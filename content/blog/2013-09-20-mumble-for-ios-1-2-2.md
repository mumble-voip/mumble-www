---
title: Mumble for iOS 1.2.2
author: mkrautz
date: 2013-09-20T10:23:32+00:00
categories:
  - iOS
  - News
  - Release

---
This Wednesday, to coincide with the release of iOS 7, we released [Mumble for iOS 1.2.2][1]. This is a primarily a release to fix incompatibilities with iOS 7, but it also includes a few other oft-requested goodies.

<!--more-->

Running the previous version of Mumble for iOS (1.2.1) on iPhone 4S or iPhone 5 with iOS 7, the app could sometimes get into a state where no audio would be recorded. This is the primary compatibility issue that this new 1.2.2 release addresses, and we recommend you upgrade as soon as you get the chance.

We&#8217;ve also added a speakerphone toggle in the app&#8217;s preferences. This allows you to change whether Mumble will output to your iPhone&#8217;s speaker or its receiver, allowing you to use it more like a regular mobile phone if you please.

As usual, the source code for the app is available on GitHub ([Mumble for iOS][2], [MumbleKit][3]), and contributions of any kind are warmly appreciated.

The full list of changes in Mumble for iOS 1.2.2:

• Fixed a bug where audio would be silent on launch or after consecutive visits to the app&#8217;s preferences on iOS 7.
  
• Fixed a bug which prevented some certificates from importing correctly.
  
• Speakerphone mode is now optional (a preference).
  
• Mumble now forces CELT mode when using Opus, for performance reasons. This can be disabled in the app&#8217;s preferences.

As always, it’s available on the [App Store][1]. And don’t forget — if you enjoy the app, please spread the word, and remember to rate it in the App Store.

 [1]: http://itunes.apple.com/us/app/mumble/id443472808?mt=8
 [2]: https://github.com/mumble-voip/mumble-iphoneos
 [3]: https://github.com/mumble-voip/mumblekit
