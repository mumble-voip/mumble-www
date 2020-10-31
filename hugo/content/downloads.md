---
title: Downloads
---
We provide downloads for the official client and server programs. A Linux distribution may provide their own packages and have their own maintainer, which we will describe below. We also link to some third party projects.

Version 1.3.3 is the latest stable version of Mumble and was released on October 5th, 2020.

## Suggested Mumble Version

<div id="suggested-download" style="display: grid; grid-template-columns: auto 1fr;">We suggest a download by determining the operating system with JavaScript. However JavaScript seems to not be enabled. Please refer to the manual download choices instead.</div>
<script>
'use strict'
/* For win always Win32 on Firefox and Chrome */
function parsePlatform(value) {
    value = value.toLowerCase()
    if (value.includes('Windows NT')) {
        if (value.includes('win64') || value.includes('wow64')) {
            return 'win64'
        }
        return 'win32'
    }
    if (value.includes('windows')) {
        return 'win64'
    }
    if (value.includes('Mac OS X') || value.includes('macOS')) {
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
    return '<a href="' + href + '" style="display: grid; grid-template-columns: auto 1fr; grid-gap: 16px; background: #3fb33f linear-gradient(#30dd30, #3fb33f); border: 1px solid #66b566; border-radius: 8px; padding: 24px 24px; color: #000; font-size: 32px; line-height: 32px;"><img src="/css/icons/' + icon + '" style="height: 32px; background: transparent;"><div style="margin: auto 0;">' + caption + '</div></a>'
}
function getPlatformContent(platform) {
    switch (platform) {
        case 'win64':
            return getButton('windows-64', 'windows.svg', 'Mumble for Windows 64-bit')
        case 'win32':
            return getButton('windows-32', 'windows.svg', 'Mumble for Windows 32-bit')
        case 'linux':
            return getButton('ubuntu', 'ubuntu.svg', 'Mumble PPA for Ubuntu')
        case 'macos':
            return getButton('osx', 'apple.svg', 'Mumble for macOS')
        default:
            return 'We could not determine the OS you are browsing this website on. Please choose the appropriate download yourself.'
            break;
    }
}
document.getElementById('suggested-download').innerHTML = getPlatformContent(getPlatform())
</script>

## Manual Download

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

For the individual files and for a zipped sources file see the [1.3.3 GitHub release page](https://github.com/mumble-voip/mumble/releases/tag/1.3.3).

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
