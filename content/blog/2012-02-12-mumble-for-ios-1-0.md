---
title: Mumble for iOS 1.0
author: mkrautz
date: 2012-02-12T20:41:27+00:00
categories:
  - iOS
  - News
  - Release

---
Today we&#8217;ve released [Mumble for iOS 1.0 on the App Store][1].

This has been brewing for some time, but we are finally comfortable with calling this release 1.0, from both a feature and a UI perspective. The app is, of course, free of charge and open source.

**UPDATE:** _The initial release was excluded from the App Store in France. We&#8217;ve since fixed this error, and apologize for any inconvenience this has caused!_

_<!--more-->_

## Availability

Mumble for iOS 1.0 runs on iPhone 3Gs, iPhone 4, iPhone 4s, iPod Touch 3rd Gen, iPod Touch 4th Gen, iPad and iPad 2. It requires iOS 5 or greater to run.

The application is currently built for the phone form factor, and work is still ongoing to create a gorgeous iPad UI for it.

## Features

The 1.0 version of the application includes the functionality that is needed to get a typical Mumble user up and running on an iOS device. However, the client is not at feature parity with its desktop sibling just yet. Current features include:

  * Voice-Activity, Push-To-Talk and Continuous transmission modes.
  * Ability to turn off audio preprocessing to save power (and reduce UI lag). This feature is provided mainly for users of iPhone 3Gs and iPod Touch 3rd Gen, where preprocessing is more taxing on the CPU.
  * Codec presets (Low, Balanced and High), for selecting your Transmission Quality.
  * Force TCP-mode for networks where that is required.
  * Support for creating self-signed certificates on-device and using them to connect to a server.
  * Support for importing certificates into the app from iTunes. Export to iTunes also possible.
  * Support for connecting to passworded servers (using either a user password or a server password).
  * Access token support for &#8216;password protected&#8217; channels, and other access token trickeries.
  * The same world-wide public server list that is available in the desktop Mumble client.
  * Excellent favourite server support.
  * Support for connecting to LAN servers discovered via Bonjour.
  * Ability to join servers by opening mumble:// URLs
  * Separate server and channel-only views for easier use on small screens. (Push-To-Talk only available in the channel view.)
  * Support for self-muting and self-deafening.
  * Support for sending and receiving text messages.

(A similar list of features, but tailored towards the casual AppStore app-cruising user is available [here][1])

We have tried hard to include the features that a regular, non-administrator user of Mumble would use on a day-to-day basis, and we think we have mostly succeeded in that. However, not everything we would like to have available at launch is there. These missing features are:

  * Self-registration
  * Local muting other people
  * Viewing channel comments and setting and viewing user comments

We&#8217;re aware that these features are nice-to-haves, and will work on making these available as soon as we can.

## The Future

For the immediate future, we are working on implementing the above list of missing features, and fixing the bugs and crashes that are found in the wild by users of the 1.0 version of the application.

As mentioned earlier in this blog post, we also plan to implement a UI for the iPad. Along with this, we hope to add many of the administrative features found in the desktop Mumble client. The improved screen real-estate available on the iPad makes the implementation of these features preferable to do on the iPad first.

## Feature Requests and Bug Reports

If you stumble upon a bug in the client, do not hesitate to report it to us on the mumble-iphoneos GitHub issues page:

<https://github.com/mumble-voip/mumble-iphoneos/issues>

Feature requests are welcome on there, as well.

## Concluding Remarks

We hope that you will enjoy using the app as much as we have enjoyed creating it. If you like it, _do not forget to rate it in the App Store_.

The app, like the Mumble desktop client, is Open Source, available under a BSD-style license.

The source code is available on our GitHub page:

<https://github.com/mumble-voip/mumble-iphoneos>
  
[https://github.com/mumble-voip/mumblekit][2]

Contributions of any kind are always warmly welcomed.

Regards,
  
The Mumble Team

 [1]: http://itunes.apple.com/us/app/mumble/id443472808?mt=8
 [2]: https://github.com/mumble-voip/mumble-iphoneos
