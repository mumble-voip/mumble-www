all: public mumble-www package

public: $(shell find hugo/ -type f)
	(cd hugo/; hugo)

mumble-www: mumble-www.go githubcache.go config.go downloads.go programflags.go
	(cd src/; GO111MODULE=on go build -o ../$@ $^)

package: public mumble-www
	7z a mumble-www.7z public mumble-www
