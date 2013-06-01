+++
title = "Mumo"
date = "2013-06-01T21:09:11"
license = "CC by-sa 2.5"
categories = []
+++
# mumo - The Mumble moderator

Mumo is meant to provide a platform on which python based Mumble server
plugins can be built upon. The goal is to reduce the boilerplate needed
to interact with the Mumble server to a minimum.

To achieve this goal tasks like Ice interface setup, basic error
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

## Contributions
If you have a module that you would like to see shipped with mumo or
have any improvements or suggestions please contact us. Whether you
prefer a pull request, visiting us in IRC in #mumble on Freenode or
starting a thread in our forums at http://sourceforge.net/apps/phpbb/mumble/ is up to you.

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

* Not tested with lower versions yet

# Installing mumo
The newest version of mumo is always available from our mumble-scripts repository at http://gitorious.org/mumble-scripts .
## Ubuntu 12.04
*Note:* This guide only shows the basic steps for trying out mumo. For a more permanent setup you'll want to run mumo with its own user and a startup script.

* Make sure you are running a recent Murmur release (1.2.4 or later). Ice should be enabled and a writesecret must be set (see configuration file).
* Install dependencies

    sudo apt-get install python-zeroc-ice python-daemon git

* Clone repository

    cd ~/
    git clone git://gitorious.org/mumble-scripts/mumo.git

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