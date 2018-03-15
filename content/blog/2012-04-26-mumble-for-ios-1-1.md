---
title: Mumble for iOS 1.1
author: mkrautz
type: post
date: 2012-04-26T20:08:44+00:00
excerpt: |
  Today we're releasing version 1.1 of the Mumble iOS app on the App Store.
  
  This release addresses quite a few of the most sought-after features missing from the 1.0 version, along with fixes for various crashes and memory leaks.
url: /?p=373
categories:
  - iOS
  - News
  - Release

---
[<img class="alignleft size-full wp-image-210" title="Mumble for iOS" src="http://blog.mumble.info/wp-uploads/2012/02/MumbleAppIcon.png" alt="Mumble for iOS" />][1]Today we&#8217;re releasing version 1.1 of the Mumble iOS app on the App Store.

This release addresses quite a few of the most sought-after features missing from the 1.0 version, along with fixes for various crashes and memory leaks.

<!--more-->

## **New Features**

  * An improved icon for the app. Thanks Karsten Bruns!
  * Searchable public server list. Thanks Jimmy Selgen Nielsen!
  * Bluetooth headset support.
  * Echo Cancellation support for the built-in mic.
  * Improved Push-to-Talk configurability and discoverability.
  * Self-registration is now supported.
  * Intermediate certificates can now, optionally, be shown in the Certificate Preferences. Intermediate certificates are colored green. Identities keep their yellow color.
  * When exporting a certificate for which a chain can be built, the whole chain will be exported.
  * Improved crash reporting. After a crash, the app will allow you to submit a crash report directly to the Mumble Developers. An &#8216;Always&#8217; option is also available, and we hope users will make use of this feature.

## **Bug Fixes**

  * Various crash bugs have been fixed.
  * Various memory leaks have been plugged. Thanks Jimmy Selgen Nielsen!
  * The 1.0 version of the app failed to parse some Distinguished Names from X.509 certificates. A more solid DN parser has been implemented.
  * When connecting to a server, the client will now attempt to build up a full certificate chain using certificates found in the app&#8217;s keychain.
  * Importing a PKCS #12 file now correctly imports an identity (certificate and private key), along with all intermediate certificates.
  * All action sheets in the app have been converted to black to better fit in with the rest of the app.
  * Transitioning between the Channel view and a modal view, such as Access Tokens, now provides a better visual experience.
  * The mumble.sqlite database is no longer a hidden file in ~/Documents. It now lives in ~/Library.
  * The public server list file is no longer a hidden file in ~/Documents. It now lives in ~/Library.

Mumble for iOS is Open Source software released under the BSD license. We&#8217;re always looking for new contributors, and valuable feedback from the community. We&#8217;d like to thank the people who&#8217;ve reported bugs, and helped beta test the 1.1 release.

Head on over to the app&#8217;s [GitHub repository][2] if you too are interested in contributing!

And don&#8217;t forget &#8212; if you enjoy the app, please spread the word, and remember to rate it in the App Store.

 [1]: http://itunes.apple.com/us/app/mumble/id443472808?mt=8
 [2]: https://github.com/mumble-voip/mumble-iphoneos