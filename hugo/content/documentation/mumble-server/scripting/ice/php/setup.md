---
title: Mumble Server Ice Scripting with PHP - Setup
date: 2019-10-27
weight: -10
license: cc-by-sa 2.5
basedon: https://wiki.mumble.info/wiki/Ice
---

Note: This documentation is quite old and may need some verification or updating.

How to convert Murmur.ice to Murmur.php for Ice >= 3.4.0 :
<http://doc.zeroc.com/display/Ice/slice2php+Command-Line+Options>

## How to setup Ice 3.4 for PHP with Apache on Linux (general)

If your Linux distribution offers a binary packet for Ice with PHP (usually the name contains Ice and php) you can skip
everything but naming the Murmur.ice slice file in the php.ini (see below). If there is no prepared package you'll have
to try to find binaries for your system or {{< wiki "Ice#Compiling_Ice" "compile" />}} Ice yourself and add the
extension to PHP and tell PHP where to find the Murmur.ice file.

To add the IcePHP extension to PHP, first check that the file IcePHP.so for linux is in your php extensions folder
specified in your php.ini as

```ini
extension_dir = /usr/lib/php5/extensions
```

If it is not, get the corresponding files from [ZeroC's downloads page](http://www.zeroc.com/download.html).

Then either in your php.ini file or in your `/etc/php.d` or `/etc/php5/conf.d` folder in `ice.ini`, add the line

```ini
extension=IcePHP.so
```

At least the Linux RPMs will do this automatically, so check that you're not doing it a second time.

Second, you have to tell the PHP parser where to find the <abbr title="Specification Language Ice">slice</abbr> file.

Add

```ini
ice.slice = /path/to/Murmur.ice
```

to your `php.ini` or other config file (e.g. `ice.ini`).

### Troubleshooting

If you encounter problems, check your Apache log.

If it tells you the PHP extension was compiled on an older api, you have to compile the IcePHP.so from source.

Download `Ice-3.4.2.tar.gz` from [ZeroCs downloads page](http://www.zeroc.com/download.html), untar, cd into PHP, as
written in the `INSTALL` file export the `ICE_HOME` environment variable pointing to your Ice install dir. If you
installed it with an RPM, type

```bash
export ICE_HOME=/usr
```

then `make`, and in the `lib` folder, there will be your `IcePHP.so` file.

One common error is

```text
PHP Warning:  PHP Startup: skipping dictionary ::Murmur::UserInfoMap - unsupported key type in Unknown on line 0
```

This is caused because the `.ice` file is slightly incompatible with older versions of php-ice. Edit the `Murmur.ice`
file and find the following lines

```ice
/** User information map.
  * Older versions of ice-php can't handle enums as keys. If you are using one of these, replace 'UserInfo' with 'byte'.
  */
```

and perform the substitution mentioned.

## How to Setup Ice for PHP with Apache on Debian/Ubuntu ==

**Note:** For PHP Ice, there are major differences between the Ice versions 3.3 and 3.4. In more detail: 3.3 uses `.ice`
definition files while 3.4 uses generated PHP files.

Thus, depending on the version you are installing, you will have to follow different configuration steps.

At the moment (2012-05-06) the current Debian stable is `squeeze`, for which the official repositories provide Ice 3.3.
For Ubuntu on the other hand, the current stable version 12.04 (LTS) provides Ice 3.4 in its repositories. So make sure
you follow the matching instructions.

### Short, expert version

- If `php-zeroc-ice` version >= `3.4`: `apt-get install php-zeroc-ice`, then add `/usr/share/Ice-3.4.2/php/lib` to your
  PHPs include path.
- If `php-zeroc-ice` version <= `3.3`: `apt-get install php-zeroc-ice`, then add the `ice.slice` parameter to with value
  `<pathto>/murmur.ice` to your PHPs configuration.

### Setup

For the following guide, a set up Apache2 and PHP environment, and a working install of Murmur is assumed.

#### Step 1 - PHP Ice extension

First we need to install the Ice library for PHP. Execute the following in a root shell:

```bash
apt-get update
apt-get install php-zeroc-ice
```

#### Step 2 - PHP Setup

Now we need to tell PHP to load the IcePHP extension. Open the file ''/etc/php5/apache2/php.ini'' for editing:

```bash
vim /etc/php5/apache2/php.ini
```

OR:

```bash
nano -w /etc/php5/apache2/php.ini
```

Paste the folowing in the dynamic extensions section of the file:

```ini
extension=IcePHP.so
```

#### Step 3 - Reload and Check

Everything should be set up so let's restart the servers so they load the updated config files.

Restart your Apache2 daemon:

```bash
/etc/init.d/apache2 restart
```

Make a file named `phpinfo.php` in your web root and paste the following:

```php
<?php phpinfo(); ?>
```

Open the page in your browser, you should be greeted with PHP's information page. Search the page for `Ice version` to
verify that the `IcePHP` extension has loaded.

If you don't see `IcePHP`’s version and config info on this page, verify that the file `IcePHP.so` exists in PHP’s
`extension_dir` (this will also be in your PHP info page). Once you’re sure that the IcePHP extension has loaded, delete
the `phpinfo.php` file for security reasons.

Now we will take a look in the mumble-server log to see if Murmur's Ice interface is listening:

```bash
tail -n10 /var/log/mumble-server/mumble-server.log
```

If you find a line similar to the following everything is fine and you can now communicate to Murmur via its Ice
interface.

```text
<W>2012-03-24 13:37:11.316 MurmurIce: Endpoint "tcp -h 127.0.0.1 -p 6502" running
```

#### Step 4 - Setting up Murmur to provide Ice interface

Now open the file ''/etc/mumble-server.ini'' for editing:

```bash
vim /etc/mumble-server.ini
```

OR:

```bash
nano -w /etc/mumble-server.ini
```

Comment the following line to disable {{< wiki DBus />}}:

```ini
#dbus=system
```

Uncomment or paste the following line to enable the Ice interface:

```ini
ice="tcp -h 127.0.0.1 -p 6502"
```

Restart your Mumble server:

```bash
/etc/init.d/mumble-server restart
```

After restarting the daemon Murmur will now listen for Ice requests.

Now you can install and use any of the PHP based [[3rd Party Applications]] compatible with IcePHP version 3.4.

## How to setup Ice for PHP with XAMPP on Windows

### Pre Setup

Download and install the following packages:

- XAMPP for Windows <http://www.apachefriends.org>
- ZeroC Ice <https://zeroc.com/downloads/ice>

Now configure and start XAMPP.

### Create developer environment for Windows

- Open System in Control Panel
- On the Advanced tab, click Environment Variables
- Add the Ice Path to `%PATH%`

Select the name `PATH` on user variables and edit. If the name `PATH` does not exist, select new. Enter now the `PATH`
to your Ice installation dir.

For example: `%PROGRAMFILES(x86)%\ZeroC\Ice-3.4.0\bin;%PATH%`

Now logout from Windows or restart. After reboot check if included Ice in our environment. Open the CMD and try

```batch
slice2php
```

The Output return the `slice2php` help.

### Configure XAMPP/PHP

Copy the File php_ice.dll from `<your path>\ZeroC\Ice-3.4.0\bin` to `<your_path>\xampplite\php\ext`

Edit the `php.ini` in `<your_path>\xampplite\php` and add

```ini
extension=php_ice.dll
ice.slice=<mumble_path>\Murmur.ice
```

Now start XAMPP and check the `phpinfo()`.

## Using different ice.slice on same host

Sometimes you run two servers of mumble on the same host and you cannot load two slices with different mumble versions
on the same time.

The solution is to use ICE profiles, that will require extra files and unfortunately modification of scripts that run
the default profile.

Description on Debian etch with apache2, PHP as fcgid and cli.

### Creating profiles

- Lets create profile directory (as root/superuser user), I have created mine at `/etc/php5/ice/`
- Then create directories per profile (it will make life easier later) `/etc/php5/ice/murmur.1.1.x`
  `/etc/php5/ice/murmur.1.2.x`
- To each of the corresponding directory I have placed the `Murmur.ice` file provided with the installs. You will notice
  that later upgrades will just consist of copying new files (instead of renaming).
- time to make profiles.ini file that ICE will use: `vi /etc/php5/ice/profiles.ini`
- Insert there below code:

```ini
[Murmur11x]
slice=/etc/php5/ice/murmur.1.1.x/Murmur.ice

[Murmur12x]
slice=/etc/php5/ice/murmur.1.2.x/Murmur.ice
```

This way you got two profiles, first (`Murmur11x`) is loading slice file for older murmur setups, and the other
(`Murmurm12x`) loads newer slices. Below example of directory structure:

```text
/etc/php5/ice
|-- murmur.1.1.x
|   `-- Murmur.ice
|-- murmur.1.2.x
|   `-- Murmur.ice
`-- profiles.ini
```

### `IcePHP.ini`

- Now time to alter `/etc/php5/cgi/conf.d/IcePHP.ini`, so it will look like this

```ini
extension = IcePHP.so
ice.profiles=/etc/php5/ice/profiles.ini
```

Notice that we are adding ice.profile setting, so that ice knows where to search for profiles.

Default profile is named `__default__` and you better avoid to make it.

I got no idea about setting up default ice.slice entry (experience told me to force editing php files anyway, read
below)

#### Changing PHP scripts

Next alter the source code of the php scripts you use. It consists of just altering the one function that tells ICE to
load profile. By default the function is ran without parameter, and it looks like:

```php
Ice_loadProfile();
```

Now we will tell that function to load specific profile that it was designed for. Below example modifications to load
different profiles for given ice profile: for older mumble:

```php
Ice_loadProfile('Murmur11x'); // load profile for scripts that talk with murmur 1.1.x
```

or for new mumble:

```php
Ice_loadProfile('Murmur12x'); // load profile for scripts that talk with murmur 1.2.x
```

Notice that single quotes are required.

### Checking out if it works

Remember to restart apache (your configuration may require it)

See php/apache error logs when something goes wrong.

If you get error like `unknown operation getUsers invoked on proxy of type ::Murmur::Server` then it means you use old
`Murmur.ice` file. This happens when you try to connect using `Murmur.ice` from version 1.1.x to Murmur server 1.2.x,
which require `Murmur.ice` 1.2.x. Update `Murmur.ice` on web server and restart it and then it should work.

### Further steps

I suggest you do the same with `/etc/php5/cli/conf.d` file and scripts that you run from console (crontab).

Altering `php.ini` `error_reporting` is advised to disable generating massive amount of warning messages:

```ini
error_reporting = E_ALL & ~E_NOTICE & ~E_WARNING
```

You can also set up config files for those profile - more info at
<http://www.zeroc.com/doc/Ice-3.3.1/manual/IcePHP.29.3.html>

{{< aside >}} This content released under
[Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This
content is based on {{< wiki Ice />}}. {{< /aside >}}
