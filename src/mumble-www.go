// Copyright 2005-2019 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setup(mux *http.ServeMux, config *Config) {
	setupDownloads(mux, config.JsonPath)
	setupGithubCache(mux, config.GitHubCacheTimeMins)
	setupStaticFiles(mux, config.PublicPath)
}

func setupGithubCache(mux *http.ServeMux, cacheTimeMins int) {
	githubCache := NewGithubCache(time.Minute * time.Duration(cacheTimeMins))
	mux.Handle("/LICENSE", githubCache.Handle("/LICENSE"))
	mux.Handle("/AUTHORS", githubCache.Handle("/AUTHORS"))
}

func setupStaticFiles(mux *http.ServeMux, path string) {
	fs := http.FileServer(http.Dir(path))
	mux.Handle("/", fs)
}

func main() {
	config := ParseInit()
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	log.Printf("Starting server on %s\n", addr)

	mux := http.NewServeMux()

	setup(mux, config)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
