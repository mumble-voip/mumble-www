---
title: New versioning scheme
author: Robert Adam
date: 2021-12-07
categories:
  - News
---
Mumble traditionally used a versioning scheme of the form _major.minor.patch_ where the patch version number was incremented for every bug-fix release
of a stable release series (e.g. 1.3.0 -> 1.3.1). Most notably though, the patch version was always starting at zero for a new stable release series
(e.g. 1.2.0, 1.3.0, etc.). This will no longer be the case.

<!--more-->

## Background

Mumble was always using the [semantic versioning scheme](https://semver.org/) which generally makes it really easy to know what to expect when
switching from one version to another. While this is working great, using only the three common version numbers (_major.minor.patch_) poses a problem
when considering snapshot builds. A snapshot build is essentially a beta-release for an upcoming release. Thus it has to use the same version numbers
as that upcoming release. Thus, in order to get a unique versioning scheme, we had to make use of the fourth version number, that semantic versioning
allows for. We called this the _build number_. This build number was just a number that was incremented with every build of the Mumble binaries that
we performed. Thus a version could now look like this: 1.4.0.62 and the stable release would then simply use whatever build number was next in the
queue.

## The problem

In theory that worked great and solved the issue at hand: we now had a unique version for every release (be that a stable or a snapshot release).
However, as it turns out Windows doesn't like that: it restricts programs to use a three-component versioning scheme. While it is possible to set a
fourth component, without causing any issues (it is even displayed when checking the program's version), it is ignored for all relevant actions such
as upgrading to a newer version. However, it seems it wasn't ignored entirely as using a four-component versioning scheme caused Windows to install an
additional version of Mumble alongside the already installed version if one was to attempt to upgrade between Mumble versions that only differed in
the build number.

## The solution

After having researched quite a bit, we came to the conclusion that it seems to be impossible to get Windows to properly respect the build number in
our version specifier (old dogs don't learn new tricks), so it was clear that we had to find a different solution to this. We evaluated a few
different possibilities but eventually settled on using the patch version number as a replacement for the build number. Consequently, the latter will
be dropped.  Thus, we are back at a 3-component version scheme but now the third component is the build number instead of the patch level
(_major.minor.build_).

This appears to be a purely semantical difference and for the most part, it is. The only notable difference is that now the third component of the
version will (usually) no longer start at zero and then increment one by one while iterating through the different versions of a stable release
series. Instead, the build number will be increased with every build we create and thus the build number of the released stable versions will be
somewhat arbitrary. Therefore, it could be that the first release of the 1.4 series will be e.g. 1.4.53 while the first bug-fix release will then
become 1.4.281.

This makes it impossible to tell (by just looking at the versions themselves) how many releases lay between two given version numbers. However, in
most scenarios this is probably not what is interesting anyway. Usually one is only interested in which version is newer and that is still immediately
obvious by looking at the version numbers.

## Summary

Versions remain in a three component format (and also compatible with the semantic versioning scheme) but the third component will now be the build
number instead of the patch level. The latter is no longer directly visible from the version specification (and in fact isn't really tracked in its
own right anymore). Therefore, you will start seeing much higher (and apparently discontinuous) numbers in the third component of future Mumble
versions.
