---
title: 1.4.0 Development Snapshot 4
author: Robert Adam
date: 2021-04-03
categories:
  - Release
  - Development Snapshot
---

About two weeks have passed since our last release ([third snapshot]({{< relref "/blog/2021-03-21-mumble-1-4-0-snapshot3" >}})). In that time we were
able to fix a few bugs and work on the overall polish a bit more, so that we can release the forth snapshot of 1.4.0 today.

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

The changes listed here are relative to the last snapshot version.

### Client

- Added: Allow gemini:// protocol for links in chat messages ({{< issue 4888 >}})
- Fixed: No users being added to the TalkingUI ({{< issue 4884 >}})
- Fixed: Line breaks not being respected in chat messages (sent directly via chat bar) ({{< issue 4902 >}})
- Revamped: Server information dialog ({{< issue 4891 >}})


### Server

- Fixed: Prevent multiple sessions with ID 0 (causing different users to appear as a single one) ({{< issue 4886 >}})


### Positional audio plugins

- Updated: Among Us plugin to work with v2021.3.31.3s ({{< issue 4901 >}})
