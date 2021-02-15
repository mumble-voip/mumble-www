# The Mumble Website

Source of the Mumble Website. The live website is available at [www.mumble.info](https://www.mumble.info).

Implemented and served through a combination of a Go webserver for dynamic content and static files generated with [hugo](https://gohugo.io/).

## Go Webserver

* Cache the `LICENSE` and `AUTHORS` files from GitHub and serve them
* Redirect predefined stable download URLs to download files
* Redirect predefined snapshot download URLs; cached self-updating target URLs
  * Redirect Ubuntu snapshots to the PPA
* Serve static website files (generated with Hugo)

The go source files reside in the `src` folder.

## Static Website

The website source resides in the `hugo` folder.

It is buildable with [Hugo](gohugo.io/). After navigating into the `hugo` folder: The `hugo` command will place the generated website files in the `public` folder. An interactive, auto-updating development server can be launched with `hugo server`.
