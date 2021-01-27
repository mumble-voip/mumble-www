---
title: First Mumble Tracker Squash Meeting
author: slicer
date: 2010-07-21T23:01:44+00:00
categories:
  - Uncategorized

---
Over the past few months, most of the Mumble team has been busy with other projects. This has lead to an unfortunate build-up of unsolved bug and feature requests. This Sunday, we had our first developer meeting to address this.

<!--more-->Every now and then, real life and other projects have to take priority over Mumble development. Since about mid-march, the only requests we&#8217;ve really done anything about has been patches. Bug reports and feature requests have been put on the &#8220;todo&#8221; list, which by now has become pretty long.

In an effort to address this, this Sunday saw our first developer meeting (held using Mumble, of course). This saw quite a few results, which I&#8217;ll go through in turn.

First of all, 1.2.3 is now more or less feature komplete. Krejler has a block of OSX interactive overlay code that is missing, and .D0T, pcgod and other contributors are working hard on finishing an initial version of audio recording. The current plan is to get these features in, resolve any bugs that remain and then release 1.2.3. With a bit of luck that will happen inside of a month.

As a consequence of the above, that means the 200 feature requests we currently have unresolved in the tracker will not be addressed for 1.2.3. Once 1.2.3 is out, we&#8217;ll start going through the requests and implement things that look fun or make a lot of sense. In the meantime, we&#8217;re currently evaluating replacements for the feature request tracker itself. While the tracker works well for bugs, it lacks a few features we&#8217;d like to see for feature requests, such as popularity tracking, milestone targets and easy marking of duplicates. SF.net is currently implementing [SF.net 2.0][1] and maybe that will improve things. If not, we&#8217;ll probably need an alternative.

Most of the meeting was spent triaging bug reports. All bug reports should now either have been assigned to a developer for fixing, set pending with request for more information, or left open because triaging them takes a lot longer. As it turns out, this triaging is a good thing, as it allows all of us to work independently on fixing specific bugs, which means we should have a fairly bug-free snapshot ready inside of a week.

 [1]: https://sourceforge.net/blog/get-ready-for-a-whole-new-forge/
