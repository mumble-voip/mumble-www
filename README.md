# The Mumble Website

**This is a work in progress not yet released.**

Source of the Mumble Website available at [mumble.info](https://mumble.info/).

Implemented and served through a combination of a Go webserver for dynamic content and static files generated with [hugo](https://gohugo.io/).

## Webserver

* Cache the `LICENSE` and `AUTHORS` files from github and serve them
* Redirect predefined stable download URLs to download files
* Redirect predefined snapshot download URLs; cached self-updating target URLs
* * Redirect Ubuntu snapshots to the PPA
* Serve static website files

The go source files reside in the `src` folder.

## Website

The website source resides in the `hugo` folder, is buildable with hugo. The target folder is `public`.
