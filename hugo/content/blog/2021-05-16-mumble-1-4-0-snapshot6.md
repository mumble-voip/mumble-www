---
title: 1.4.0 Development Snapshot 6
author: Robert Adam
date: 2021-05-16
categories:
  - Release
  - Development Snapshot
---

Following the [last snapshot]({{< relref "2021-04-21-mumble-1-4-0-snapshot5" >}}) we proudly present our sixth development snapshot of the upcoming 1.4.0 stable
release. Note that this snapshot is intended to be the last one before the final release and as such with its publication we are introducing a
feature-freeze of the 1.4.x series. From now on only bug fixes will make it into this series.

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

- Added: Global shortcut to toggle TalkingUI ({{< issue 4958 >}})
- Added: Implement a mute cue ({{< issue 4926 >}})
- Added: Implement native support for PipeWire ({{< issue 4970 >}})
- Added: Implement raw input (Replaces hooks and DirectInput) on Windows ({{< issue 4941 >}})
- Added: Improvements for using --multiple command-line flag ({{< issue 4949 >}})
- Added: Revamp positional audio settings page ({{< issue 4995 >}})
- Added: Search dialog ({{< issue 4967 >}})
- Changed: Default min volume to 25% (positional audio) ({{< issue 4954 >}})
- Fixed: Links containing special characters (e.g. "&") being broken ({{< issue 5000 >}})
- Fixed: Echo cancellation options not translated ({{< issue 4994 >}})
- Fixed: Memory management issue in the source code ({{< issue 4982 >}})
- Fixed: Properly show currently selected audio device in settings dialog ({{< issue 4974 >}})
- Fixed: TalkingUI font not scaling properly ({{< issue 4991 >}})
- Removed: Classic theme ({{< issue 4969 >}})


### Server

- Fixed: ChannelListener IDs colliding across VServers ({{< issue 5003 >}})


## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode
- Autoscroll of chat window not working properly on Linux ({{< issue 4638 >}}, {{< issue 2504 >}})
- Certain special characters are not rendered on Windows ({{< issue 4939 >}})
