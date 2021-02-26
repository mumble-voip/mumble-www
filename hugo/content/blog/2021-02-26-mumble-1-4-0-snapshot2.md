---
title: 1.4.0 Development Snapshot 2
author: Robert Adam
date: 2021-02-26
categories:
  - News
  - Development Snapshot
---

After the [first snapshot]({{< relref "/blog/2020-12-24-mumble-1-4-0-snapshot" >}}) has been released
on Christmas last year (2020), we now present you an updated version of the snapshot version of the upcoming
feature release 1.4.0.

After having understood what it means for a release to be a _snapshot release_ (see below), you can download the new version from our [**Downloads
page**]({{< relref "/downloads#development-snapshots" >}}).

<!--more-->

## Development Snapshots

A _development snapshot_ can be considered an open-beta for an upcoming feature release. This means that it contains the latest and greatest features
that were introduced since the last feature release. Note though that in contrast to a _release candidate_ the list of features that will be included
in the final 1.4.0 release is _not_ yet frozen and is thus subject to change.

This does also mean that these features have not been tested by a broad audience yet and are therefore likely to still contain bugs. These could be
minor ones, but in theory these could also be major. Therefore it is important that you always back up your data (e.g. configuration files) if you
have a previous version of Mumble installed.

Due to the reasoning stated above, we do not recommend using snapshots in an environment where stability of the software is key (e.g. in a business
context).

That being said though, we strongly encourage you to try snapshots out if you are a tech-savvy person that can deal with potentially arising problems
and report any problems you encounter to our [**issue tracker**](https://github.com/mumble-voip/mumble/issues). This kind of feedback is very
important for us to be able to make the actual release as stable and reliable as possible.

## Changelog

The changes listed here are relative to the last snapshot version. Note also that the development snapshot also
includes all fixes that were shipped in the recently released version [1.3.4](https://www.mumble.info/blog/mumble-1.3.4-release-announcement/).

### Client

- Added: "starttalking" and "stoptalking" to the socket RPC interface ({{< issue 4754 >}})
- Added: New icons (padlocks, new talking state, listener symbol) to "What's this?" ({{< issue 4690 >}})
- Added: DBus calls to toggle mute and deaf state ({{< issue 4667 >}})
- Added: Prompt user for microphone permission on macOS if not granted already ({{< issue 4676 >}})
- Added: Echo Cancellation for macOS ({{< issue 4694 >}})
- Fixed: Various issues in the "About" dialog ({{< issue 4785 >}})
- Fixed: Channel description available via lock icon ({{< issue 4688 >}})
- Fixed: Crash when pressing Esc in nickname dialog ({{< issue 4671 >}})
- Fixed: Don't show URL password in clear text in chat log ({{< issue 4779 >}})
- Fixed: Duplicate keyboard shortcut "c" in server browser ({{< issue 4777 >}})
- Fixed: Fix protocol version not being set correctly ({{< issue 4683 >}})
- Fixed: Wrong "Image not supported" message ({{< issue 4661 >}})
- Fixed: Crash after rename ({{< issue 4764 >}})
- Fixed: Issues with Mumble URL versioning ({{< issue 4778 >}})
- Improved: Shortcut representation and interaction in settings ({{< issue 4722 >}})


### Server

- Changed: Make default max bandwidth the highest supported ({{< issue 4700 >}})
- Fixed: Fix protocol version not being set correctly ({{< issue 4683 >}})
- Fixed: Older clients not being warned about ChannelListener ({{< issue 4689 >}})


### Positional audio plugins

- Updated: Among Us plugin to work with v2020.12.09s ({{< issue 4663 >}})

