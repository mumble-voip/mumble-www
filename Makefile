all: public mumble-www

public: $(shell find hugo/ -type f)
	(cd hugo/; hugo)

mumble-www: src/mumble-www.go src/githubcache.go src/snapshotcache.go
	go build -o $@ $^
