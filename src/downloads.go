package main

import (
	"net/http"
	"time"
)

func setupDownloadsStable(mux *http.ServeMux) {
	stable := map[string]string{
		"windows-32":          "https://github.com/mumble-voip/mumble/releases/download/1.3.4/mumble-1.3.4.msi",
		"windows-64":          "https://github.com/mumble-voip/mumble/releases/download/1.3.4/mumble-1.3.4.winx64.msi",
		"osx":                 "https://github.com/mumble-voip/mumble/releases/download/1.3.3/Mumble-1.3.3.dmg",
		"osx-hfs+":            "https://github.com/mumble-voip/mumble/releases/download/1.3.3/Mumble-1.3.3-HFS+.dmg",
		"ios":                 "http://itunes.apple.com/us/app/mumble/id443472808?ls=1&mt=8",
		"ubuntu":              "https://launchpad.net/~mumble/+archive/release",
		"linux-static-server": "https://github.com/mumble-voip/mumble/releases/download/1.3.4/murmur-static_x86-1.3.4.tar.bz2",
		"osx-static-server":   "https://github.com/mumble-voip/mumble/releases/download/1.3.3/Murmur-OSX-Static-1.3.3.tar.bz2",
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
