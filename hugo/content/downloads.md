---
title: Downloads
---
We provide downloads for the official client and server programs. A Linux distribution may provide their own packages and have their own maintainer,
which we will describe below. We also link to some third party projects.

Version **1.3.4** is the latest stable version of Mumble and was released on February 10th, 2021.

**Note:** Due to problems with our macOS builder, we are currently unable to provide 1.3.4 binaries for macOS. We are working on it, but for the time
being the **macOS download** will still refer to version **1.3.3**.

## Suggested Mumble Version

<div id="suggested-download" style="display: grid; grid-template-columns: auto 1fr;">We suggest a download by determining the operating system with
JavaScript. However JavaScript seems to not be enabled. Please refer to the manual download choices instead.</div>
<script>
'use strict'
/* For win always Win32 on Firefox and Chrome */
function parsePlatform(value) {
    if (!value)
        return false;
    value = value.toLowerCase()
    if (value.includes('windows nt')) {
        if (value.includes('win64') || value.includes('wow64')) {
            return 'win64'
        }
        return 'win32'
    }
    if (value.includes('windows')) {
        return 'win64'
    }
    if (value.includes('mac os x') || value.includes('macos')) {
        return 'macos'
    }
    if (value.includes('android')) {
        return 'android'
    }
    if (value.includes('linux')) {
        return 'linux'
    }
}
function getPlatform() {
    return parsePlatform(navigator.oscpu) || parsePlatform(navigator.appVersion) || parsePlatform(navigator.userAgent) || parsePlatform(navigator.platform)
}
function getButton(href, icon, caption) {
    return '<a class="suggested-download-button" href="' + href + '"><img class="suggested-download-button-icon" src="/css/icons/' + icon + '"><div class="suggested-download-button-caption">' + caption + '</div></a>'
}
function getPlatformContent(platform) {
    switch (platform) {
        case 'win64':
            return getButton('https://dl.mumble.info/latest/stable/client-windows-x64', 'windows.svg', 'Mumble for Windows (x64)')
        case 'win32':
            return getButton('https://dl.mumble.info/latest/stable/client-windows-x86', 'windows.svg', 'Mumble for Windows (x86)')
        case 'linux':
            return getButton('https://launchpad.net/~mumble/+archive/release', 'ubuntu.svg', 'Mumble PPA for Ubuntu')
        case 'macos':
            return getButton('https://dl.mumble.info/latest/stable/client-macos-x64', 'apple.svg', 'Mumble for macOS >= 10.13 (x64)')
        default:
            return 'We could not determine the OS you are browsing this website on. Please choose the appropriate download yourself.'
            break;
    }
}
document.getElementById('suggested-download').innerHTML = getPlatformContent(getPlatform())
</script>

## Manual Download

{{< content-layout/downloads >}}
{{< content-layout/download name="Windows (x64)" href="https://dl.mumble.info/latest/stable/client-windows-x64" osclass="windows">}}
{{< content-layout/download name="Windows (x86)" href="https://dl.mumble.info/latest/stable/client-windows-x86" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="macOS >= 10.13 (x64)" href="https://dl.mumble.info/latest/stable/client-macos-x64" osclass="mac">}}
{{< content-layout/download name="macOS <= 10.12 (x64)" href="https://dl.mumble.info/latest/stable/client-macos-hfs+-x64" osclass="mac">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="Ubuntu" href="https://launchpad.net/~mumble/+archive/release" osclass="ubuntu">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="Static macOS server (x64)" href="https://dl.mumble.info/latest/stable/server-macos-x64" osclass="mac">}}
{{< content-layout/download name="Static Linux server (x86)" href="https://dl.mumble.info/latest/stable/server-linux-x86" osclass="linux">}}
{{< /content-layout/downloads >}}

**Note:** `Static Linux Server` is using an outdated version of OpenSSL - see [#4001](https://github.com/mumble-voip/mumble/issues/4001) for details. 

Server note: *Depending on the context and packager our server program is called “mumble-server” or “Murmur”*

For the individual files and for a zipped sources file see the [1.3.4 GitHub release page](https://github.com/mumble-voip/mumble/releases/tag/1.3.4).

Instructions on [verifying GPG signatures of Mumble downloads](https://github.com/mumble-voip/mumble-gpg-signatures/blob/master/gpg.txt) can be found
in the linked document.

## Linux

Your distribution probably provides official packages for Mumble. Please refer to your OS packages. Distributions with official packages include
Debian, Ubuntu, Fedora, openSUSE, Arch Linux, Mandriva/ROSA/Unity.

### Ubuntu

Ubuntu provides their own packages. But we also provide a PPA (Personal Package Archive) which you can add to use more recent stable or development
snapshot versions.

    sudo add-apt-repository ppa:mumble/release

### Snap/Snapcraft

A Snap package is published and maintained by the third party Snapcraft community at <https://snapcraft.io/mumble>

### Flatpak

A Flatpak Package is published and maintained by the third party Flatpak community at <https://www.flathub.org/apps/details/info.mumble.Mumble>

## Mobile Clients

{{< content-layout/downloads >}}
{{< content-layout/download name="iOS" href="ios" osclass="ios">}}
{{< /content-layout/downloads >}}

**Note:** The iOS app is unmaintained; last update in 2017.

### Android (third party)

We do not currently provide an official Android client. However, you can use one of the third party clients (that we do not develop or support
ourselves).

{{< content-layout/downloads >}}
{{% content-layout/download-ext osclass="android" %}}Mumla ([Play](https://play.google.com/store/apps/details?id=se.lublin.mumla),<br>[Play-Donate](https://play.google.com/store/apps/details?id=se.lublin.mumla.donation), [F-Droid](https://f-droid.org/packages/se.lublin.mumla)){{%/ content-layout/download-ext %}}
{{< /content-layout/downloads >}}

## Development snapshots

Development snapshots contain unreleased features and changes that will eventually be available in the next stable release. Please report any problems
you encounter on [our issue tracker](https://github.com/mumble-voip/mumble/issues).

The most recent snapshot version is actually the first release candidate for Mumble 1.4.0
([Release Announcement]({{< relref "/blog/2021-09-12-mumble-release-candidate1" >}})).

Note that we are currently not able to provide static server binaries for Linux nor are we able to provide a snapshot PPA for Linux just yet.

{{< content-layout/downloads >}}
{{< content-layout/download name="Client for Windows (x64)" href="https://dl.mumble.info/latest/snapshot/client-windows-x64" osclass="windows">}}
{{< content-layout/download name="Server for Windows (x64)" href="https://dl.mumble.info/latest/snapshot/server-windows-x64" osclass="windows">}}
{{< content-layout/download name="Client for Windows (x86)" href="https://dl.mumble.info/latest/snapshot/client-windows-x86" osclass="windows">}}
{{< content-layout/download name="Server for Windows (x86)" href="https://dl.mumble.info/latest/snapshot/server-windows-x86" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="Client for macOS >= 10.13 (x64)" href="https://dl.mumble.info/latest/snapshot/client-macos-x64" osclass="mac">}}
{{< content-layout/download name="Server for macOS >= 10.13 (x64)" href="https://dl.mumble.info/latest/snapshot/server-macos-x64" osclass="mac">}}
{{< /content-layout/downloads >}}

Note that we are no longer able to support macOS < 10.13 since [Qt does not support](https://doc.qt.io/qt-5/macos.html#supported-versions) these
versions anymore.

## Source Code

As a free software project the source code is publicly readable and under a permissive license.

Our Mumble project source code is hosted on GitHub as [`mumble`](https://github.com/mumble-voip/mumble). The iOS app as
[`mumble-iphoneos`](https://github.com/mumble-voip/mumble-iphoneos).

