---
title: Mumble Server Ice Scripting with Python
date: 2024-12-14
weight: -40
---

This guide demonstrates how you can get started with using the [Ice](https://zeroc.com/ice) scripting interface of the Mumble server. Effectively, it
is a summary of [the official docs](https://doc.zeroc.com/ice/latest/release-notes/using-the-python-distribution) tuned specifically to the Mumble use
case.

Prerequisite for this to work is that you have [enabled Ice]({{< ref "../server-setup.md" >}}) in your server's configuration.

## Setup

First, you have to install the Ice package for Python via
```bash
pip install zeroc-ice
```

Additionally, you need the slice file that defines the Ice interface. Make sure you obtain the slice file that corresponds to the actual server
version you are using. You can find the relevant slice files on GitHub by checking out the branch corresponding to your release series (i.e. `1.5.x`
for Mumble 1.5.x) and then locating the `MumbleServer.ice` file under `src/murmur/`. Note that the slice file for development versions of Mumble can
always be found on the `master` branch. See also the [list of branches](https://github.com/mumble-voip/mumble/branches) in the GitHub repository.

The relevant slice file may also be part of your installation already but its exact install location can vary. Typically, it can be found under e.g.
`/usr/local/share/mumble-server/` or `/usr/share/mumble-server/` on Unix systems.

{{< aside >}}
In older versions (< 1.5), the legacy name `Murmur` was used instead of `MumbleServer`. That means that the relevant file was called `Murmur.ice` and
also all mentions of `MumbleServer` in the following instructions have to be replaced with `Murmur` in those cases.
{{< /aside >}}


## Generating stubs

Before being able to use Python to access a Mumble server, you need to generate the relevant Python stubs from the slice file. In order to do so, you
can use the `slice2py` executable that should have been as part of the `zeroc-ice` Python package. To do so, execute
```bash
slice2py <path_to_slice_file>
```
where `<path_to_slice_file>` is the path to the slice file. So if you have the slice file in the current directory, the call would be
```bash
slice2py MumbleServer.ice
```

This will generate a `MumbleServer` directory and a `MumbleServer_ice.py` in the current directory.


## First steps

Once everything has been set up and all stubs have been generated, you can start with the actual Python scripting. The only thing to take care of is
that the generated stubs have to be in your `PYTHONPATH` in order to be able to import the `MumbleServer` module. The easiest way to accomplish this
is to have the stubs generated in the same directory in which your script lives.

The following script illustrates the essentials of establishing a connection to your server and how to use the Ice interface:
```python3
import sys
import Ice
import MumbleServer


def main():
    with Ice.initialize(sys.argv) as communicator:
        # Establish the connection to the server - make sure the used IP and port match the ones
        # in your server's configuration
        # "Meta" identifies the object adapter to use and the remaining string after the colon
        # specifies the endpoint to connect to. In this case, a TCP endpoint where
        # -h specifies the host (IP)
        # -p specifies the port
        # -t specifies the timeout of the connection in milliseconds
        base = communicator.stringToProxy("Meta:tcp -h 127.0.0.1 -p 6502 -t 60000")

        # Cast the proxy to a Meta object
        meta = MumbleServer.MetaPrx.checkedCast(base)

        assert meta is not None

        # Iterate over all existing virtual servers
        for server in meta.getAllServers():
            # Print status of current server making use of the id() and isRunning() methods
            print(
                f"Found server with ID {server.id()} ({'online' if server.isRunning() else 'offline'})"
            )

            if server.isRunning():
                # For running servers, print how long they have been running already
                uptime = server.getUptime()
                print(f"  -> Has been online for {uptime} seconds")


if __name__ == "__main__":
    main()
```

Note that the first object you have access to after connecting to the server is a `Meta` object, which can be seen as the entity managing all
configured virtual servers. It can be used to start/stop existing servers and to create new ones (or delete existing ones).

In most cases you'll want to use the `Meta` object to get a handle of the actual `Server` object corresponding to the virtual server you are
interested in. This can then be used to manipulate this specific server (e.g. set configuration options or querying connected users).


## Interface documentation

The documentation of what kind of functions can be called (and with what parameters) on `Meta` or `Server` objects can be found in the slice file
itself. Search for `interface Meta` and `interface Server` respectively.


## Using passwords

If the server has been configured to require a password for read and/or write requests via Ice (cmp. {{< ref "../server-setup.md" >}}), you have to
pass the correct password in form of a context object containing the password. The context object is a dictionary and the password is provided as the
value for the `secret` key of said dictionary. All Ice functions related to the Mumble server take the context as an optional last argument when being
called.

As an example consider the following script which sets the SuperUser password on a server that has been configured with `icesecret="mypw"`:
```python3
import sys
import Ice
import MumbleServer


def main():
    with Ice.initialize(sys.argv) as communicator:
        base = communicator.stringToProxy("Meta:tcp -h 127.0.0.1 -p 6502 -t 60000")
        meta = MumbleServer.MetaPrx.checkedCast(base)

        assert meta is not None

        context = {"secret": "mypw"}

        # We just assume the server we want to modify has ID 1
        server = meta.getServer(1, context)
        assert server is not None

        # Not providing the context with the correct password would make this
        # request fail with an InvalidSecretException
        server.setSuperuserPassword("Strong password", context)


if __name__ == "__main__":
    main()
```


## Troubleshooting

### `InvalidSecretException`

Your request has been blocked because a password is required to perform the given request and you didn't provide one or the one you provided is
incorrect. Alternatively, it could be that the server has completely disabled any requests of this kind (read or write requests).
