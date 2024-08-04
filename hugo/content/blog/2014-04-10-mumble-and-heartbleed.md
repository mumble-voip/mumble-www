---
title: Mumble and Heartbleed
author: mkrautz
date: 2014-04-10T23:21:53+00:00
categories:
  - Security
  - Tech talk
---

Most of the Mumble project&#8217;s communication this week regarding the [Heartbleed][1] vulnerability has happened via
IRC. This blog post attempts to fix that by providing answers to the most frequently asked questions here on our blog.

<!--more-->

If you&#8217;re reading this blog post, you&#8217;re probably wondering whether your Mumble client or Murmur server is
vulnerable to the [Heartbleed][1] ([CVE-2014-0160][2]) vulnerability in [OpenSSL][3].

The answer is: &#8220;It depends.&#8221;

The binary Mumble and Murmur packages that are available to download from [SourceForge][4] and [mumble.info][5] are not
affected. These packages use [OpenSSL][3] 1.0.0, and as such are not vulnerable to [Heartbleed][1]. (This is also why
you have not seen any new releases from us to fix this issue.)

So, if you&#8217;re using a Mumble client/server on Windows or Mac OS X that you downloaded from [SourceForge][4] or
[mumble.info][5], you&#8217;re not vulnerable. If you&#8217;re running the &#8216;static&#8217; Linux server, you are
not vulnerable either. If you&#8217;re running the iOS client, you&#8217;re also good.

However, if you are running Mumble client or a Murmur server that you didn&#8217;t download from [SourceForge][4] or
[mumble.info][5], you are most likely vulnerable. This includes Mumble and Murmur packages from Linux and other
Unix-like systems&#8217;s package managers, and importantly also the Ubuntu PPA archives that we link to from the front
page of [mumble.info][5].

If you&#8217;re on a Unix-like system, you should ensure that your [OpenSSL][3] package is up-to-date and that it
includes a fix for [Heartbleed][1]. Once that is the case, you are no longer vulnerable. (Make sure you restart your
server instances after updating [OpenSSL][3] for the update to have any effect.)

Once you have patched [OpenSSL][3] on any vulnerable systems, you should also strongly consider to revoke and reissue
any certificates, private keys and passwords that have been used by either Mumble or Murmur on the vulnerable machine,
as these might have been readable by attackers.

The Mumble Team

[1]: http://heartbleed.com/
[2]: https://www.cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2014-0160
[3]: https://www.openssl.org/
[4]: https://sourceforge.net/projects/mumble/
[5]: https://mumble.info
