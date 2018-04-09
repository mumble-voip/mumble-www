MUMBLE_DIR=/tmp/mumble-voip-mumble
MUMBLE_GIT=git@github.com:mumble-voip/mumble.git

public: external
	hugo

.PHONY: public

external: static/LICENSE static/AUTHORS

.PHONY: external

gitfetch:
	test -d $(MUMBLE_DIR)/.git || git clone --depth 1 $(MUMBLE_GIT) $(MUMBLE_DIR)
	test $(MUMBLE_GIT) = $(shell git --git-dir=$(MUMBLE_DIR)/.git remote get-url origin)
	git --git-dir=$(MUMBLE_DIR)/.git fetch -q origin
	git --git-dir=$(MUMBLE_DIR)/.git checkout -q origin/master

.PHONY: gitfetch

static/LICENSE: gitfetch
	cp $(MUMBLE_DIR)/LICENSE $@

static/AUTHORS: gitfetch
	cp $(MUMBLE_DIR)/AUTHORS $@
