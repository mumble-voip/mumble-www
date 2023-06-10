---
title: ACL setup for a Guest Channel
date: 2023-06-10
---
This guide will

1. Set up a default channel for users to connect to
2. Deny not registered users to join other channels

This can be helpful as a safeguard against malicious users and guests; as a gate to new users until someone confirms their intentions or acceptance.

## Setting the default channel

The Mumble server configuration has the setting `defaultchannel` to set the default channel.

[The `defaultchannel` setting and it's description](https://github.com/mumble-voip/mumble/blob/f09c0746450c71f7f26dc1942707d68c987fee05/auxiliary_files/mumble-server.ini#L187-L198):

```ini
; If a user has no stored channel (they've never been connected to the server
; before, or rememberchannel is set to false) and the client hasn't been given
; a URL that includes a channel path, the default behavior is that they will
; end up in the root channel.
;
; You can set this setting to a channel ID, and the user will automatically be
; moved into that channel instead. Note that this is the numeric ID of the
; channel, which can be a little tricky to get (you'll either need to use an
; RPC mechanism, watch the console of a debug client, or root around through
; the server database to get it).
;
;defaultchannel=0
```

Create the new channel you want users to join by default (if no last channel of the rejoining registered user is known). To add a channel, with adequate permissions, or when logged in as SuperUser, right click and select Add Channel.

After creating the channel check the server log for a line similar to `<W>2010-02-02 22:30:34.873 1 => <26:SuperUser(0)> Added channel Entrance[4:0] under Root[0:-1]`.

Where you can find the log file depends on your server configuration or operating system. If you changed the `logfile` setting that's where or under what name you will find it. Otherwise:

* Windows default: `C:\Program Files\Mumble\server\mumble-server.log`
* Linux default: `/var/log/mumble-server/mumble-server.log`

In the example `<W>2010-02-02 22:30:34.873 1 => <26:SuperUser(0)> Added channel Entrance[4:0] under Root[0:-1]`, a channel named `Entrance` was created with channel ID `4`.

In your server settings file (`C:\Program Files\Mumble\server\mumble-server.ini` or `/etc/mumble-server.ini`) uncomment and set `defaultchannel` to the channel ID:

```ini
defaultchannel=4
```

## Permissions

1. Open the ACL settings on the root channel (right click -> Edit -> ACL)
2. Add a `all` group with *Deny* on *Enter*, *Register User*, *Register Self*
3. Ensure `Applies to sub-channels` is set
4. If you have other rules added, move the rule as far up as possible (right below the uneditable default `@all` rule)
5. Confirm the changes with OK
6. Open the ACL settings for the new *Entrance* channel
7. Ensure the other rule from the root channel is inherited
8. Add a new `@all` rule with *Allow* *Traverse*, *Enter*, *Speak*, and *Deny* all other permissions
9. Add a new `@auth` rule with *Allow* *Traverse*, *Enter*, *Speak*; don't set any other permissions
10. Optional: Allow *Move* to `@auth` so anyone can move guests out of the entrance
11. Confirm the changes with OK

Now, new connecting users will join the default channel. They will be able to speak but not be able to join other channels.

Someone of the authenticated users will be able to join the entrance channel and welcome and accept them in.
