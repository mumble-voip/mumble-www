---
title: Lost in translation
author: slicer
type: post
date: 2010-10-05T17:21:47+00:00
categories:
  - Community
tags:
  - translation

---
[<img class="alignleft size-thumbnail wp-image-198" title="spelling" src="http://blog.mumble.info/wp-uploads/2010/10/spelling-150x150.jpg" alt="" width="150" height="150" />][1]One important milestone for any Mumble release is freezing the text-strings. After a freeze, no changes are allowed, which means our translators can start working. For this release, we&#8217;d like a bit of help proofing and checking the strings before we freeze them.

<!--more-->The usual method for freezing strings is to export them all to a webpage, load that webpage into Word and use the spelling and grammar checker. This catches the worst mistakes, and there always are quite a few. However, it does not catch nuances of language, and it certainly doesn&#8217;t catch wrong information.

As I was looking through the strings for this release, I notice that the tooltip for the bitrate in the audiostats mention that you can adjust the bitrate by changing complexity in the audio settings. We replaced the complexity slider with a direct bitrate slider before 1.2.0, so that text has been wrong for quite a few releases. Thankfully, it appears nobody has noticed.

Translations are one of the last hurdles before 1.2.3 ships, and since translations can&#8217;t start until strings are frozen, we need a bit of help. The dev team wrote these strings, and we&#8217;ve seen them so many times we aren&#8217;t really reading them anymore. It&#8217;s therefore hard for us to spot mistakes and errors.

So, if you find any spelling, grammar or textual logic errors in the current snapshot, or in the [exported strings][2], please let us know 🙂

**UPDATE:** Strings are frozen now. Thanks for your help.

 [1]: http://blog.mumble.info/wp-uploads/2010/10/spelling.jpg
 [2]: http://mumble.info/en-ac2b6ca497387652bfe96ef75d444ac1cb622d76.html "Exported strings"