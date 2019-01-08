package main

import (
	"net/http"
	"time"
)

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
