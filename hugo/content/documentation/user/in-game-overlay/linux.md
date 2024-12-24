---
title: "Mumble Overlay on Linux"
---
To use the Overlay in an application you start that appliaction through the mumble-overlay binary:

```bash
mumble-overlay gameexecutable
```

The mumble-overlay binary is a utility for an *LD preload* which you can also execute manually:

```bash
LD_PRELOAD=/path/to/libmumble.so.1.1 gameexecutable
```
