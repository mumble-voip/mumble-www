---
title: Mumble for iOS 1.2.1
author: mkrautz
date: 2013-06-20T21:57:41+00:00
categories:
  - Release
tags:
  - ios

---
Today, we&#8217;ll be releasing [Mumble for iOS 1.2.1][1]. This is a minor bug fix release to the 1.2 release we did earlier this month.

We found a nasty race condition which could cause crashes when the audio subsystem of the app was started or restarted. This bug seemed to happen very often on 4th generation iPod Touch devices, but could happen on all iOS devices with some bad luck.

<!--more-->

We also fixed an issue where users using the &#8220;Low&#8221; quality preset would be unable to speak to others when on a 1.2.4 server in Opus mode.  To allow the &#8220;Low&#8221; quality to mode

to work correctly on all servers, we&#8217;ve changed it to use Opus in Opus mode, and CELT 0.7 (our baseline codec) on everything else. Previously, we always used Speex for it.

As usual, the source code for the app is available on GitHub ([Mumble for iOS][2], [MumbleKit][3]), and contributions of any kind are warmly appreciated.

The full list of changes in Mumble for iOS 1.2.1:

• Fixed a bug where users using the &#8216;Low&#8217; audio quality mode could not send audio on 1.2.4 servers.

• Fixed a bug where chat messages sometimes had incorrect timestamps.

• Fixed an audio recording bug that would sometimes crash the app.

• Fixed several crash bugs which were triggered when using the the &#8216;Low&#8217; audio quality mode.

As always, it’s available on the [App Store][1]. And don’t forget — if you enjoy the app, please spread the word, and remember to rate it in the App Store.

 [1]: https://itunes.apple.com/us/app/mumble/id443472808?mt=8
 [2]: https://github.com/mumble-voip/mumble-iphoneos
 [3]: https://github.com/mumble-voip/mumblekit
