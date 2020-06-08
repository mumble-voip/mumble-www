---
title: Downloads
---
We provide downloads for the official client and server programs. A Linux distribution may provide their own packages and have their own maintainer, which we will describe below. We also link to some third party projects.

Version 1.3.1 is the latest stable version of Mumble and was released on June 08, 2020.

{{< content-layout/downloads >}}
{{< content-layout/download name="Windows 64-bit" href="windows-64" osclass="windows">}}
{{< content-layout/download name="Windows 32-bit" href="windows-32" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="macOS" href="osx" osclass="mac">}}
{{< content-layout/download name="Static macOS Server" href="osx-static-server" osclass="mac">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="Ubuntu" href="ubuntu" osclass="ubuntu">}}
{{< content-layout/download name="Static Linux Server" href="linux-static-server" osclass="linux">}}
{{< /content-layout/downloads >}}

Server note: *Depending on the context and packager our server program is called “mumble-server” or “Murmur”*

For the individual files and for a zipped sources file see the [1.3.1 GitHub release page](https://github.com/mumble-voip/mumble/releases/tag/1.3.1).

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
{{< content-layout/download name="Mumla (F-Droid)" href="https://f-droid.org/packages/se.lublin.mumla" osclass="android">}}
{{< /content-layout/downloads >}}

## Development snapshots

[Development snapshots](https://dl.mumble.info/) contain unreleased features and changes that will eventually be available in the next stable release. Please report any problems you encounter on [our issue tracker](https://github.com/mumble-voip/mumble/issues).

When a new feature release is published the snapshots will naturally transform into them. If you want to keep using snapshots after the next feature release, you will have to manually upgrade to one again.

<table class="development-snapshots">
    <tr>
        <th>Windows 64-bit</th>
        <th>Windows 32-bit</th>
        <th>macOS</th>
        <th>macOS (Universal, Legacy)</th>
        <th>Ubuntu</th>
        <th>Static Linux Server</th>
        <th>Static macOS Server</th>
    </tr>
    <tr>
        <td>
            <a href="windows-64/snapshot">.msi</a>
        </td>
        <td>
            <a href="windows-32/snapshot">.msi</a>
        </td>
        <td>
            <a href="osx/snapshot">.dmg</a>
        </td>
        <td>
            <a href="osx-universal/snapshot">.dmg</a>
        </td>
        <td>
            <a href="ubuntu/snapshot">PPA</a>
        </td>
        <td>
            <a href="linux-static-server/snapshot">.tar.bz2</a>
        </td>
        <td>
            <a href="osx-static-server/snapshot">.tar.bz2</a>
        </td>
    </tr>
</table>
