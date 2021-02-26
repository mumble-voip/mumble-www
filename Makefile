all: public package

public: $(shell find hugo/ -type f)
	(cd hugo/; hugo)

package: public
	7z a mumble-www.7z public
