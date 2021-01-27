---
title: Mumble Security Advisory 2014-006
date: 2014-05-14
---

ID:              Mumble-SA-2014-006 ([txt](../Mumble-SA-2014-005.txt), [sig](../Mumble-SA-2014-005.txt.sig), [patch](../Mumble-SA-2014-005.patch), [patch sig](../Mumble-SA-2014-005.patch.sig), [Blog post](/news/mumble-1.2.6/))

Date:            May 2014

CVE Reference:   Not assigned

Product:         Mumble

Mumble Website:  http://mumble.info

Permalink:       http://mumble.info/security/mumble-sa-2014-006

Last Updated:    2014-05-14

## Vulnerability

The Mumble client did not properly HTML-escape some external strings before using them in a rich-text (HTML) context.

In some situations, this could be abused to perform a Denial of Service attack (hang and/or resource exhaustion) on a Mumble client by causing it to load external files via the HTML.

## Affected versions and configurations

All 1.2.x versions of Mumble are affected.

## Mitigation

No known mitigation strategies exist for this issue.

## Details

By default, many Qt widgets sniff their text content to determine whether or not to use to render the text as HTML. However, in some places, the Mumble client neglected to properly escape external strings when used rich-text enabled Qt widgets.

One case where this was an issue was when including user and channel names in Qt dialogs. However, this vector is automatically mitigated if using the default channel- and username regular expressions, since angle brackets are disallowed by it.

Another case where this issue could be triggered was in the Certificate Wizard, where the subject common name and email addresses were not properly escaped.

A similar problem exists with the tooltip for servers in the Mumble server list, where the server name was not properly escaped.

While finding and fixing the above issues, any other instances where an explicit HTML-escape was deemed appropriate had one inserted.

In some cases, instead of performing explicit HTML escapes, Qt widgets that were determined not to ever contain HTML content were changed to only be able to render plain text.

## Credits

This issue was originally fixed in the Mumble git repo by various commits from Tim Cooper and Mikkel Krautz.

## Fix

A fix for this issue has been released in Mumble 1.2.6.

A patch that applies on top of Mumble 1.2.5 can be found at http://mumble.info/security/Mumble-SA-2014-006.patch
	(for a detached PGP signature, append '.sig').

The fix is also available in the Mumble git repo as e30d7acda6c04b667618ac86f49786cf966a08fb (on the v1.2.6 branch).
