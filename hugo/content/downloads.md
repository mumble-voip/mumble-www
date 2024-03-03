---
title: Downloads
---
We provide downloads for the official client and server programs. A Linux distribution may provide their own packages and have their own maintainer,
which we will describe below. We also link to some third party projects.

Version **1.4.287** is the latest stable version of Mumble and was released on August 21st, 2022.

Note that we are no longer able to support macOS < 10.13 since [Qt does not support](https://doc.qt.io/qt-5/macos.html#supported-versions) these
versions anymore.

Note also that you have to **uninstall any previous Mumble version** before installing 1.4, since the **upgrade path is unfixably broken**. See
[here](https://github.com/mumble-voip/mumble/issues/5076) for more info.


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
function getWinRedistNotice() {
    return '<div style="grid-row: 2; font-style: italic;">Requires installed Microsoft Visual C++ 2015-2022 Redistributable <a href="https://aka.ms/vs/17/release/vc_redist.x64.exe">x64</a> and <a href="https://aka.ms/vs/17/release/vc_redist.x86.exe">x86</a> to run</div>'
}
function getPlatformContent(platform) {
    switch (platform) {
        case 'win64':
            return getButton('https://dl.mumble.info/latest/stable/client-windows-x64', 'windows.svg', 'Mumble for Windows (x64)')
                + getWinRedistNotice()
        case 'win32':
            return getButton('https://dl.mumble.info/latest/stable/client-windows-x86', 'windows.svg', 'Mumble for Windows (x86)')
                + getWinRedistNotice()
        case 'linux':
			return 'For Linux, please refer to the dedicated section below.';
            // return getButton('https://launchpad.net/~mumble/+archive/release', 'ubuntu.svg', 'Mumble PPA for Ubuntu')
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

*Note: On Windows Microsoft Visual C++ 2015-2022 Redistributable (<a href="https://aka.ms/vs/17/release/vc_redist.x64.exe">x64 download</a> and <a href="https://aka.ms/vs/17/release/vc_redist.x86.exe">x86 download</a>) is required to be installed.*

{{< content-layout/downloads >}}
{{< content-layout/download name="Windows client (x64)" href="https://dl.mumble.info/latest/stable/client-windows-x64" osclass="windows">}}
{{< content-layout/download name="Windows client (x86)" href="https://dl.mumble.info/latest/stable/client-windows-x86" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="Windows server (x64)" href="https://dl.mumble.info/latest/stable/server-windows-x64" osclass="windows">}}
{{< content-layout/download name="Windows server (x86)" href="https://dl.mumble.info/latest/stable/server-windows-x86" osclass="windows">}}
{{< /content-layout/downloads >}}

{{< content-layout/downloads >}}
{{< content-layout/download name="macOS >= 10.13 (x64)" href="https://dl.mumble.info/latest/stable/client-macos-x64" osclass="mac">}}
{{< /content-layout/downloads >}}

<!-- The PPA is completely outdated
{{< content-layout/downloads >}}
{{< content-layout/download name="Ubuntu" href="https://launchpad.net/~mumble/+archive/release" osclass="ubuntu">}}
{{< /content-layout/downloads >}}
-->

{{< content-layout/downloads >}}
{{< content-layout/download name="Static macOS server (x64)" href="https://dl.mumble.info/latest/stable/server-macos-x64" osclass="mac">}}
<!--
{{< content-layout/download name="Static Linux server (x64)" href="https://dl.mumble.info/latest/stable/server-linux-x64" osclass="linux">}}
-->
{{< /content-layout/downloads >}}

Server note: *Depending on the context and packager our server program is called “mumble-server” or “Murmur”*

For the individual files and for a zipped sources file see the [1.4.282 GitHub release page](https://github.com/mumble-voip/mumble/releases/tag/v1.4.287).

Instructions on [verifying GPG signatures of Mumble downloads](https://github.com/mumble-voip/mumble-gpg-signatures/blob/master/gpg.txt) can be found
in the linked document.

## Linux

Your distribution probably provides official packages for Mumble. Please refer to your OS packages. Distributions with official packages include
Debian, Ubuntu, Fedora, openSUSE, Arch Linux, Mandriva/ROSA/Unity.

### Snap/Snapcraft

A Snap package is published and maintained by the third party Snapcraft community at <https://snapcraft.io/mumble>

### Flatpak

A Flatpak Package is published and maintained by the third party Flatpak community at <https://www.flathub.org/apps/details/info.mumble.Mumble>

## Mobile Clients

{{< content-layout/downloads >}}
{{< content-layout/download name="iOS" href="https://apps.apple.com/us/app/mumble/id443472808" osclass="ios">}}
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

The most recent snapshot version is the second release candidate for Mumble 1.5.x
([Release Announcement]({{< relref "/blog/2024-03-03-mumble-1-5-613" >}})).

Note that we are currently not able to provide static server binaries for Linux nor are we able to provide a snapshot PPA for Linux just yet.

{{< content-layout/downloads >}}
{{< content-layout/download name="Client for Windows (x64)" href="https://dl.mumble.info/latest/snapshot/client-windows-x64" osclass="windows">}}
{{< content-layout/download name="Server for Windows (x64)" href="https://dl.mumble.info/latest/snapshot/server-windows-x64" osclass="windows">}}
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

