---
title: Using Glacier2 with Ice
date: 2019-10-27
weight: 100
---
**NOTE: Since Mumble 1.2.2 you can set `icesecretread` and `icesecretwrite` in your server configuration and use it as a password. This is a lot easier to set up and use than Glacier2.**

**Glacier2** is a Ice **routing and firewall utility**, and allows you to securely run the server on one machine and murmur on another. Note that if both server and client are on a secure LAN, you can just use `iptables` to protect the Ice port, which is a lot easier than setting up Glacier2.

The examples here assume that `1.2.3.4` is the public IP address of the server running Murmur. We're going to use the username `magic` with the password `pink`.

## Configuring Glacier2

Create a config file called `config.glacier2` and put the following in it:

```ini
Glacier2.Client.Endpoints=tcp -h 1.2.3.4 -p 4063
Glacier2.SessionTimeout=60
Glacier2.CryptPasswords=passwords.txt
```

Your endpoint host should be the public IP that you are running Glacier on. If you don't specify a client via `-h`, then Glacier will bind to all listening interfaces.

Then, create a password hash using the OpenSSL utility.

```bash
openssl passwd pink
```

this will spit out a hash, which looks something like `CTThafhdv9Lz2`

Create a file called `passwords.txt` containing:

```bash
magic CTThafhdv9Lz2
```

Start glacier2 as this:

```bash
glacier2router --Ice.Config=config.glacier2
```

You will need to have Ice installed ([download](https://zeroc.com/downloads/ice)). `glacier2router` is a binary that is located in `<location_of_Ice_installation>/bin/glacier2router.exe`.

## Configuring Murmur

There is nothing to do in murmur. Seriously. Leave the default setting of binding to `127.0.0.1` alone.

## Configuring Client (PHP)

This is where it starts getting slightly ugly. Note that this requires Ice >= 3.3.1, as Ice 3.3.0 has a bug in it which prevents this from working. The following is the adaptation necessary to `weblist.php` to get it to work:

```php
try {
  $router = $ICE->stringToProxy("Glacier2/router:tcp -p 4063 -h 1.2.3.4");
  $router = $router->ice_uncheckedCast("::Glacier2::Router")->ice_router(null);
  $session = $router->createSession("magic", "pink");
  $base = $ICE->stringToProxy("Meta:tcp -h 127.0.0.1 -p 6502")->ice_router($router);
  $meta = $base->ice_checkedCast("::Murmur::Meta")->ice_router($router);
  …
```

For each object you get a proxy to (including the return from `$meta->getServer`), you need to add `->ice_router($router)`

## Configuring Client (Ruby)

There is a set of classes for easily working with Ice directly and through Glacier [available at GitHub](https://github.com/cheald/murmur-manager/tree/master/interfaces/). However, if you want to do it manually, it’s not too hard.

```ini
glacierHost = "example.com"
glacierPort = 1234
user = "glacieruser"
pass = "glacierpass"
server_id = 1

prx = ic.stringToProxy("Glacier2/router:tcp -h #{glacierHost} -p #{glacierPort}")
router = ::Glacier2::RouterPrx::uncheckedCast(prx).ice_router(nil)
router.createSession(user, pass)
meta = Murmur::MetaPrx::checkedCast(ic.stringToProxy("Meta:tcp -h #{host} -p #{port}")).ice_router(router)
server = meta.getServer(server_id).ice_router(router)
```

For each object you get a proxy to (including the return from `Murmur::MetaPrx::getServer`), you need to add `#ice_router(router)`.

{{< aside >}}
This content released under [Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This content is based on {{< wiki Ice />}}.
{{< /aside >}}
