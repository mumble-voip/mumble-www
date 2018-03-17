---
title: Version numbering change
author: .D0T
type: post
date: 2014-02-10T21:19:13+00:00
categories:
  - Release

---
<img class="alignleft size-thumbnail wp-image-541" title="Version numbers" src="http://blog.mumble.info/wp-uploads/2014/02/kablam_Number_Animals-e1391992112887-150x100.png" alt="" width="150" height="100" />Mumble&#8217;s version numbers have been a great source of confusion for a long time. Version 1.1.8 being incompatible with 1.2.0, only incrementing the last place of the version number even though big changes were made and so on. This will change.

**TL;DR:** The next version will be 1.3.0 not 1.2.6. The next feature release after that 1.4.0 and so on.

<!--more-->In practice we used &#8220;

_1.<major>.<minor>_&#8221; in which the _major_ number is incremented for breaking changes and the _minor_ number for feature releases. This wasn&#8217;t actually intended but we were never quite happy with the amount and kind of features we added to actually name a release Mumble 2. Even with the massive and backwards compatibility breaking changes introduced in <a title="1.2.0" href="http://mumble.sourceforge.net/1.2.0" target="_blank">1.2.0</a>.

We do realize this is kinda ridiculous so for future releases we will switch to a &#8220;_<major>.<minor>.<patch>_&#8221; numbering schema. _Major_ and _minor_ have the same meaning as above while _patch_ is used only for security and bugfix releases like <a title="1.2.5" href="http://mumble.sourceforge.net/1.2.5" target="_blank">1.2.5</a>.

This means that the next feature release of Mumble will be <a title="1.3.0" href="http://mumble.sourceforge.net/1.3.0" target="_blank">1.3.0</a> and still will be fully compatible with 1.2.X clients and servers. The next backwards compatibility breaking changes will be in Mumble 2 but there will likely be multiple feature releases before we will feel the need to take that step.

Regards,
  
The Mumble team