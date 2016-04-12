+++
title = "Mumo"
date = "2016-04-12T20:28:38"
license = "CC by-sa 2.5"
categories = ["3rd Party"]
+++
# mumo - The Mumble moderator

Mumo is meant to provide a platform on which Python based Mumble server
plugins can be built upon. The goal is to reduce the boilerplate needed
to interact with the Mumble server to a minimum.

To achieve this goal tasks like [Ice](https://wiki.mumble.info/wiki/Ice) interface setup, basic error
handling, configuration management, logging and more are provided
by mumo. Developers can focus on their specific functionality instead.

## Currently mumo comes with the following modules
### bf2
Battlefield 2 game management plugin that can dynamically
move players into appropriate channels and groups to fit
the in-game command structure. This is achieved by using
data gathered from Mumble's positional audio system and does
not require cooperation by the game server. 

### idlemove
Plugin for moving players that have been idle for
a configurable amount of time into a idle channel. Optionally
the players can be muted/deafened on move.

### seen
Makes the server listen for a configurable keyword to ask for
the last time a specific nick was seen on the server.

### [source](https://wiki.mumble.info/wiki/mumo_source)
Source game management plugin that can dynamically move
players into on-the-fly created channel structures representing
in-game team setup. This is achieved by using data gathered from
Mumble's positional audio system and does not require cooperation
by the game server. Currently the following source engine based
games are supported: Team Fortress 2, Day of Defeat: Source,
Counter-Strike: Source, Half-Life 2: Deathmatch.

For more information on how to configure and use the source plugin see [here](https://wiki.mumble.info/wiki/mumo_source).

### test
A debugging plugin that registers for all possible events and
outputs every call with parameters into the debug log.

## 3rd Party Modules
### Image Injection Modules
#### URL to image
[urltoimg on GitHub](https://github.com/aciid/urltoimg-for-mumo)

If you send an image URL to the channel, the server fetches it and sends the image instead of the image link to the current channel.

#### Chat Img
[chatimg on GitHub](https://github.com/aselus-hub/chatimg-mumo)

A more full featured implementation of the same functionality as Url to Image. Allows injection of photos into chat, re-sizing them if they are larger then the size accepted by the mumble protocol and allowing the server admin to set a max width/height for the image so that it is scaled through html or thumbnailing to never be larger then prescribed. Allows the conversion of images within regular chat messages w/o bang commands as an option. Finally the injection of any number of images present after the bang or within a message.

### Message forwarder
[forwarder-for-mumo on GitHub](https://github.com/Natenom/forwarder-for-mumo)

Forwards messages to all linked channels.

### Wrong Version
[wrongversion on GitHub](https://github.com/Natenom/wrongversion)

Users with older Mumble versions will get a message (triggering version is configurable).

### Lowbw
[mumo-lowbw on GitHub](https://github.com/Natenom/mumo-lowbw)

Creates a Low bw channel whenever a member of group "bots" enters a room and removes it later if empty...

### Videoinfo
[mumo-videoinfo on GitHub](https://github.com/while-loop/mumo-videoinfo)

Mumo plugin to provide YouTube video information to Mumble.

### Set Status
[setstatus on GitHub](https://github.com/Natenom/setstatus-for-mumo)

Users can add a temporary status to your username via chat command.

### Deaf to AFK
[deaftoafk on GitHub](https://github.com/Natenom/deaftoafk-for-mumo)

Moves deafened users into an AFK channel.

### Anti recording
* [antirec on GitHub](https://github.com/Natenom/antirec-for-mumo/)
* [antirec with contextmenus on GitHub](https://github.com/Natenom/antirec-for-mumo-contextmenu)

If a user starts to record in Mumble he will be deafened or kicked from the server.

### Max users
[mumo-maxusers on GitHub](https://github.com/ExplodingFist/mumo-maxusers/)

This is a MuMo module to provide an administrator the capability of enforcing granular user limits by channel in mumble.

### Opcommand
[mumo-opcommand on GitHub](https://github.com/ExplodingFist/mumo-opcommand)

Temporarily add user or remove user to/from a group via GUI command line.

### Sticky
* [mumo-sticky on GitHub](https://github.com/Natenom/mumo-sticky)
* [mumo-sticky with contextmenus on GitHub](https://github.com/Natenom/mumo-sticky-contextmenu)

A user who gets the sticky status can't do more than sitting in one special channel. Even admins will loose their permissions while sticked.

### AntiFlood
[antiflood-for-mumo on GitHub](https://github.com/Natenom/antiflood-for-mumo)

If a user exceeds a limit of actions within a timeframe (both can be changed) he will be kicked from the server.

### GetSupport
[mumo-getsupport on GitHub](https://github.com/Natenom/mumo-getsupport)

A user can write for example "!helpme support message up to 160 characters..." to a channel/user/himself and everyone on the server who is member of a special group gets a notification.

### OS-Suffixes
[mumo-os-suffixes on GitHub](https://github.com/Natenom/mumo-os-suffixes)

Adds suffixes to user names based on operating system.

## Contributions
If you have a module that you would like to see shipped with mumo or
have any improvements or suggestions please contact us. Whether you
prefer a pull request, visiting us in IRC in #mumble on Freenode or
starting a thread in our forums at http://forums.mumble.info is up to you.

## Configuration
To configure and run mumo take a look at the mumo.ini and the module
specific configurations in modules-available folder. Enabling modules
is done by linking the configuration in modules-available to the
modules-enabled folder.

## Requirements
mumo requires:

    - python 2.7*
    - python-zeroc-ice
    - murmur >=1.2.3*
    - murmur >=1.2.4 if you want to use a module which needs user interaction via text commands

* Not tested with lower versions yet

# Installing mumo
The newest version of mumo is always available from our mumo repository at https://github.com/mumble-voip/mumo .

## Ubuntu 12.04
*Note:* This guide only shows the basic steps for trying out mumo. For a more permanent setup you'll want to run mumo with its own user and a startup script.

* Make sure you are running a recent Murmur release (1.2.4 or later). Ice should be enabled and a writesecret must be set (see configuration file).
* Install dependencies

    sudo apt-get install python-zeroc-ice python-daemon git

* Clone repository

    cd ~/
    git clone https://github.com/mumble-voip/mumo.git

* Adjust configuration

    cd mumo
    nano mumo.ini

In the editor set your server's Ice writesecret as the secret variable so mumo can control your server.

    secret = mysecretwritesecret

Close and save by pressing Ctrl + X followed by Y and Enter.
* Configure the modules you want to use by editing their ini file in the modules-available folder
* Enable modules by linking their config file into the modules-enabled folder

    cd modules-enabled
    ln -s ../modules-available/moduleyouwanttouse.ini

* Run mumo

    ./mumo.py

Mumo should now be working with your mumble server. If it doesn't work check the *mumo.log* file for more information.

