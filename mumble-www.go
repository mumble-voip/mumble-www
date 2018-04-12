// Copyright 2005-2018 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "bind address")
	public := flag.String("public", "public", "public directory")
	flag.Parse()

	log.Printf("Starting server on %s\n", *addr)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(*public))

	mux.HandleFunc("/grumble/", func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = "/grumble/"
		fs.ServeHTTP(w, r2)
	})

	{
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

	{
		// TODO: snapshots
		mux.Handle(
			"/downloads/ubuntu/snapshot",
			http.RedirectHandler("https://launchpad.net/~mumble/+archive/snapshot", http.StatusTemporaryRedirect),
		)
	}

	githubCache := NewGithubCache(time.Minute * 10)
	mux.Handle("/LICENSE", githubCache.Handle("/LICENSE"))
	mux.Handle("/README", githubCache.Handle("/README"))

	mux.Handle("/", fs)

	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
