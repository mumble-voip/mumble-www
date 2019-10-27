---
title: Mumble Server Ice Scripting with PHP
date: 2019-10-27
weight: -30
---
Setup documentation that you may want to check out prior to this article: [PHP Setup]({{< relref "setup" >}}), [Ice]({{< relref ".." >}})

With IcePHP establishing the connection to the interface **differs** between the Ice versions 3.3 and prior and 3.4 and later.

## Ice <= 3.3

There's an example script using the '''Ice 3.3''' approach (defining the ice.slice directive in the PHP settings) included in the source; have a look at `/mumble/blob/1.2.4/scripts/icedemo.php`.

The establishing, minimum code, is:

```php
Ice_loadProfile();
// initialize ice connection
global $ICE;
$base = $ICE->stringToProxy('Meta:tcp -h 127.0.0.1 -p 6502');
$meta = $base->ice_checkedCast("::Murmur::Meta");
```

## Ice >= 3.4

First, you will have to **generate PHP code** from the slice definitions `.ice` file. With Ice >= 3.4 installed, use the `slice2php` executable to generate it.

For your PHP code, you’ll have to have the Ice.php and other libs (scripts provided by zeroc) in your PHPs include path to include them.

```php
require_once 'Ice.php';
require_once 'Murmur.php';
$ICE = Ice_initialize();
$meta = Murmur_MetaPrxHelper::checkedCast($ICE->stringToProxy('Meta:tcp -h 127.0.0.1 -p 6502'));
```

Where `Murmur.php` is the generated file.

{{< aside >}}
This content released under [Creative Commons Attribution Share Alike](http://creativecommons.org/licenses/by-sa/2.5/) unless otherwise noted. This content is based on {{< wiki Ice />}}.
{{< /aside >}}
