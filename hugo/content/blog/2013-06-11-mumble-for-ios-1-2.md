---
title: Mumble for iOS 1.2
author: mkrautz
date: 2013-06-11T22:03:56+00:00
categories:
  - Release
tags:
  - ios
---

Earlier today, we released [Mumble for iOS 1.2][1].

This release is compatible with the 1.2.4 protocol and includes support for the Opus codec.

It also fixes many of the most irritating bugs people have found and reported with the previous release.

<!--more-->

As usual, the source code is availble on GitHub ([Mumble for iOS][2], [MumbleKit][3]), and contributions of any kind are
warmly appreciated.

The full list of changes in Mumble for iOS 1.2:

Protocol:

• Fully protocol compatible with 1.2.4.

Devices:

• The UI is now tuned for 4 inch retina devices (iPhone 5 and 5th gen iPod Touch).

Audio:

• The Opus codec is now available for use in Mumble for iOS.

• Echo Cancellation is now re-enabled for all iOS 6 devices.

• An Audio Sidetone option is now available.

Certificates:

• Certificate information is now copyable.

• SHA1 and SHA256 fingerprints are now shown in the certificate UI.

User Interface:

• The channel-only view mode tab has been removed.

• Clicking the Mumble icon when connected to a server will now toggle between showing all channels, and the one you are
in.

Bug fixes:

• The bug that caused the &#8216;red bar&#8217; to appear randomly when using Mumble along with Bluetooth accessories,
Siri, or other audio apps has been fixed.

• Various minor bug fixes.

As always, it’s available on the [App Store][1]. And don’t forget — if you enjoy the app, please spread the word, and
remember to rate it in the App Store.

[1]: https://itunes.apple.com/us/app/mumble/id443472808?mt=8
[2]: https://github.com/mumble-voip/mumble-iphoneos
[3]: https://github.com/mumble-voip/mumblekit
