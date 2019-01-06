// Copyright 2005-2018 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func setup(mux *http.ServeMux) {
	setupDownloadsStable(mux)
	setupDownloadsSnapshot(mux)
	setupGithubCache(mux)
	setupStaticFiles(mux)
}

func setupDownloadsStable(mux *http.ServeMux) {
	stable := map[string]string{
		"windows-32":          "https://github.com/mumble-voip/mumble/releases/download/1.2.19/mumble-1.2.19.msi",
		"osx":                 "https://github.com/mumble-voip/mumble/releases/download/1.2.19/Mumble-1.2.19.dmg",
		"osx-universal":       "https://github.com/mumble-voip/mumble/releases/download/1.2.10/Mumble-Universal-1.2.10.dmg",
		"ios":                 "http://itunes.apple.com/us/app/mumble/id443472808?ls=1&mt=8",
		"ubuntu":              "https://launchpad.net/~mumble/+archive/release",
		"linux-static-server": "https://github.com/mumble-voip/mumble/releases/download/1.2.19/murmur-static_x86-1.2.19.tar.bz2",
		"osx-static-server":   "https://github.com/mumble-voip/mumble/releases/download/1.2.19/Murmur-OSX-Static-1.2.19.tar.bz2",
	}

	for source, target := range stable {
		mux.Handle("/downloads/"+source, http.RedirectHandler(target, http.StatusTemporaryRedirect))
	}
}

func setupDownloadsSnapshot(mux *http.ServeMux) {
	snapshotCache := NewSnapshotCache(time.Hour)
	for source := range snapshotRegexps {
		mux.Handle("/downloads/"+source+"/snapshot", snapshotCache.Handle(source))
	}
	mux.Handle(
		"/downloads/ubuntu/snapshot",
		http.RedirectHandler("https://launchpad.net/~mumble/+archive/snapshot", http.StatusTemporaryRedirect),
	)
}

func setupGithubCache(mux *http.ServeMux) {
	githubCache := NewGithubCache(time.Minute * 10)
	mux.Handle("/LICENSE", githubCache.Handle("/LICENSE"))
	mux.Handle("/AUTHORS", githubCache.Handle("/AUTHORS"))
}

func setupStaticFiles(mux *http.ServeMux) {
	public := flag.String("public", "public", "public directory")
	fs := http.FileServer(http.Dir(*public))
	mux.Handle("/", fs)
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "bind address")
	flag.Parse()

	log.Printf("Starting server on %s\n", *addr)

	mux := http.NewServeMux()

	setup(mux)

	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
