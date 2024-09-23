---
title: Mumble Server Scripting
date: 2019-10-27
---
The Mumble Server can listen to commands from other applications to allow scripting specific actions and functionality on triggers.

For example:

* Authenticators can be used to authenticate connecting users through an existing database (for example Active Directory, a Forum Database, Website Database, etc)
* Context Menu extensions can provide additional context menu (right click) actions to users
* Adjust user state
* Channel Viewers (online users)
  * Connected directly
  * Making use of [Channel Viewer Protocol]({{< relref "/documentation/developer/channel-viewer-protocol" >}}) provider
* Management of server, user or database state, for example with webinterface/website

Different technologies can be used for scripting. Ice is the most stable and complete one.

| Technology | State | Documentation |
|---|---|---|
| Ice | Stable | [Ice Scripting]({{< relref "ice" >}}), [Slice Types Docs](/documentation/slice/) |
| <abbr title="Command Line Interface">CLI</abbr> | Stable | [RPC subcommand]({{< relref "/documentation/user/rpc" >}}) |
| DBus | Obsolete, Incomplete | N/A |
| GRPC | Unreleased, Buildable, Incomplete | N/A |

