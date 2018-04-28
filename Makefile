all: public mumble-www

public: $(shell find hugo/ -type f)
	(cd hugo/; hugo)

mumble-www: mumble-www.go githubcache.go snapshotcache.go
	go build -o $@ $^
