---
title: Contribute
toc: true
---

Mumble is a great program, but that is no reason to stop polishing it. There are numerous areas where the development team needs help, and if you feel like contributing, then any of the below is for you.

# Find and report Bugs, suggest features

Use the developer snapshots of Mumble and Murmur (see [Downloads]({{< relref "downloads" >}})) and see if you can find any issues with them. If you do, check if the issue has already been reported, add details if you have more, or add a new tracker item if it has not been reported yet. The issue tracker can be found at https://github.com/mumble-voip/mumble/issues.

You may also want to suggest new features at https://github.com/mumble-voip/mumble/issues. 

# Translation

Mumble currently (March 2014) supports 18 different languages. As more people get to know and use Mumble versions in their native languages become more and more important.

We are thankful for any contribution on translating. Please check the [Language Translation wiki page](https://wiki.mumble.info/wiki/Language_Translation) for more information.

# Documentation

Mumble and all around it needs documentation. Great documentation for a great program.

**Website**: This website has a public hugo source repository at [mumble-www](https://github.com/mumble-voip/mumble-www). You can report [issues](https://github.com/mumble-voip/mumble-www/issues), create [requests](https://github.com/mumble-voip/mumble-www/issues), and create pull requests for implemented change suggestions.

**Wiki**: Our [documentation wiki] covers some more advanced topics than this website; for example for server administrators and contributors. Any contributions welcome, be it improving existing pages or creating new ones with useful content. Also see the [Please Improve Category](https://wiki.mumble.info/wiki/Category:Please_Improve).

**Wikipedia**: Someone that is not a development team member needs to maintain the [article at Wikipedia](https://en.wikipedia.org/wiki/Mumble_%28software%29). At the moment it’s missing references; try to find some. (This requires more work than just googling). Contributing to or creating a mumble article on non-English Wikipedia are also appreciated. See Promotion section below.

**Ice methods**: We have a [generated documentation](http://mumble.sourceforge.net/slice/) for exposed Ice methods. Although probably only developers using the doc will be able to understand it, please give us a note on any inconsistencies, errors or needed clarifications if you find them.

# Promotion

Get the word out. This means more than writing about it on your blog ;)

* Ask (voice-)server hosters on if they also provide Mumble servers if you can’t find them.
* Coordinate with one of the hosters providing free mumble servers and start promoting Mumble as the voicechat of choice for your favorite game.
* Add mumble hosters you found to the [Hosters](https://wiki.mumble.info/wiki/Hosters) wiki page / update information.
* Create How-Tos, Screencasts, Tutorials or Benchmarks on typical issues and topics, and publish them on prominent platforms.
* Spread the word in groups – join our [Steam Group](https://steamcommunity.com/groups/Mumble_VoIP)

# Skinning

Mumble skins bring variety into Mumbles design and choices for different tastes.

Create new skins for Mumble! Information can be found on our [Skinning wiki page](https://wiki.mumble.info/wiki/Skinning) and existing skins we know about on the [Skins wiki page](https://wiki.mumble.info/wiki/Skins).

# Coding

## 3rd party Interfaces

Web interfaces, web-viewers, administration programs, CMS website modules – all neat things we would love to see! These could be written in any language supported by our Ice based RPC Interface (using the deprecated [DBus](https://wiki.mumble.info/wiki/DBus) Interface is no longer advised). You could start your own project or link up with an [already existing one](https://wiki.mumble.info/wiki/3rd_Party_Applications). To get an idea about how interaction with Murmur might look like you can take a look at the example scripts in our git repositories script folder.

## Ice authenticators

See [Authenticators](https://wiki.mumble.info/wiki/Authenticators). We’re happy for every authenticator and authenticator improvement you and the community can bring us!

## Positional Audio

Most of our positional audio plugins regularly break on game updates. To read on how you can help us update the plugins see the [Pluginguide](https://wiki.mumble.info/wiki/Pluginguide) page. You can even fix the update-issue forever for open source games, or if you want to add positional audio to your game, see [Link](https://wiki.mumble.info/wiki/Link).

## Mumble itself

We are always looking for helping hands willing to contribute to mumble. You'll need a good understanding of C++ and previous knowledge in Qt is definitely helpful. The best way to get into mumble development is to link up with us on [IRC](https://wiki.mumble.info/wiki/IRC) ([#mumble on freenode](irc://chat.freenode.net/mumble)), there is always small stuff to do to get you started. Also check the [Development wiki page](https://wiki.mumble.info/wiki/Development) on more info about the source code and building it.
