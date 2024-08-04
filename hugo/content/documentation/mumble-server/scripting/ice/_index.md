---
title: Mumble Server Scripting with Ice
date: 2019-10-27
---

The Mumble Server supports remote scripting using the [ZeroC Ice](https://zeroc.com/products/ice) RPC protocol and
framework.

There are bindings (libraries) for C++, Java, .NET, Python, PHP and Ruby, and this is supported on all our platforms
(Win32, Linux and OSX). Ice works locally and also over a network. This means you can create a web application that
interfaces with a Murmur process running on the same or other machines.

Note that after you have ICE set up on your machine, you can install a
{{< wiki "3rd_Party_Applications#Web-Interfaces" "web interface" />}}.

## Developing for the Murmur Ice interface

How to use Ice differs from language to language. The parameters and method names will remain the same, but the syntax
will naturally be different. Murmur will, by default, open up an adapter on port `6502` (or `10000` for homedir
installs), which has a single accessible object named `Meta`. This is the Meta server, and from it you can retrieve
adapters for any configured server. (One server process can run multiple Mumble servers.)

The ice interface is fully documented, and you can browse the [generated documentation](/documentation/slice/).

Further detailed, feature and language specific documentation is available in subpages.

{{< aside >}} This content released under
[Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This
content is based on {{< wiki Ice />}}. {{< /aside >}}
