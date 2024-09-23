---
title: Mumble Config File
date: 2024-09-23
---

### Configuration template

The most recent version of the configuration file is kept up-to-date in the Mumble GitHub repository:
<https://github.com/mumble-voip/mumble/blob/master/auxiliary_files/mumble-server.ini>

### Overview

The ``mumble-server`` configuration file consists of single line configuration settings, in the format
``key=value``. Empty lines and anything after ``#`` to the end of a line is ignored.

Some settings are process-wide (``database``, etc.) while others are used as defaults on a per-server basis.
``registername`` for exmaple will be used as the default "title" for all virtual servers, unless it is overridden
via RPC.

A unique example is ``port`` - the first server will attempt to bind to this port, with
each subsequent server incrementing the port by one for its own port.

### ICE

#### ``ice``

The ``mumble-server`` supports extensive RPC via [ZeroC ICE](https://doc.zeroc.com/ice/latest).
The ``ice`` setting specifies the "endpoint". Note that all virtual servers share
an ICE service. **Without secrets configured** anyone with access to the ICE port has
**full administrative access** to every virtual server on the ``mumble-server`` process.
It's not recommended to expose it to the world for that reason, but rather to protect
the service via a firewall and configure ICE secrets.

Example:

```text
ice="tcp -h 127.0.0.1 -p 6502"
```

#### ``icesecretread`` and ``icesecretwrite``

The default ICE configuration listens on a local TCP socket, which means that anyone with
access to the local machine can issue commands over it.
Setting the ICE secrets will require anyone to posess them to be able to access the
ICE service.

Access is split into two levels: "read" (look only) and "write" (modify). "write" access
implies "read" access.

Example:

```text
icesecretread=letmelook
icesecretwrite=letmechangestuff
```

Note that if either of these entries is **uncommented** and **empty**, access will be denied.
To **disable** secrets, **comment** these settings out.

### Server Registration

Mumble supports a public server registry, which your ``mumble-server`` can automatically configure itself
to register with if you choose. The server must be public (no ``serverpassword``), it must be accessible
worldwide, and it must have configured a valid certificate with TLS bit enabled (or a ``mumble-server`` automatically
generated self-signed certificate).

``serverpassword`` must be empty, and all the other ``register*`` settings must be filled in. ``registerHostname``
**must** be valid and resolve properly.

#### ``registerName``

The ``registerName`` setting specifies the "name" of your server in the public server list. The list is no longer
sorted alphabetically by default, so please, in the interests of being user-friendly, don't set it to something silly
trying to arbitrarily "rank" your server "further up the list".

Example:

```text
registerName=Super Awesome Mumble Server
```

#### ``registerPassword``

The ``registerPassword`` is a simple, plain-text secret between your server and the registration server.
Its sole purpose is to prevent other servers from impersonating your server in the public list.

Leaving this setting blank will disable registration with the public server list.

If you somehow lose the secret, you will obviously not be able to register your server until the
public server list "forgets" about your server. This usually takes a few days.

Exmaple:

```text
registerPassword=CorrectHorseBatteryStaple
```

#### ``registerUrl``

This is a URL pointing to the website for your server.

Example:

```text
registerUrl=https://mumble.info/
```

#### ``registerHostname``

This setting is the DNS hostname where your server can be reached. It only needs to be set, if you want
your server to be addressed by its hostname instead of by IP. But if it is set, it **must** resolve on the
public internet or registration will fail.

Example:

```text
registerHostname=mumble.some.host
```

#### ``registerLocation``

Deprecated.

### Misc

#### ``bandwidth``

You can configure a per-user upper bandwidth limit, in bits per second. This allows you to clamp your users
down to a sensible setting, if you are limited in bandwidth. This does not set a minimum bandwidth - the Mumble client
will check its own bandwidth setting and the server's bandwidth setting, and choose the lower of the two.

The default is 558000. Note that this is **per user**.

```text
bandwidth=558000
```

#### ``obfuscate``

By default, Mumble will show IP addresses in log file. In some situations you may find this unwanted behavior.
When ``obfuscate`` is set to ``true``, the ``mumble-server`` will randomize the IP addresses of connecting users.

Note that this setting currently only applies to IPv4 addresses, and that it is currently just a simple XOR with a random value.
It is probably trivially broken, if a user knows their IP address and can compare it with the obfuscated one assigend to
them.

Example:

```text
obfuscate=true
```

#### ``uname``

On Unix-like OSes, you can start ``mumble-server`` as root and have it drop privileges to another account such as ``mumble``.
The username must already exist on your system and have write access to the database, logfiles etc.

This option is ignored, if ``mumble-server`` is not started as root.

Example:

```text
uname=mumble
```

