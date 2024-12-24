---
title: Mumble 1.3.2 Release Announcement
author: Robert Adam
date: 2020-07-09
categories:
  - Release
---
The Mumble team has released [**version 1.3.2**](https://github.com/mumble-voip/mumble/releases/tag/1.3.2) of the Mumble VoIP application. This is a
**bugfix release** of the current and stable 1.3 series.

You can download the new version from our [**Downloads page**]({{< relref "/downloads" >}}) or the
[**GitHub release page**](https://github.com/mumble-voip/mumble/releases/tag/1.3.2) or from within your Windows client or software package management
system.

<!--more-->

## Changes in this Version

### Client

Fixed: Overlay not starting ({{<issue 4282 >}})

### Server

Fixed: keychain-error on macOS for custom certificates ({{<issue 4345 >}})


## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode
