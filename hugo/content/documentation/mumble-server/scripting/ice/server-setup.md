---
title: Mumble Server Configuration to Enable Ice
date: 2024-12-14
weight: -100
---
To enable the Ice interface in your `mumble-server.ini` configuration file, add or uncomment the following line

```ini
ice="tcp -h 127.0.0.1 -p 6502"
```

This will make the Ice backend listen via TCP on localhost (that's what the IP address `127.0.0.1` corresponds to) on port `6502`. It is possible to
change the IP address and the port to something else, but you need to be aware of two things:
- Once enabled, all connections made from the given IP address to the given port are granted access to the Mumble Ice interface. **There is no
   additional security checks** that the server performs on connection requests. Therefore, you need to make sure that the setup is in such a way that
   only authorized clients can ever make connections via that route. With the aforementioned defaults, anyone who can access the machine the server is
   running on can also access the Ice interface.
- If you change IP address and port in the server configuration, you need to adapt any further instructions to match your setup.

If you want to expose your Ice interface to connections made from outside the machine the server is on (i.e. anything on the internet), you have to
change the IP address to the public IP address of the server's machine. You can find this IP address by executing `hostname --all-ip-addresses` on
your server's machine. This command will usually return the `127.0.0.1` address we used above and then either one or two more addresses. Take one of
the latter ones.

Additionally, you can make use of the `icesecretread` and `icesecretwrite` configuration options. If set, they will require any Ice
read (write) requests to the server to supply the respective string as a password. If the password is not given or is wrong, the request will fail.
Note that in most cases, providing the write password also grants read access.

If any of these options is set to an empty value (to be contrasted with simply not setting the option at all), all associated requests (i.e. read
requests for `icesecretread` and write requests for `icesecretwrite`) will be denied, regardless of whether a password was provided or not. If this is
used to block read access, even correctly providing a write password won't unlock read accesses.

Now restart the server so the change takes effect.

To check if Ice in fact does listen, on UNIX type:

```bash
netstat -apn | grep 6502
```

and on Windows type:

```batch
netstat -an
```

and look for the process listening on port `6502`.

If the port is not being listened on, check the server's log. It should state enabling ice on startup. If it does not, something of your configuration
went wrong.

```text
MumbleServerIce: Endpoint "tcp -h 127.0.0.1 -p 6502" running
```

{{< aside >}}
This content released under [Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted.
{{< /aside >}}
