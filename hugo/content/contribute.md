---
title: Contribute
toc: true
---

Mumble is a great program, but that is no reason to stop polishing it. There are numerous areas where the development team needs help, and if you feel like contributing, then any of the below is for you.

# Find and report Bugs, suggest features

Use the developer snapshots of Mumble and Murmur (see [Downloads]({{< relref "downloads" >}})) and see if you can find any issues with them. If you do, check if the issue has already been reported, add details if you have more, or add a new tracker item if it has not been reported yet. The issue tracker can be found at https://github.com/mumble-voip/mumble/issues.

The issue tracker is also used to track [feature requests, which are marked with the `feature-request` label](https://github.com/mumble-voip/mumble/issues?q=is%3Aopen+is%3Aissue+label%3Afeature-request+sort%3Areactions-%2B1-desc). Please vote with a thumbs up on those you would like to see being fulfilled, describe your use cases or otherwise participate in the discussion on existing tickets. And if you want to request a feature not suggested yet, please create a new feature request ticket possibly with some elaboration and use cases.

# Translation

Mumble currently (September 2019) supports over 20 different languages (41 if including unfinished translations). As more people get to know and use Mumble, versions in their native languages become more and more important.

We are thankful for any contribution on translating. Please check the {{< wiki "Language Translation" "Language Translation wiki page" />}} for more information.

# Documentation

Mumble and all around it needs documentation. Great documentation for a great program.

**Website**: This website has a public hugo source repository at [mumble-www](https://github.com/mumble-voip/mumble-www). You can report [issues](https://github.com/mumble-voip/mumble-www/issues), create [requests](https://github.com/mumble-voip/mumble-www/issues), and create pull requests for implemented change suggestions.

**Wiki**: Our [documentation wiki] covers some more advanced topics than this website; for example for server administrators and contributors. Any contributions welcome, be it improving existing pages or creating new ones with useful content. Also see the {{< wiki "Category:Please_Improve" "Please Improve Category" />}}.

**Wikipedia**: Someone that is not a development team member needs to maintain the [article at Wikipedia](https://en.wikipedia.org/wiki/Mumble_%28software%29). At the moment it’s missing references; try to find some. (This requires more work than just googling). Contributing to or creating a mumble article on non-English Wikipedia are also appreciated. See Promotion section below.

**Ice methods**: We have a [generated documentation](http://mumble.sourceforge.net/slice/) for exposed Ice methods. Although probably only developers using the doc will be able to understand it, please give us a note on any inconsistencies, errors or needed clarifications if you find them.

# Promotion

Get the word out. This means more than writing about it on your blog ;)

* Ask (voice-)server hosters on if they also provide Mumble servers if you can’t find them.
* Coordinate with one of the hosters providing free mumble servers and start promoting Mumble as the voicechat of choice for your favorite game.
* Add mumble hosters you found to the {{< wiki Hosters />}} wiki page / update information.
* Create How-Tos, Screencasts, Tutorials or Benchmarks on typical issues and topics, and publish them on prominent platforms.
* Spread the word in groups – join our [Steam Group](https://steamcommunity.com/groups/Mumble_VoIP)

# Skinning

Mumble skins bring variety into Mumbles design and choices for different tastes.

Create new skins for Mumble! Information can be found on our {{< wiki Skinning "Skinning wiki page" />}} and existing skins we know about on the {{< wiki Skins "Skins wiki page" />}}.

# Coding

## 3rd party Interfaces

Web interfaces, web-viewers, administration programs, CMS website modules – all neat things we would love to see! These could be written in any language supported by our Ice based RPC Interface (using the deprecated {{< wiki DBus />}} Interface is no longer advised). You could start your own project or link up with an {{< wiki "3rd Party Applications" "already existing one" />}}. To get an idea about how interaction with Murmur might look like you can take a look at the example scripts in our git repositories script folder.

## Ice authenticators

See {{< wiki Authenticators />}}. We’re happy for every authenticator and authenticator improvement you and the community can bring us!

## Positional Audio

Most of our positional audio plugins regularly break on game updates. To read on how you can help us update the plugins see the {{< wiki Pluginguide />}} page. You can even fix the update-issue forever for open source games, or if you want to add positional audio to your game, see {{< wiki Link />}}.

## Mumble itself

We are always looking for helping hands willing to contribute to mumble. You'll need a good understanding of C++ and previous knowledge in Qt is definitely helpful. The best way to get into mumble development is to link up with us on {{< wiki IRC />}} ([#mumble on freenode](irc://chat.freenode.net/mumble)), there is always small stuff to do to get you started. Also check the {{< wiki Development "Development wiki page" />}} on more info about the source code and building it.

{{<aside>}}This pages content is based on our {{< wiki Contributing "Contributing wiki page" />}}, which is licensed as [CC-By-SA](http://creativecommons.org/licenses/by-sa/2.5/).{{</aside>}}
