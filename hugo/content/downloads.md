---
title: Downloads
---
We provide downloads for the official client and server programs. A Linux distribution may provide their own packages and have their own maintainer, which we will describe below. We also link to some third party projects.

Version 1.3.2 is the latest stable version of Mumble and was released on July 08, 2020.

{{< content-layout/downloads >}}
{{< content-layout/download name="Windows 64-bit" href="windows-64" osclass="windows">}}
{{< content-layout/download name="Windows 32-bit" href="windows-32" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="macOS" href="osx" osclass="mac">}}
{{< content-layout/download name="Static macOS Server" href="osx-static-server" osclass="mac">}}
{{< /content-layout/downloads >}}

Note (macOS): Apple introduced a stricter verification and warning process. macOS will now warn you when installing the Mumble application. This is expected and not a security issue. Apple will not currently verify and acknowledge our package unless we pay a significant amount of money. *(Issue tracked in {{< issue 4263 >}})*

{{< content-layout/downloads >}}
{{< content-layout/download name="Ubuntu" href="ubuntu" osclass="ubuntu">}}
{{< content-layout/download name="Static Linux Server" href="linux-static-server" osclass="linux">}}
{{< /content-layout/downloads >}}

Server note: *Depending on the context and packager our server program is called “mumble-server” or “Murmur”*

For the individual files and for a zipped sources file see the [1.3.2 GitHub release page](https://github.com/mumble-voip/mumble/releases/tag/1.3.2).

Instructions on [verifying GPG signatures of Mumble downloads](https://github.com/mumble-voip/mumble-gpg-signatures/blob/master/gpg.txt) can be found in the linked document.

## Linux

Your distribution probably provides official packages for Mumble. Please refer to your OS packages. Distributions with official packages include Debian, Ubuntu, Fedora, openSUSE, Arch Linux, Mandriva/ROSA/Unity.

### Ubuntu

Ubuntu provides their own packages. But we also provide a PPA (Personal Package Archive) which you can add to use more recent stable or development snapshot versions.

    sudo add-apt-repository ppa:mumble/release

### Snap/Snapcraft

A Snap package is published and maintained by the third party Snapcraft community at <https://snapcraft.io/mumble>

### Flatpak

A Flatpak Package is published and maintained by the third party Flatpak community at <https://www.flathub.org/apps/details/info.mumble.Mumble>

## Mobile Clients

{{< content-layout/downloads >}}
{{< content-layout/download name="iOS" href="ios" osclass="ios">}}
{{< /content-layout/downloads >}}

### Android (third party)

We do not currently provide an official Android client. However, you can use one of the third party clients (that we do not develop or support ourselves).

{{< content-layout/downloads >}}
{{% content-layout/download-ext osclass="android" %}}Mumla ([Play](https://play.google.com/store/apps/details?id=se.lublin.mumla),<br>[Play-Donate](https://play.google.com/store/apps/details?id=se.lublin.mumla.donation), [F-Droid](https://f-droid.org/packages/se.lublin.mumla)){{%/ content-layout/download-ext %}}
{{< /content-layout/downloads >}}

## Development snapshots

[Development snapshots](https://dl.mumble.info/) contain unreleased features and changes that will eventually be available in the next stable release. Please report any problems you encounter on [our issue tracker](https://github.com/mumble-voip/mumble/issues).

At the moment we are in the process of changing our build infrastructure for releases. Our next snapshots will be based on the new system. **Until the new system is established we are not publishing development snapshots** (using the old system).

Note (macOS): Apple introduced a stricter verification and warning process. macOS will now warn you when installing the Mumble application. This is expected and not a security issue. Apple will not currently verify and acknowledge our package unless we pay. *(Issue tracked in {{< issue 4263 >}})*

## Source Code

As a free software project the source code is publicly readable and under a permissive license.

Our Mumble project source code is hosted on GitHub as [`mumble`](https://github.com/mumble-voip/mumble). The iOS app as [`mumble-iphoneos`](https://github.com/mumble-voip/mumble-iphoneos).
