---
title: Mumble for iOS 1.2.3
author: mkrautz
date: 2014-02-05T15:00:12+00:00
categories:
  - iOS
  - News
  - Release
  - Security

---
<img class="alignleft" src="http://blog.mumble.info/wp-uploads/2012/02/MumbleAppIcon.png" alt="" />The Mumble team has released [version 1.2.3 of the iOS Mumble client][1].

This new version contains two important client-side security fixes, and we advise users to download this update from the App Store as soon as possible.

<!--more-->

Alongside these security fixes, this release also includes a few minor bug fixes:

  * Increased the size of Mumble&#8217;s encoding buffer for Opus packets
  
    to be able to encode all Opus packets without triggering Opus&#8217;s
  
    internal rate limiting.
  * Fixed a bug that could cause the certificate accept alert view that is shown
  
    upon connecting to a server with an unknown certificate to sometimes be hidden
  
    on iOS 7.
  * Several external libraries have been synced to their latest stable versions.

Security advisories for the two fixed vulnerabilities are available below:

[Mumble-SA-2014-003][2] [[sig][3]]
  
&#8211; A malformed Opus voice packet sent to a MumbleKit client could trigger a NULL pointer dereference.

[Mumble-SA-2014-004][4] [[sig][5]]
  
&#8211; A malformed Opus voice packet sent to a MumbleKit client could trigger a heap-based buffer overflow.

The Mumble team

 [1]: http://itunes.apple.com/us/app/mumble/id443472808?mt=8
 [2]: http://mumble.info/security/Mumble-SA-2014-003.txt
 [3]: http://mumble.info/security/Mumble-SA-2014-003.txt.sig
 [4]: http://mumble.info/security/Mumble-SA-2014-004.txt
 [5]: http://mumble.info/security/Mumble-SA-2014-004.txt.sig