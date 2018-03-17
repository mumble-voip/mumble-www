---
title: 'Progress Report: April 2017'
author: Kissaki
type: post
date: 2017-05-03T10:36:33+00:00
categories:
  - News
tags:
  - mumble
  - progress

---
This month, we continued work on features for 1.3, some bug fixes and code improvements. We also improved our wiki documentation and categorization and in small parts our README, INSTALL and manual texts.<!--more-->

We did some fixes and improvements on the overlay exceptions system (supporting whitelisting of programs started from launchers).

After restructuring our unit tests for easier testing we implemented more of them for new code we merged.

### New Features

  * Introduced option to hide OS information from the server you connect to PR [#3009][1] [#3015][2]

### Mumble Client

  * Fixed issue with default local volume dialog size PR [#3007][3]
  * Fixed tooltips that could show when the Mumble window was not active PR [#3010][4]
  * Updated translations from Transifex (our translation platform) PR [#3044][5] [#3016][6]
  * Fixed a wizard issue by switching to the &#8220;classic&#8221; style PR [#3020][7]
  * Fix an issue of the Mumble window no longer being visible in some cases with hide in tray enabled and specific Unix window managers (no tray icons) PR [#3025][8]
  * Dropped overlay not being attached if not being injected because of issues in specific cases
  * Added the itch.io launcher and Daybreak Games&#8217; launcher as known launchers to overlay exceptions PR [#3049][9] [#3046][10]
  * Allow drag-dropping files and folders into the overlay exception lists (configuration) PR [#3001][11]

### Positional Audio Plugins

  * Updated our Battlefield 1 Positional Audio Plugin PR [#3003][12]
  * Updated our Battlefield 4 Positional Audio Plugin PR [#2999][13]
  * Updated our Battlefield 2142 Positional Audio Plugin PR [#3040][14]

### Mumble Server

  * Check for validity of configured database driver PR [#3039][15]

### Documentation

  * Cleaned up wiki documentation and categorization
  * Updated documentation on [how to use (free) Let&#8217;s Encrypt server certificates][16]
  * Describe SIGHUP and SIGUSR1 in the manual file (SIGHUP exists for quite some time, SIGUSR1 is a more recent new feature) PR [#3050][17] 
      * SIGHUP for log rotation
      * SIGUSR1 for reloading TLS settings (e.g. reloading Let&#8217;s Encrypt server certificates without downtime)

### Miscellaneous

  * Improved code quality and potential issues (identified issues with CppCheck and PVS-Studio) PR# [#3051][18] [#3054][19] [#3008][20] #
  * Some more work on our public builds (no-pch builds) (travis-ci and appveyor) PR [#3014][21] [#3013][22] [#3012][23] [#3011][24]
  * Introduce build flag/configuration CONFIG(dpkg-buildflags) PR [#3033][25] [#3034][26]

And as always, apart from these features that landed in our master development branch, we had discussions, community/user communication and work on features and improvements that are still work in progress. Our main goal continues to be to bring version 1.3 to a release candidate state.  You can already use 1.3, with most of the features mentioned above, when you install our development snapshots available from our [homepage][27].

 [1]: https://github.com/mumble-voip/mumble/pull/3009
 [2]: https://github.com/mumble-voip/mumble/pull/3015
 [3]: https://github.com/mumble-voip/mumble/pull/3007
 [4]: https://github.com/mumble-voip/mumble/pull/3010
 [5]: https://github.com/mumble-voip/mumble/pull/3045
 [6]: https://github.com/mumble-voip/mumble/pull/3016
 [7]: https://github.com/mumble-voip/mumble/pull/3020
 [8]: https://github.com/mumble-voip/mumble/pull/3025
 [9]: https://github.com/mumble-voip/mumble/pull/3049
 [10]: https://github.com/mumble-voip/mumble/pull/3046
 [11]: https://github.com/mumble-voip/mumble/pull/3001
 [12]: https://github.com/mumble-voip/mumble/pull/3003
 [13]: https://github.com/mumble-voip/mumble/pull/2999
 [14]: https://github.com/mumble-voip/mumble/pull/3040
 [15]: https://github.com/mumble-voip/mumble/pull/3039
 [16]: http://wiki.mumble.info/wiki/Obtaining_a_Let%27s_Encrypt_Murmur_Certificate
 [17]: https://github.com/mumble-voip/mumble/pull/3050
 [18]: https://github.com/mumble-voip/mumble/pull/3051
 [19]: https://github.com/mumble-voip/mumble/pull/3054
 [20]: https://github.com/mumble-voip/mumble/pull/3008
 [21]: https://github.com/mumble-voip/mumble/pull/3014
 [22]: https://github.com/mumble-voip/mumble/pull/3013
 [23]: https://github.com/mumble-voip/mumble/pull/3012
 [24]: https://github.com/mumble-voip/mumble/pull/3011
 [25]: https://github.com/mumble-voip/mumble/pull/3033
 [26]: https://github.com/mumble-voip/mumble/pull/3034
 [27]: http://mumble.info/