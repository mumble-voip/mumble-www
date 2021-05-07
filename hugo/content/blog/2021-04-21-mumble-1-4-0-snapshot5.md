---
title: 1.4.0 Development Snapshot 5
author: Robert Adam
date: 2021-04-21
categories:
  - News
  - Development Snapshot
---

With the arrival of the brand new plugin framework for Mumble, we release the fifth snapshot for the upcoming 1.4.0 stable release. Apart from that
there were also a couple of other changes since the [last snapshot](2021-04-03-mumble-1-4-0-snapshot4.md).

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

- Added: Plugin framework ({{< issue 3743 >}})
- Changed: Disable echo cancellation on macOS by default (was reported to cause trouble) ({{< issue 4914 >}})
- Fixed: CrashReporter considering 2xx codes as errors ({{< issue 4929 >}})
- Fixed: Memory leak in ALSA implementation ({{< issue 4916 >}})
- Fixed: Minor bugs in ALSA implementation ({{< issue 4920 >}})


### Server

- Changed: Allow spaces in username by default ({{< issue 4925 >}})
- Fixed: Fix undefined behavior on Linux hosts ({{< issue 4915 >}})


### Miscellaneous

- Fixed: Cryptographic init potentially failing silently ({{< issue 4903 >}})

