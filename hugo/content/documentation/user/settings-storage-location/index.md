---
title: "Mumble Settings Storage Location"
---
The Mumble client uses two different locations for storing data:

1. A database that contains the server list, server certificates, favorites, friends, access tokens, comments, etc.
2. Settings for the client configuration

## Windows

1. Database: `%APPDATA%\Mumble\mumble.sqlite`
2. Settings: In the Windows registry under `HKEY_CURRENT_USER\Software\Mumble\Mumble`

## Mac OS X

1. Database: `$HOME/Library/Application Support/Mumble/Mumble/.mumble.sqlite`
2. Settings: `$HOME/Library/Preferences/net.sourceforge.mumble.Mumble.plist`

## Linux

1. Database `$HOME/.local/share/Mumble/Mumble/mumble.sqlite`
2. Settings `$HOME/.config/Mumble/Mumble.conf`
