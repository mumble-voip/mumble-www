---
title: 1.4.0 Release Candidate 1
author: Robert Adam
date: 2021-09-12
categories:
  - Release
  - Release Candidate
---

Quite some time has passed since our [last snapshot]({{< relref ""2021-05-16-mumble-1-4-0-snapshot6.md >}}) back in May of this year. Since then we
have worked hard to polish things up further and are therefore now able to present to you the first _release candidate_ (RC) of Mumble 1.4.0.

You can download the new version from our [**Downloads page**]({{< relref "/downloads#development-snapshots" >}}).

<!--more-->


## Changelog

The changes listed here are relative to the last snapshot version.

### Client

- Changed: Increased minimum macOS version to 10.13 (which is what Qt requires) {{< issue 5250 >}}
- Changed: Display unit as suffix in SpinBox ({{< issue 5017 >}})
- Changed: Do not play the mute cue when push-to-muted ({{< issue 5047 >}})
- Changed: Resize posted images to 800x600 instead of 420x270 ({{< issue 5050 >}})
- Fixed: Ambiguity in plugin installer ({{< issue 5123 >}})
- Fixed: Application freeze when cancelling audio wizard ({{< issue 5083 >}})
- Fixed: Crash when using the PipeWire audio backend ({{< issue 5080 >}})
- Fixed: Issues when updating installed plugins on Windows ({{< issue 5167 >}})
- Fixed: Migrate Windows shortcuts from versions older than 1.4.0 properly ({{< issue 5205 >}})
- Fixed: PluginInstaller unable to extract zips on Windows ({{< issue 5109 >}})
- Fixed: Positional audio not working properly after canceling audio wizard ({{< issue 5046 >}})
- Fixed: Potential deadlocks in plugins ({{< issue 5163 >}})
- Fixed: PulseAudio not initializing ({{< issue 5184 >}})
- Fixed: Remove unnecessary waiting during application startup (Mumble now starts up way faster) ({{< issue 5168 >}})
- Fixed: onAudioOutputAboutToPlay plugin API function used wrong parameter order ({{< issue 5115 >}})
- Fixed: requestLocalUserTransmissionMode plugin API function now properly integrates with UI ({{< issue 5116 >}})


### Server

- Fixed: Always bind to both IPv6 and IPv4 by default ({{< issue 5212 >}})
- Fixed: Database upgrade path ignoring specific data field ({{< issue 5142 >}})
- Fixed: Missing lock in Ice-call setACL ({{< issue 5136 >}})
- Fixed: Possible DB corruption due to missing locks ({{< issue 5044 >}})
- Fixed: Tray icon not shown on Windows ({{< issue 5173 >}})
- Fixed: Wrong "Unable to find matching CELT codec" warning upon connecting ({{< issue 5112 >}})


### Positional audio plugins

- Added: Update & port GTA5 to new Plugin API ({{< issue 5162 >}})
- Fixed: Update Among Us plugin to work with v2021.6.30s ({{< issue 5189 >}})
- Fixed: Update Source Engine plugin to work with L4D2 2.2.2.0 ({{< issue 5190 >}})


## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode
- Autoscroll of chat window not working properly on Linux ({{< issue 4638 >}}, {{< issue 2504 >}})
- Certain special characters are not rendered on Windows ({{< issue 4939 >}})
- The RC is still being displayed as a "snapshot"
