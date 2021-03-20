---
title: Contribute
---

You can contribute to the Mumble Project in multiple ways:

{{< toc >}}

## Report bugs and Feature Requests

1. Take a look at existing [issues in our GitHub repo](https://github.com/mumble-voip/mumble/issues) to check whether your issue or request was already reported.
   If that's the case, then vote for it with a thumbs up or add useful information & ideas with a comment.
2. If there is no issue report yet, then create a [new one](https://github.com/mumble-voip/mumble/issues/new/choose).

**Note:** For security issues, see [Report a vulnerability]({{< relref "security#reporting-a-vulnerability" >}}).

## Development

We are always looking for helping hands willing to contribute to Mumble.

### Coding

Mumble is mostly written in **C++** using the **Qt** library.
Thus previous knowledge in C++ and Qt is definitely useful, but we also encourage interested people who want to learn it to give it a try. We're happy to help you out, if needed.

#### Getting Started

The best way to get into Mumble development is to link up with us on IRC ([#mumble-dev on freenode](irc://chat.freenode.net/mumble-dev)), also available via Matrix (Join: `#freenode_#mumble-dev:matrix.org` or use the [Direct Link via Element.io (Web App)](https://app.element.io/#/room/!VNUpYnUPdhTAqagvUu:matrix.org)).

Alternatively you can create or comment (on) a pull request or issue report on GitHub.

For more information, see the categories below:

##### Source Code

Our source code is hosted on GitHub:

- [Overview of Mumble Projects/Repos](https://github.com/mumble-voip/)
- [Main Mumble Repo](https://github.com/mumble-voip/mumble): includes the Server and Client
- [iOS app](https://github.com/mumble-voip/mumble-iphoneos): unmaintained; help wanted

##### Documentation regarding Development

Take a look at the following documentation:

- [Contributing guidelines](https://github.com/mumble-voip/mumble#contributing)
- [Development Documentation](https://github.com/mumble-voip/mumble/tree/master/docs/dev)
- [Build Instructions](https://github.com/mumble-voip/mumble#building)

<!-- Should be implemented once it's ready:
- [C++ guidelines & tools](https://github.com/mumble-voip/mumble/discussions/4261) 
-->

##### Open issues and ongoing development

Also take a look at open [issues](https://github.com/mumble-voip/mumble/issues) and [pull requests](https://github.com/mumble-voip/mumble/pulls) on GitHub, especially with the following Labels:

- [Good first issue](https://github.com/mumble-voip/mumble/labels/good%20first%20issue)
- [Help wanted](https://github.com/mumble-voip/mumble/labels/help%20wanted)
- [Help needed](https://github.com/mumble-voip/mumble/labels/help-needed)
- [Feature Request](https://github.com/mumble-voip/mumble/labels/feature-request)
- [Bug](https://github.com/mumble-voip/mumble/labels/bug)

##### Tips

- To avoid unnecessary work and duplication, comment on issues and inform others what you plan to work on.
- You can also create [Draft Pull Requests](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests#draft-pull-requests); this way you can work on the Pull Request, while others can see the progress and give you feedback.

#### Additional projects

##### 3rd party Interfaces

Web interfaces, web-viewers, administration programs, CMS website modules – all neat things we would love to see! These could be written in any language supported by our Ice- or gRPC-based RPC Interfaces. You could start your own project or link up with an {{< wiki "3rd Party Applications" "already existing one" />}}. To get an idea about how interaction with the Mumble server might look like, you can take a look at the example scripts in our [GitHub repository for mumble-scripts](https://github.com/mumble-voip/mumble-scripts).

##### Ice authenticators

See {{< wiki Authenticators />}}. We’re happy for every authenticator and authenticator improvement.

##### Positional Audio plugins

Most of our [positional audio plugins](https://wiki.mumble.info/wiki/Games#Positional_audio) regularly break on game updates. To read on how you can help us update the plugins see the {{< wiki Pluginguide />}} page. 

You can also add positional audio support to Games via our standard `Link` plugin. This way it will not get outdated, but you need permission to change (and access to) the Source Code of a game (e.g. via an open license or developer permission). See {{< wiki "Link" "Linking a game to Mumble" />}}.

##### Other 3rd party projects

If you are interested in more projects (Clients, Servers, cli tools etc.), you can also take a look at {{< wiki "3rd Party Applications" />}}.

Many of these projects will either gladly accept help or can (often) be forked easily, because of the open licenses.

### Non-Coding

#### Translation

Mumble is available in various languages.

You can help improve or update existing translations, translate new texts on existing languages, or introduce additional languages.

We have a [Mumble project on Weblate](https://hosted.weblate.org/projects/mumble/mumble-client/) where you can register and start translating.

Weblate is a web translation platform. As a Website it is easy to access and use with any browser. For translators a global and project memory of translations as well as machine translations can help in finding good translations. Nevertheless, a native or good speaker is required to implement good translations.

The current status of our translations can be seen in the following graph:

[![Weblate translation status by language](https://hosted.weblate.org/widgets/mumble/-/mumble-client/horizontal-auto.svg)](https://hosted.weblate.org/engage/mumble/)

#### Testing

To find bugs and polish new releases we need testers.

If you want to test the newest development versions, you have two options:

- We provide development snapshots. You can download them from our [Downloads page]({{< relref "downloads#development-snapshots" >}}).

- Build the latest development version of Mumble, see [Building Mumble](https://github.com/mumble-voip/mumble#building).

If you find any issues, take a look at [Report Bugs](#report-bugs-and-feature-requests).

#### Development Documentation

See [Section "Documentation" below](#documentation). 

#### Mumble theme

You find the official Mumble theme in its own [repository on GitHub](https://github.com/mumble-voip/mumble-theme).

##### Additional themes

Additional Mumble themes bring variety into Mumble's visual design and choices for different tastes.

You can find more information about themes in general and how to create themes on our {{< wiki Themes "Themes wiki page" />}}.

## Documentation & Website

Improve our Documentation and Website.

### Documentation

Right now you find most of the **(User) Documentation** in our [wiki](https://wiki.mumble.info).

Any contributions are welcome: improving existing pages or creating new ones with useful content. Also see the {{< wiki "Category:Please_Improve" "Please Improve Category" />}}.

You can create a wiki account {{< wiki "Special:RequestAccount" "here" />}}.

The **Development Documentation** is now in [Mumble's GitHub repo](https://github.com/mumble-voip/mumble/tree/master/docs/dev).

It is written in Markdown.

You can create Pull Requests on GitHub to edit it or to add new files.

<!--
Future content:

The new documentation can be found in the following repos:

- **User Documentation** in the [Website's GitHub repo](https://github.com/mumble-voip/mumble-www/tree/master/hugo/content/documentation)
- **Development Documentation** in [Mumble's GitHub repo](https://github.com/mumble-voip/mumble/tree/master/docs)

It is written in Markdown.
-->

#### Documentation Issues & Pull requests

Also take a look at open Issues and Pull Requests:

- [Documentation issues in Mumble's GitHub repo](https://github.com/mumble-voip/mumble/labels/documentation)
- [Pull Request with TYPE label "Docs"](https://github.com/mumble-voip/mumble/pulls?q=is%3Apr+is%3Aopen+docs)
- [Documentation issues in Mumble's Website repo](https://github.com/mumble-voip/mumble-www/labels/docs)

#### Special Documentation

##### Ice methods

We have a [generated documentation](https://www.mumble.info/documentation/slice/1.3.0/html/Murmur.html) for exposed Ice methods. Although probably only developers will be able to understand it, please give us a note on any inconsistencies, errors or needed clarifications if you find them.sk

#### External Documentation

- **Wikipedia**: Someone that is not a development team member needs to maintain the [article(s) at Wikipedia](https://en.wikipedia.org/wiki/Mumble_%28software%29).

### Website

Take a look at the [Website's GitHub repo](https://github.com/mumble-voip/mumble-www).

### Community

Be a part of our community and help others!

You find the [Community Channels on the Contact page](https://www.mumble.info/contact#community-channels).

## Hosting

Host your own Mumble Server and register it on our Public Server List (optional).
See our {{< wiki "Running_Murmur" "Server guide" />}} and [Server Registration details](https://wiki.mumble.info/wiki/Murmur.ini#Server_Registration).

## Promotion

Recommend Mumble to your friends and colleagues ;).

Furthermore you can:

* Write about Mumble (e.g. in your blog or in social media).
* Ask (voice-)server hosters to consider providing Mumble servers.
* Create Videos, Tutorials or Benchmarks and publish them on prominent platforms.
