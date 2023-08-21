---
title: "Mumble URL"
---
Mumble defines it's own `mumble://` URL scheme to share links to servers and channels and to be able to connect to them directly.

On Windows the installer will register `mumble://` links to be opened with Mumble.

Mumble links can be opened on the command line by passing them to mumble.exe too:

```
mumble.exe mumble://example.com/?version=1.2.0
```

or with a channel path:

```
mumble.exe mumble://example.com/ChanA/ChanB?version=1.2.0
```

You can get get a URL in the Channel context menu (right click):

![Screenshot of Mumble channel context menu Copy URL item](screenshot-chan-copy-url.jpg)

## Connect Behavior

When opening a Mumble link the Mumble client will connect to the server.

* If you are already registered you will connect with your account
* If the URL contains a username that username will be used
* If no username is determined, the client will ask with what name you want to connect with

## URL Format

```
mumble://[username[:password]@]<address>[:port]/[channelpath][?version=<serverversion>][&title=<servername>][&url=<serverurl>]
```

* optional `username` and `password`
* mandatory `address`; a hostname, IPv4, or IPv6 address
* optional `port`; defaults to 64738
* optional `channelpath`; specifies a channel subtree to join; special characters must be [URL-encoded](https://en.wikipedia.org/wiki/URL_encoding); if the Mumble client can join the server but not the channel it will stay in the default channel
* optional `serverversion` - protocol version of the server; Should no longer be necessary; Relevant for backwards compatibility to pre-1.2.0 versions
* optional `servername` and `serverurl`; only used if the user drags the link and drops it in the server browser in Mumble for filling the server name and URL

"Special" non-ASCII characters must be %-encoded. Most utilities should do this for you. For more information see [Wikipedia - URL encoding](https://en.wikipedia.org/wiki/URL_encoding).

Simple form:

```
mumble://<address>:<port>/?version=1.2.0
```

Examples:

```
mumble://mumble.example.com/?version=1.2.0
mumble://172.16.31.199/?version=1.2.0
mumble://example.com:23840/?version=1.2.0
mumble://172.16.31.199:2000/?version=1.2.0
```

## Server Password

A server password can be specified without a username:

```
mumble://:mypassword@example.com/?version=1.2.0
```

## URL Handler Installation

A manual handler installation should not be necessary for normal users.

### Linux Gnome

```bash
gconftool-2 -s -t string /desktop/gnome/url-handlers/mumble/command 'mumble "%s"'
gconftool-2 -s -t bool /desktop/gnome/url-handlers/mumble/enabled true
```

If the `mumble` command is not globally accessible through the `PATH` environment variable you must specify the full path to the mumble executable.
