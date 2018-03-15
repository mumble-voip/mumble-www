---
title: ‘Mumble for iOS’ 1.2.3 is buggy when running on iOS 7.1
author: mkrautz
type: post
date: 2014-03-10T22:47:03+00:00
url: /?p=553
categories:
  - iOS
  - News

---
<img class="alignleft" src="http://blog.mumble.info/wp-uploads/2012/02/MumbleAppIcon.png" alt="" />A word of warning: &#8216;Mumble for iOS&#8217; 1.2.3 is incredibly unusable on iOS 7.1 &#8211; if you depend on it, you should probably not update to iOS 7.1 just yet.﻿

There is a problem with audio subsystem, which is causing it to endlessly try to restart itself. This manifests itself via iOS&#8217;s &#8216;red bar&#8217; (signalling that an app is using background audio) toggling itself on and off very fast (even when inside the app!), making the app unusable.

We&#8217;ve identified the problem, and we&#8217;re working on a fix. If you&#8217;ve already updated to iOS 7.1, fear not! We&#8217;ve also found a workaround which, while inconvenient, should help you work around the issue until a proper fix has been released.

<!--more-->

To work around the issue, you disable Echo Cancellation in the app. Here&#8217;s how:

1. In Mumble&#8217;s Preferences, navigate to &#8216;Advanced&#8217; under Audio.

2. When in there, look for the &#8216;Audio Input&#8217; section.

3. There should be a checkbox for Echo Cancellation. Disable it.
  
It can be hard when the bug is occurring. If you&#8217;re having trouble, swipe/drag the checkbox&#8217;s button to disable it.

4. Double tap your home button.

5. Swipe upwards on the Mumble app to close it.

6. Re-open Mumble. The issue should be gone.

7. If the issue is still present, navigate to &#8216;Echo Cancellation&#8217; again, and check that it&#8217;s really disabled. It might not work the first time, due to the &#8216;red bar&#8217; making it very hard to actually execute these steps.

We hope it works for you!

Furthermore, we&#8217;d like to note that we hope to be able to submit an updated app with this issue fixed tonight.﻿ This update will also bring some other goodies, so stay tuned.

The Mumble Team