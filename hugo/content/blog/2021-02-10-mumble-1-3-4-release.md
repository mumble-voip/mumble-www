---
title: Mumble 1.3.4 Release Announcement
author: Robert Adam
date: 2021-01-10
categories:
  - Release
  - Security
---
The Mumble team has released [**version 1.3.4**](https://github.com/mumble-voip/mumble/releases/tag/1.3.4) of the Mumble VoIP application. This is a
**bugfix and security release** of the current and stable 1.3 series.

You can download the new version from our [**Downloads page**]({{< relref "/downloads" >}}) or the
[**GitHub release page**](https://github.com/mumble-voip/mumble/releases/tag/1.3.4) or from within your Windows client or software package management
system.

<!--more-->

## Changes in this Version

### Client

- Fixed: Don't use outdated (non-existent) notification icon names on Linux ({{< issue 4705 >}})
- Fixed: Security vulnerability caused by allowing non http/https URL schemes in public server list ({{< issue 4733 >}})

### Server

- Fixed: Exit status for actions like `--version` or `--supw` was always set to non-zero ({{< issue 3998 >}})

### General

- Fixed: Packet loss & audio artifacts caused by OCB2 XEX* mitigation ({{< issue 4720 >}})

## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode.

## Warning

- The static server binary for Linux is using an outdated version of OpenSSL - see https://github.com/mumble-voip/mumble/issues/4001 for details. This
  warning only applies to the package released by us named  `murmur-static_x86`. It does not apply to our Ubuntu PPA releases or other maintained
  packages.

