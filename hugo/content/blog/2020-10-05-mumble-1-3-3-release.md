---
title: Mumble 1.3.3 Release Announcement
author: Robert Adam
date: 2020-10-05
categories:
  - News
  - Release
  - Security
---
The Mumble team has released [**version 1.3.3**](https://github.com/mumble-voip/mumble/releases/tag/1.3.3) of the Mumble VoIP application. This is a
**bugfix and security release** of the current and stable 1.3 series.

You can download the new version from our [**Downloads page**]({{< relref "/downloads" >}}) or the
[**GitHub release page**](https://github.com/mumble-voip/mumble/releases/tag/1.3.3) or from within your Windows client or software package management
system.

<!--more-->

## Changes in this Version

### Client

- Fixed: Chatbox invisble (zero height) (#4388)
- Fixed: Handling of invalid packet sizes (#4394)
- Fixed: Race-condition leading to loss of shortcuts (#4430)
- Fixed: Link in About dialog is now clickable again (#4454)
- Fixed: Sizing issues in ACL-Editor (#4455)
- Improved: PulseAudio now always samples at 48 kHz (#4449)

### Server

- Fixed: Crash due to problems when using PostgreSQL (#4370)
- Fixed: Handling of invalid package sizes (#4392) 

## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode.

## Warning

- The static server binary for Linux is using an outdated version of OpenSSL - see https://github.com/mumble-voip/mumble/issues/4001 for details.
