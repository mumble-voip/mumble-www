// Copyright 2005-2018 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

package main

import (
	"context"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/github"
)

type GithubCache struct {
	client *github.Client

	newFile chan struct{}

	mu    sync.Mutex
	files map[string]*githubCacheFile
}

type githubCacheFile struct {
	mu          sync.RWMutex
	lastUpdated time.Time
	contents    string
}

func NewGithubCache(cacheTime time.Duration) *GithubCache {
	cache := &GithubCache{
		client: github.NewClient(nil),

		newFile: make(chan struct{}, 1),

		files: make(map[string]*githubCacheFile),
	}
	go cache.updater(cacheTime)
	return cache
}

func (c *GithubCache) updater(cacheTime time.Duration) {
	ticker := time.NewTicker(cacheTime)
	defer ticker.Stop()

	for {
		c.mu.Lock()
		for filename, entry := range c.files {
			fileContent, _, _, err := c.client.Repositories.GetContents(context.Background(), "mumble-voip", "mumble", filename, nil)
			if err != nil {
				log.Printf("GitHubCache GetContents Error: %s", err)
				continue
			}
			// TODO: support files larger than 1MB
			content, err := fileContent.GetContent()
			if err != nil {
				log.Printf("GitHubCache GetContents Error: %s", err)
				continue
			}

			entry.mu.Lock()
			entry.lastUpdated = time.Now()
			entry.contents = content
			entry.mu.Unlock()
		}
		c.mu.Unlock()

		select {
		case <-ticker.C:
		case <-c.newFile:
		}
	}
}

func (c *GithubCache) signalNewFile() {
	select {
	case c.newFile <- struct{}{}:
	default:
	}
}

func (c *GithubCache) Handle(filename string) http.Handler {
	if !path.IsAbs(filename) {
		panic("filename not an absolute path")
	}
	filename = path.Clean(filename)
	basename := filepath.Base(filename)

	c.mu.Lock()
	cacheEntry, ok := c.files[filename]
	if !ok {
		cacheEntry = &githubCacheFile{}
		c.files[filename] = cacheEntry
		c.signalNewFile()
	}
	c.mu.Unlock()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheEntry.mu.RLock()
		if cacheEntry.lastUpdated.IsZero() {
			// File has not been fetched yet
			cacheEntry.mu.RUnlock()
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		reader := strings.NewReader(cacheEntry.contents)
		cacheEntry.mu.RUnlock()

		http.ServeContent(w, r, basename, time.Time{}, reader)
	})
}
