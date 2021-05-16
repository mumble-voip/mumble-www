---
title: "Mumble Overlay Configuration"
---
The Mumble settings have an Overlay category which allows you to configure how and what the overlay displays as well as in what applications it displays.

## Configuring in what applications it displays

The *Overlay Exceptions* define in what applications the overlay will display.

| Mode            | Description |
| --------------- | ----------- |
| Launcher Filter | Supports launchers where child processes will show an overlay (e.g. Steam, Origin, etc. launch a game) |
| Whitelist       | Overlay is disabled by default. The specified applications will show the overlay. |
| Blacklist       | Overlay is enabled by default (for any applications using rendering). The specified applications will not show the overlay. |
| `nooverlay` file | If a file named `nooverlay` (no file extension) is placed next to an application executable, the overlay will not show. |
| `debugoverlay` file | If a file named `debugoverlay` (no file extension) is placed next to an application executable, the more debug information messages are logged to the debug log. |
