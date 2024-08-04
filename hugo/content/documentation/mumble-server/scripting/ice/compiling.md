---
title: Compiling Ice
date: 2019-10-27
weight: 10
---

On some platforms there are no officially supported binaries available. In that case you will either have to change your
platform or compile Ice yourself.

## Building Ice on Redhat/CentOS machines

**Note: 3.3.1 is no longer the current version of Ice. Feel free to try this guide with the new version and update it if
it works, or fix it if it does not.**

### 1. Download and unpack ICE

```bash
wget http://www.zeroc.com/download/Ice/3.3/Ice-3.3.1.tar.gz
tar -xzf Ice-3.3.1.tar.gz
```

### 2. Compile the ICE CPP bindings

(these are required for all other bindings)

You will need `mcpp-devel` from the Zeroc website installed to compile. <http://www.zeroc.com/download.html> - Ctrl-F
and look for `mcpp-devel`. There is a big package of various Ice RPMs to download. You will install a few dependencies,
along with the `mcpp-devel package`.

```bash
cd Ice-3.3.1/cpp
make
```

(wait 20 minutes)

```bash
sudo make install
```

### 3. Compile and install the Ice bindings for your preferred language

For example, to install Ruby bindings:

```bash
cd ../rb
make
sudo make install
```

### 4. Export the paths for your newly-installed libraries

These will be different for each language - check the `INSTALL` or `README` files in each language’s subdirectory for
exact instructions.

For Ruby:

```bash
export RUBYLIB=/opt/Ice-3.3.1/ruby:$RUBYLIB
export LD_LIBRARY_PATH=/opt/Ice-3.3.1/lib:LD_LIBRARY_PATH
```

If you don't want to always have to keep running those export lines, also add them to your ~/.bashrc:

```bash
export RUBYLIB=/opt/Ice-3.3.1/ruby:$RUBYLIB
export LD_LIBRARY_PATH=/opt/Ice-3.3.1/lib:LD_LIBRARY_PATH
```

### Verifying

At this point, Ice should be available to your language (in this case, Ruby):

```bash
$ irb
irb(main):001:0> require 'Ice'
=> true
```

### Generate the Slice file for your language

To generate it for ruby, we use the `slice2rb` program, which is in the `Ice/cpp/bin` directory. Similar binaries for
your language of choice will be there, too.

```bash
wget -O Murmur.ice "https://raw.github.com/mumble-voip/mumble/master/src/murmur/Murmur.ice"
../cpp/bin/slice2rb Murmur.ice
cp Murmur.rb #{MANAGER_ROOT}/vendor/ice
```

Congratulations! Ice should be set up and fully functional.

{{< aside >}} This content released under
[Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This
content is based on {{< wiki Ice />}}. {{< /aside >}}
