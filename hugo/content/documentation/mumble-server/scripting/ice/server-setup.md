---
title: Mumble Server Configuration to Enable Ice
date: 2019-10-27
weight: -100
---

To enable the Ice interface in your `murmur.ini` configuration file, we recommend to first disable DBus by commenting
out:

```ini
dbus=session
```

Then enable Ice for ''localhost'' on port ''6502'' by adding or uncommenting:

```ini
ice="tcp -h 127.0.0.1 -p 6502"
```

Now restart Murmur so the change takes effect.

To check if Ice in fact does listen, on UNIX type:

```bash
netstat -apn | grep 6502
```

and on Windows type:

```batch
netstat -an
```

and look for the process listening on port `6502`.

If the port is not being listened on, check Murmurs log. It should state enabling ice on startup. If it does not,
something of your configuration went wrong.

```text
MurmurIce: Endpoint "tcp -h 127.0.0.1 -p 6502" running
```

{{< aside >}} This content released under
[Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This
content is based on {{< wiki Ice />}}. {{< /aside >}}
