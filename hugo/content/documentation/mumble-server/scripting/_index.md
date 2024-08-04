---
title: Mumble Server Scripting
date: 2019-10-27
---

The Mumble Server can listen to commands from other applications to allow scripting specific actions and functionality
on triggers.

For example:

- {{< wiki "3rd_Party_Applications#Authenticators" "Authenticators" />}} can be used to authenticate connecting users
  through an existing database (for example Active Directory, a Forum Database, Website Database, etc)
- Context Menu extensions can provide additional context menu (right click) actions to users
- Adjust user state
- Channel Viewers (online users)
  - Connected directly
  - Making use of {{< wiki "Channel Viewer Protocol" "Channel Viewer Protocol" />}} provider
- Management of server, user or database state, for example with webinterface/website

Different technologies can be used for scripting. Ice is the most stable and complete one.

| Technology                                      | State                             | Documentation                                                                    |
| ----------------------------------------------- | --------------------------------- | -------------------------------------------------------------------------------- |
| Ice                                             | Stable                            | [Ice Scripting]({{< relref "ice" >}}), [Slice Types Docs](/documentation/slice/) |
| <abbr title="Command Line Interface">CLI</abbr> | Stable                            | {{< wiki "RPC subcommand" />}}                                                   |
| DBus                                            | Obsolete, Incomplete              | {{<wiki DBus />}}                                                                |
| GRPC                                            | Unreleased, Buildable, Incomplete | {{< wiki GRPC />}}                                                               |
