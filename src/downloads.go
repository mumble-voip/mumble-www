package main

import (
	"log"
	"os"
	"sync"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/fsnotify/fsnotify"
)

var mutex sync.RWMutex
var redirectMap = make(map[string]string)

type Release struct {
	Version string `json:"version"`
	Url string `json:"url"`
	SigUrl string  `json:"sigUrl"`
}

func setupDownloads(mux *http.ServeMux, jsonPath string) {
	mux.HandleFunc("/downloads/", downloadsHandler)
	refreshDownloads(jsonPath)
	go fileWatcher(jsonPath)
}

func refreshDownloads(jsonPath string) {
	mutex.Lock()

	redirectMap = make(map[string]string)

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var jsonMap map[string]map[string]Release
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	if value, ok := jsonMap["stable"]; ok {
		setupDownloadsStable(value)
	}

	if value, ok := jsonMap["snapshot"]; ok {
		setupDownloadsSnapshot(value)
	}

	for key, value := range redirectMap {
		log.Println("downloads/" + key + " -> " + value)
	}

	mutex.Unlock()
}

func setupDownloadsStable(jsonMap map[string]Release) {
	redirectMap["ubuntu"] = "https://launchpad.net/~mumble/+archive/release"
	redirectMap["ios"] = "https://apps.apple.com/app/mumble/id443472808"

	for key, value := range jsonMap {
		redirectMap[key] = value.Url
	}
}

func setupDownloadsSnapshot(jsonMap map[string]Release) {
	redirectMap["ubuntu/snapshot"] = "https://launchpad.net/~mumble/+archive/snapshot"

	for key, value := range jsonMap {
		redirectMap[key + "/snapshot"] = value.Url
	}
}

func downloadsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the part after "/downloads/"
	path := r.URL.Path[len("/downloads/"):]

	mutex.RLock()

	value, ok := redirectMap[path]

	mutex.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, value, http.StatusTemporaryRedirect)
}

func fileWatcher(filePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					log.Println(event)

					if event.Op&fsnotify.Write == fsnotify.Write {
						refreshDownloads(filePath)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}

					log.Println(err)
			}
		}
	}()

	if err := watcher.Add(filePath); err != nil {
		log.Println(err)
	}

	<-done
}
