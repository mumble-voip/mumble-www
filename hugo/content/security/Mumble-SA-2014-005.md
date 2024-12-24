---
title: Mumble Security Advisory 2014-005
date: 2014-05-14
---

ID:              Mumble-SA-2014-005 ([txt](../Mumble-SA-2014-005.txt), [sig](../Mumble-SA-2014-005.txt.sig), [patch](../Mumble-SA-2014-005.patch), [patch sig](../Mumble-SA-2014-005.patch.sig), [Blog post](/blog/mumble-1.2.6/))

Date:            May 2014

CVE Reference:   Not assigned

Product:         Mumble

Mumble Website:  http://mumble.info

Permalink:       http://mumble.info/security/mumble-sa-2014-005

Last Updated:    2014-05-14

## Vulnerability

The Mumble client is vulnerable to a Denial of Service attack when rendering crafted SVG files that contain references to files on the local computer.

This can cause the Mumble client to hang and/or become resource exhausted.

This issue can be triggered remotely by an entity participating in a Mumble voice chat. Triggering the issue remotely can be done using text messages, channel comments, user comments and user textures/avatars.

## Affected versions and configurations

All 1.2.x versions of Mumble are affected.

## Mitigation

This issue can be mitigated by using a build of Qt that does not include the QSvg module.

However, since Mumble uses SVGs for some of its UI graphics, this might not be a satisfactory solution.

## Details

Qt's QSvg module's SVG renderer will follow file references found in SVG's image tag and XML stylesheet references.

For image tags, Qt tries to load the referenced file using QImage's constructor that takes a file path. Further processing is then delegated to an image format plugin.

For XML stylsheets, Qt will attempt to open the referenced file using QFile followed by a call to the readAll(), which will read the file until EOF.

These two possibilities makes it easy to cause a Mumble client to hang by reading a file or file-like object that is exposed using the filesystem namespace that QFile can access.

## Credits

This issue was reported to the Mumble team by Tim Cooper on April 16, 2014.

## Fix

A fix for this issue has been released in Mumble 1.2.6.

The fix for Mumble 1.2.6 addresses this issue by removing the Mumble client's ability to render SVG files received from external sources. The binary releases of the Mumble client for Windows and OS X also contain a patched version of Qt that removes the code paths in QSvg that can trigger local file access.

A patch that applies on top of Mumble 1.2.5 can be found at http://mumble.info/security/Mumble-SA-2014-005.patch (for a detached PGP signature, append '.sig').

The fix is also available in the Mumble git repo as c7aecb2956f214cd83b7a862a4fcf15cc76c4450 (on the v1.2.6 branch).

The patch has been applied to Mumble's Qt tree is available in the mumble-developers-qt tre as
	2147fa767980fe27a14f018b1528dbf880b96814.
