---
title: First Mumble 1.4.0 Development Snapshot
author: Robert Adam
date: 2020-12-24
categories:
  - News
  - Development Snapshot
---

This year we have a Christmas present for you: we are proud to announce that the first **snapshot release of Mumble 1.4.0** is released as of now. It
is a very important step towards the actual release of version 1.4.0.

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


## Feature spotlight

As there are many, many changes with this new version, here is a selection of things that we think are most prominent. If you're interested in all
changes, have a look at the full [changelog](#Changelog).

### Channel listeners

This feature can to some extent be compared to the ability to link channels. Except that this does not operate on a per-channel basis but on a
per-user basis.

Concretely this new feature allows a user to "listen to" a channel. In that case all audio that is heard by people in this particular channel (be it
by direct communication, shouts or via linked channels) is also heard by the listening user. Conceptually you could picture "listening to" a channel
with placing a microphone in a room and then listening to what is recorded.

This feature is one-way only meaning that in order to speak to people in the listened channel, you'll have to either join the channel or shout to it.

It should be noted though that you will not be secretly listened to with this new feature. Listening to a channel creates a new entry in the channel's
user-list (referred to as an "ear" or a "listener") and is thereby just as visible as a user actually being in that channel. There is also a dedicated
ACL and several new server config options that can restrict this new feature.

Older clients (< 1.4.0) will not indicate listeners in the UI though! They are informed if this feature is enabled on a server when they join it though
(via chat).

Note: This feature requires server-side support and is thus only available on servers at least running Mumble 1.4.0.

![Channel listeners](/blog/media/ChannelListeners.png)


### TalkingUI

Have you ever been in a situation where you heard someone talking but couldn't quite figure out who it was? Then this feature is probably for you!

The TalkingUI can be considered the equivalent of the already existing overlay but for non-gamers or in other terms "for everyday usage".

This feature adds an optional floating window containing a list of currently speaking users and their respective channels. This allows you to quickly
see who's currently talking and in which channel that person currently is in (this becomes important if you want to respond to that person).

The big difference between the TalkingUI and Mumble's main window (which of course also contains this information) is that the TalkingUI does not
contain all the other UI elements nor does it necessarily show the entire channel-tree. Furthermore it automatically resizes itself to the minimal
size needed to display the respective information. Thus it usually does not take up a lot of space and can therefore be left floating in the corner of
the screen without consuming too much screen real estate.

Note that the TalkingUI is highly configurable (Settings -> User Interface) and also does provides selection and context-menu support just as you get
it in the main window. In fact the selection between TalkingUI and main window is synchronized and thus this plays really nice with the new "whisper
to selected" feature allowing to whisper to whatever currently is selected in the UI (be it a channel or a user).

![TalkingUI](/blog/media/TalkingUI.png)


### UI indication for access restrictions

In previous versions of Mumble there was no indication of whether you are able to join a channel or not. You just had to try it out and either it
worked or you got a Permission Denied message. This feature solves this problem by marking "access-restricted" channels with a lock icon. Access
restricted channels are all channels that are either only open to a specific group on the server or require a password (access token) to join.

In total a channel can be in one of three states: either it is not access-restricted in which case it just looks like channels used to look prior to
the introduction of this feature or it is indeed access-restricted. In the latter case a lock icon is shown to the right of channel name. This look
can either be open and green or closed and red. Green means that you can enter the channel nonetheless (you belong to the required group or have the
correct password (access token) set) whereas red means that you are not allowed to enter the channel.

Note that this is also a feature that requires server-side support and thus only takes effect on servers that are at least running Mumble version
1.4.0.

![Access restricted channels](/blog/media/AccessRestrictedChannels.png)


### Markdown

For those of you who don't know what markdown is, we recommend a quick read of the respective [Wikipedia
article](https://en.wikipedia.org/wiki/Markdown). In simple words it allows formatting of text by simply typing a few extra characters in your text.

Mumble now supports markdown for text messages written in the normal chat bar. It is not supported for cases in which you use the dedicated message
editor. Furthermore only a subset of markdown is supported: sections, inline-code, code-blocks, bold, italic and strike-through text, links and
quotes. _Not_ supported are lists and tables.

This feature is probably best explained by an example. Try pasting the following text into the chat bar and see for yourself:

````
# Markdown
## Code
We can have `inline` code and
```
code
    blocks
```

## Emphasis
**bold** and *italic* (__bold__ and _italic_ also works) text is supported. ~~Scratch that~~

## Links
We can have [links](google.de) and www.google.de

## Quotes
> I am a quote
Although these are not pretty, they do work and are supported.

## Escaping
We can escape special characters using the \\ character: \*not italic\* and \*\*not bold\*\*
````

![Markdown](/blog/media/Markdown.png)


### Stereo playback

With version 1.4.0 Mumble (finally) learned how to play back stereo audio streams. This means that it is now possible to send a stereo stream to a
Mumble client and it will actually be played back without being mixed down to a mono stream.

Note however that this feature is currently restricted to playback. The official Mumble client will still continue sending audio as mono. This feature
is primarily intended to be made use of by (music) bots.

If you want to try this feature out yourself, you can either write a bot using [pymumble](https://github.com/azlux/pymumble) version 1.2 or later or
you can simply use a recent version of [botamusqiue](https://github.com/azlux/botamusique).


### Nicknames

For those of you who are annoyed by people using excessively long or complicated names or people changing their nickname frequently, Mumble 1.4.0
introduces the possibility to set a nickname for every client.

Depending on your configuration, a nickname will either be displayed next to a user's name or replace it entirely.

The nickname will persist across restarts of client and server. In fact it'll last until you either remove it again or the person the nickname was
assigned to changes their client-certificate (Mumble's way of keeping track of a user's identity).

Note that nicknames are _not_ tied to a particular server. That means that if you assign a nickname to a user on server A and then you meet that
person again on server B, they will still have that very same nickname set on server B.

![Nicknames](/blog/media/Nicknames.png)


## Special thanks

We want to thank everyone who spent time on the project. Without this kind of help from you, we would not be able to provide the software the way it is.


## Changelog

This changelog is the product of 1082 commits contributed to the Mumble repository on GitHub.

<details>
	<summary>Changelog</summary>

### Client

- Added: DBus calls to control push-to-talk state ({{<issue 3675 >}})
- Added: Possibility to toggle 24h time-format for the chat ({{<issue 3827 >}})
- Added: UI indication for access-restricted channels (e.g. via password) ({{<issue 3929 >}})
- Added: Ability to use the currently selected user/channel as whisper/shout target ({{<issue 4048 >}})
- Added: ChannelListeners (ability to listen to channels without joining them) ({{<issue 4011 >}})
- Added: Markdown support in chat ({{<issue 4076 >}})
- Added: Ability to set local nicknames for users ({{<issue 4624 >}})
- Added: TalkingUI (~ overlay for non-gamers) ({{<issue 4066 >}})
- Added: DBus calls to set/query transmission mode ({{<issue 4119 >}})
- Added: "Join user's channel" context menu action ({{<issue 4149 >}})
- Added: CLI option to specify custom window title extension ({{<issue 4155 >}})
- Added: Stereo decoding and playback support ({{<issue 4209 >}})
- Added: Ability to disable text-to-speech for a specific user ({{<issue 4287 >}})
- Added: Ability to send images by pasting them to the chat box ({{<issue 4265 >}})
- Added: New talking state for users that are locally muted but sending audio ({{<issue 4322 >}})
- Added: Public server list can be disabled ({{<issue 4316 >}})
- Added: Ability to specify config file via commandline ({{<issue 4369 >}})
- Added: Ability to specify database path in config file ({{<issue 4369 >}})
- Added: Ability to permanently display local volume adjustments in the UI ({{<issue 4439 >}})
- Added: Select all functionality in Messages settings page ({{<issue 4465 >}})
- Added: Ability to paste and send chat message via shortcut ({{<issue 4531 >}})
- Added: Ability to reset all settings at once ({{<issue 4546 >}})
- Added: Shorcut to hide/show the main window ({{<issue 4562 >}})
- Fixed: Handling of protocol violation by the server ({{<issue 3866 >}})
- Fixed: Micro-freezes at startup due to version check ({{<issue 3987 >}})
- Fixed: Distorted positional audio ({{<issue 4172 >}})
- Fixed: Echo cancellation not working properly ({{<issue 4167 >}})
- Fixed: Prevent hooking if a screen-reader is active on Windows ({{<issue 3896 >}})
- Fixed: Crash due to audio buffer resizing ({{<issue 4250 >}})
- Fixed: Dialog for local volume adjustments could be hidden if always-on-top nehavior was active ({{<issue 4244 >}})
- Fixed: Loading a sample from a file would fail silently ({{<issue 4497 >}})
- Fixed: Client now respects the server-setting for unlimited image size ({{<issue 4611 >}})
- Improved: Automatically select Opus's low delay mode for decreased latency ({{<issue 3753 >}})
- Improved: The shortcut dropdown window is now sorted alphabetically ({{<issue 3815 >}})
- Improved: JackAudio support ({{<issue 3826 >}}, {{<issue 3876 >}}, {{<issue 3887 >}})
- Improved: Include Windows-only plugins on Linux as well for use through Proton/Wine ({{<issue 3511 >}})
- Improved: PortAudio support ({{<issue 3889 >}})
- Improved: Use HTTPS links for presenting in the UI ({{<issue 3921 >}})
- Improved: Formatting & spacing of messages in chat ({{<issue 4026 >}})
- Improved: Audio wizard (appearance & explanations) ({{<issue 4100 >}})
- Improved: Echo cancellation settings ({{<issue 4113 >}}, {{<issue 4174 >}})
- Improved: Users can now always choose to receive update-notifications for the client ({{<issue 4138 >}}, {{<issue 4182 >}})
- Improved: Echo cancellation is now enabled by default on all platforms but MacOS ({{<issue 4214 >}})
- Improved: Accessibility ({{<issue 4211 >}}, {{<issue 4312 >}})
- Improved: Some settings pages ({{<issue 4240 >}}, {{<issue 4243 >}})
- Improved: CoreAudio implementation ({{<issue 4254 >}})
- Improved: The user is notified if a change (e.g. local mute) could not be saved permanently ({{<issue 4301 >}})
- Improved: Server Browser UI ({{<issue 4291 >}})
- Improved: Clear selection in shortcut settings after having removed shortcut to avoid accidental subsequent removal ({{<issue 4358 >}})
- Improved: Noise cancelling & associated UI ({{<issue 4212 >}})
- Improved: Use zeroconf instead of Bonjour on Windows if available ({{<issue 4494 >}})
- Improved: Tooltips for shortcut settings ({{<issue 4543 >}})
- Updated: Opus to v1.3.1 ({{<issue 3813 >}})
- Changed: Text-to-speech is now disabled by default ({{<issue 4627 >}})
- Removed: DirectSound support ({{<issue 3828 >}})
- Removed: Qt4 support ({{<issue 3602 >}})
- Removed: CELT 0.11.0 support ({{<issue 2045 >}})

### Server

- Added: Ability to log ACL and group changes (helpful for debugging) ({{<issue 4017 >}})
- Added: Option to not ban connections from an IP that managed to connect succcessfully ({{<issue 4087 >}})
- Added: Dedicated ACL for resetting comments/avatars ({{<issue 4196 >}})
- Added: Config option which allows the server to remeber a user's channel for limited time only ({{<issue 4147 >}})
- Added: Ability to load welcome message from file ({{<issue 4344 >}})
- Fixed: Use a temporary keychain on macOS in order to avoid permission issues ({{<issue 4345 >}})
- Fixed: Undefined behavior when logging SSL error ({{<issue 4452 >}})
- Improved: Also log a client's OS ({{<issue 4035 >}})
- Improved: CPU utilization by using TCP_NODELAY mode by default ({{<issue 4054 >}})
- Improved: Use zeroconf instead of Bonjour on Windows if available ({{<issue 4494 >}})

### Installer

- Improved: Chinese translations ({{<issue 3807 >}}, {{<issue 3613 >}})
- Improved: Uninstall no longer deletes murmur.ini file ({{<issue 4096 >}})
- Complete refactoring ({{<issue 4574 >}})


### Positional audio plugins

- Added: General Source Engine support ({{<issue 3771 >}})
- Added: Among Us ({{<issue 4571 >}})
- Updated: GTA V ({{<issue 4059 >}})
- Updated: Manual plugin's UI now shows a speaker's location ({{<issue 4352 >}})


### Overall

- The project has been migrated to be built with cmake instead of qmake ({{<issue 4252 >}})

</details>

## Known issues

- Overlay blocked by BattleEye. A request to whitelist it has been made.
- Overlay blocked by CS:GO Trusted Mode
