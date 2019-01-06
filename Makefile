all: public mumble-www package

public: $(shell find hugo/ -type f)
	(cd hugo/; hugo)

mumble-www: src/mumble-www.go src/githubcache.go src/snapshotcache.go
	go build -o $@ $^

package: public mumble-www
	7z a mumble-www.7z public mumble-www
