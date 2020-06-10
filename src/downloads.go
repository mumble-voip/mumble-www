package main

import (
	"net/http"
	"time"
)

func setupDownloadsStable(mux *http.ServeMux) {
	stable := map[string]string{
		"windows-32":          "https://github.com/mumble-voip/mumble/releases/download/1.3.1/mumble-1.3.1.msi",
		"windows-64":          "https://github.com/mumble-voip/mumble/releases/download/1.3.1/mumble-1.3.1.winx64.msi",
		"osx":                 "https://github.com/mumble-voip/mumble/releases/download/1.3.1/Mumble-1.3.1.dmg",
		"osx-universal":       "https://github.com/mumble-voip/mumble/releases/download/1.2.10/Mumble-Universal-1.2.10.dmg",
		"ios":                 "http://itunes.apple.com/us/app/mumble/id443472808?ls=1&mt=8",
		"appimage-x64":        "https://dl.mumble.info/Mumble-1.3.1-x64_64.AppImage",
		"ubuntu":              "https://launchpad.net/~mumble/+archive/release",
		"linux-static-server": "https://github.com/mumble-voip/mumble/releases/download/1.3.1/murmur-static_x86-1.3.1.tar.bz2",
		"osx-static-server":   "https://github.com/mumble-voip/mumble/releases/download/1.3.1/Murmur-OSX-Static-1.3.1.tar.bz2",
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
