package main

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var snapshotRegexps = map[string]*regexp.Regexp{
	"windows-64":    regexp.MustCompile(`^mumble-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.winx64\.msi$`),
	"windows-32":    regexp.MustCompile(`^mumble-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.msi$`),
	"osx":           regexp.MustCompile(`^Mumble-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.dmg$`),
	"osx-universal": regexp.MustCompile(`^Mumble-Universal-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.dmg$`),
	// "tarball": regexp.MustCompile(`^mumble-[\.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.tar\.gz$`),
	"linux-static-server": regexp.MustCompile(`^murmur-static_x86-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.tar\.bz2$`),
	"osx-static-server":   regexp.MustCompile(`^Murmur-OSX-Static-[.0-9]*~[0-9]*~g[0-9a-z]*~snapshot\.tar\.bz2$`),
}

type SnapshotCache struct {
	mu    sync.RWMutex
	links map[string]string
}

func NewSnapshotCache(cacheTime time.Duration) *SnapshotCache {
	s := &SnapshotCache{}
	go s.updater(cacheTime)
	return s
}

func (s *SnapshotCache) updater(cacheTime time.Duration) {
	ticker := time.NewTicker(cacheTime)
	defer ticker.Stop()

	const root = `https://dl.mumble.info/?C=M;O=D` // sorted by last modified first
	rootURL, _ := url.Parse(root)

	for {
		doc, err := goquery.NewDocument(root)
		if err == nil {
			newLinks := make(map[string]string, len(snapshotRegexps))

			doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
				for snapshotName, snapshotRe := range snapshotRegexps {
					if _, alreadySeen := newLinks[snapshotName]; alreadySeen {
						continue
					}
					href, hasHref := s.Attr("href")
					if !hasHref {
						continue
					}

					if !snapshotRe.MatchString(href) {
						continue
					}

					linkUrl, err := url.Parse(href)
					if err != nil {
						continue
					}

					newLinks[snapshotName] = rootURL.ResolveReference(linkUrl).String()

					if len(newLinks) == len(snapshotRegexps) {
						return false
					}
				}
				return true
			})

			s.mu.Lock()
			s.links = newLinks
			s.mu.Unlock()
		} else {
			log.Printf("Snapshot cache error: %s", err)
		}
		<-ticker.C
	}
}

func (s *SnapshotCache) Handle(linkName string) http.Handler {
	if _, isTracking := snapshotRegexps[linkName]; !isTracking {
		return http.NotFoundHandler()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.RLock()
		link, hasLink := s.links[linkName]
		s.mu.RUnlock()

		if !hasLink {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		http.Redirect(w, r, link, http.StatusTemporaryRedirect)
	})
}
