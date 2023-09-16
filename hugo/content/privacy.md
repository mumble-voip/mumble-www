---
title: Privacy
---
This privacy policy describes our various services and products and how we handle data within their context.

## This Website at `www.mumble.info`

The webserver may log non-person specific information.

As the website itself has no form of user input, we do not handle any user and personal information beyond simple and non-associated access logging.

Contributors to the website source code (content and technical) contributors will committed to the source code history in the form the contributor submits.  be shown as such under the names they identify themselves as and contribute as. Contributors may be shown on individual webpages for attribution.

## The Documentation Wiki at `wiki.mumble.info`

The Mumble documentation wiki makes use of Google Analytics for website usage statistics. A cookie with name `_ga` specific to our domain is used for tracking. Please refer to [this Google page for more information](https://support.google.com/analytics/answer/6004245).

The webserver may log non-person specific information.

Users registering and account and making changes to the wiki content will do so under the specified name, account and license. They will be listed in a pages history as a contributor along with their contribution.

## Mumble Client

### Optional Anonymous Usage Statistics

When a user decides to support us by submitting anonymous usage statistics some general usage and environment information will be sent to us. This includes operating system, ip address (for country information), Mumble version, and some other information. For the current specifics the client source code can be inspected.

### Optional Update Checks

If enabled the client will automatically check for available application updates. It will only send the required information to be able to do so: the application version and the operating system/type of application.

If enabled the positional audio plugin and overlay update check will send the corresponding version and operating system type in a similar fashion for this purpose.

The settings are located in the client applications settings dialog: Settings ➡ Network ➡ Privacy section.

### Public Server List

When opening the public server list in the connect dialog, the client will request this list from our server. Only the data required for doing so is sent.

This functionality can be disabled through a hidden configuration setting (not visible in the user interface).

### Data Sent to Listed and Connected To Servers (not us)

While this is not data being sent to us we want to mention this for completeness.

When you see servers listed in the connect dialog a “ping” request packet will be sent to it in order to display the ping (connection latency) and user count as you will see them in the user interface.

When connecting to a server the technically required data is sent to and received from the server. Please refer to the server providers privacy policy for details.

The user can prevent sending OS information to servers they connect to in the Settings ➡ Network ➡ Privacy section; Do not send OS information to Mumble and web servers.

## Mumble Server (previously called “Murmur”)

When the server is configured to register with our public server list ([`register*` settings][1]) the configured or automatically determined information (as described in the configuration documentation) is sent to us and consequently being listed publicly on our public server list.

 [1]: https://wiki.mumble.info/wiki/Murmur.ini#Server_Registration
