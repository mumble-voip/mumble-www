---
title: 1.4.0 Development Snapshot 3
author: Robert Adam
date: 2021-03-21
categories:
  - News
  - Release
  - Development Snapshot
---

After the [second snapshot]({{< relref "/blog/2021-02-26-mumble-1-4-0-snapshot2" >}}) has been released about a month ago, we now present you an
updated version of the snapshot of the upcoming feature release 1.4.0.

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

- Added: Improvements to local translation dir search (useful for translators) ({{< issue 4820 >}})
- Fixed: Output devices being ignored on macOS for echo cancellation ({{< issue 4814 >}})
- Fixed: Make "Defaults (all)" button in config dialog translatable ({{< issue 4867 >}})
- Fixed: Ordering of Users inconsistent on machines with different locale settings ({{< issue 4875 >}})
- Fixed: Heap corruption caused by RNNoise ({{< issue 4850 >}})
- Fixed: Use Utf8 for name field in certificate (allowing non-ASCII characters in it) ({{< issue 4872 >}})


### Server

- Added: Ability to explicitly broadcast welcome text via Ice ({{< issue 4809 >}})
- Added: Remove automatic broadcast of new welcometext when changing via Ice ({{< issue 4794 >}})
- Fixed: Ctrl-C leaves UDP port open ({{< issue 4819 >}})
- Fixed: Images being corrupted when transmitted via Ice ({{< issue 4798 >}})
- Fixed: Unauthenticated connections no longer add to user count ({{< issue 4817 >}})
- Fixed: 100ms timeout freezing when using gRPC ({{< issue 4833 >}})


### Positional audio plugins

- Updated: Call of Duty 2 ({{< issue 4868 >}})
- Updated: Among Us plugin to work with v2021.3.5s ({{< issue 4858 >}})

