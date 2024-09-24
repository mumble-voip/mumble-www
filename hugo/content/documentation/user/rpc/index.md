---
title: Mumble Client RPC subcommands
date: 2024-09-23
---
Various [RPC](https://en.wikipedia.org/wiki/Remote_procedure_call) subcommands allow users to issue requests to a running Mumble client starting with version 1.3.0.

| Command  | Description         |
| -------- | ------------------- |
| `mute`   | Mute self |
| `unmote`    | Unmute self |
| `togglemute`   | Toggle between mute and unmute |
| `deaf`  | Deafen self |
| `undeaf`  | Undeafen self |
| `toggledeaf`  | Toggle between deaf and undeaf |
| `starttalking`  | Start talking (like Push-To-Talk) |
| `stoptalking`  | Stop talking (like Push-To-Talk) |

Use ``mumble rpc --help`` to get a list of valid commands.

### Example call
```text
$ mumble rpc mute
$ mumble rpc unmute
```

{{< aside >}}
This content released under [Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted.}.
{{< /aside >}}
