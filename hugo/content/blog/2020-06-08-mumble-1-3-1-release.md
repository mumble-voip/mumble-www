---
title: Mumble 1.3.1 Release Announcement
author: Jan Klass
date: 2020-06-08
categories:
  - News
  - Release
  - Security
---
The Mumble team has released [**version 1.3.1**](https://github.com/mumble-voip/mumble/releases/tag/1.3.1) of the Mumble VoIP application. This is a **bugfix and security release** of the current and stable 1.3 series.

You can download the new version from our [**Downloads page**]({{< relref "/downloads" >}}) or the [**GitHub release page**](https://github.com/mumble-voip/mumble/releases/tag/1.3.1) or from within your Windows client or software package management system.

<!--more-->

## Changes in this Version

### Security

* Fixed: Potential exploit in the OCB2 encryption ({{< issue 4227 >}})

### ICE

* Fixed: Added missing UserKDFIterations field to UserInfo  =>  Prevents getRegistration() from failing with enumerator out of range error ({{< issue 3835 >}})

### GRPC

* Fixed: Segmentation fault during murmur shutdown ({{< issue 3938 >}})

### Client

* Fixed: Crash when using multiple monitors ({{< issue 3756 >}})
* FIxed: Don't send empty message from clipboard via shortcut, if clipboard is empty ({{< issue 3864 >}})
* Fixed: Talking indicator being able to freeze to indicate talking when self-muted ({{< issue 4006 >}})
* Fixed: High CPU usage for update-check if update server not available ({{< issue 4019 >}})
* Fixed: DBus getCurrentUrl returning empty string when not in root-channel ({{< issue 4029 >}})
* Fixed: Small parts of whispering leaking out to normal talk ({{< issue 4051 >}})
* Fixed: Last audio frame of normal talking sent to last whisper target instead when using VoiceActivation ({{< issue 4050 >}})
* Fixed: LAN-icon not found in ConnectDialog ({{< issue 4058 >}})
* Improved: Set maximal vertical size for User Volume Adjustment dialog ({{< issue 3801 >}})
* Improved: Don't send empty data to PulseAudio ({{< issue 3316 >}})
* Improved: Use the SRV resolved port for UDP connections ({{< issue 3820 >}})
* Improved: Manual Plugin UI ({{< issue 3919 >}})
* Improved: Don't start Jack server by default ({{< issue 3990 >}})
* Improved: Overlay doesn't hook into all other processes by default ({{< issue 4041 >}})
* Improved: Wait longer before disconnecting from a server due to unanswered Ping-messages ({{< issue 4123 >}})

### Server

* Fixed: Possibility to circumvent max user-count in channel ({{< issue 3880 >}})
* Fixed: Rate-limit implementation susceptible to time-underflow ({{< issue 4004 >}})
* Fixed: OpenSSL error 140E0197 with Qt >= 5.12.2 ({{< issue 4032 >}})
* Fixed: VersionCheck for SQL for when to use the WAL feature ({{< issue 4163 >}})
* Fixed: Wrong database encoding that could lead to server-crash ({{< issue 4220 >}})
* Fixed: DB crash due to primary key violation (now performs "UPSERT" to avoid this) ({{< issue 4105 >}})
* Improved: The fields in the `Version` ProtoBuf message are now size-restricted in order to avoid attacks that can render another client unresponsive ({{< issue 4101 >}})

### Windows Installer

* Improved: Mumble icon (now properly displayed) ({{< issue 4204 >}})
