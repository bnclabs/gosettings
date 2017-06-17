Settings and configurations for golang application
==================================================

[![Build Status](https://travis-ci.org/prataprc/gosettings.png)](https://travis-ci.org/prataprc/gosettings)
[![Coverage Status](https://coveralls.io/repos/prataprc/gosettings/badge.png?branch=master&service=github)](https://coveralls.io/github/prataprc/gosettings?branch=master)
[![GoDoc](https://godoc.org/github.com/prataprc/gosettings?status.png)](https://godoc.org/github.com/prataprc/gosettings)

* Build components, libraries and applications that are easy to configure.
* Easy to learn and easy to use settings for your applications.
* Settings object is represented as ``map[string]interface{}`` object.
* Settings can be marshalled to JSON or unmarshalled from JSON.
* Possible to add more formats for marshalling and unmarshalling settings.
* All methods exported on settings object are immutable, except ``Mixin``.

Settings as Key value pairs
---------------------------

Golang map is the chosen structure to pass settings, hold them within
components and interpret them when ever it is needed. To create new
settings object:

```go
    setts := make(Settings)
```

Subsequently ``setts`` can be populated with {key,value} pairs, where each
{key,value} pair correspond to a single configuration parameter. To make
most out of ``Settings`` it is better to avoid nested values, as in:

```go
    setts["llrb"] = Settings{"maxkeylen": 1024, "maxvallen": 1024}
```

Instead compose the settings as:

```go
    setts := make(Settings)
    setts["llrb.maxkeylen"] = 1024
    setts["llrb.maxvallen"] = 1024
```

Although the former style is quite natural to manage a tree of settings at
component level and at sub-component level, it can quickly become complex when
we start passing settings object around the application. The later style
attempt to preserve the topology of settings by using the dot-prefixed
namespaces. There are three APIs available to filter out component level
settings and merge it with container namespace, ``Section``, ``Trim``,
and ``AddPrefix``:

```go
    setts := make(Settings)
    setts["numvbuckets"] = 8
    setts["memalloc"] = 1000000
    setts["llrb.maxkeylen"] = 1024
    setts["llrb.maxvallen"] = 1024
    llrbsetts := setts.Section("llrb.")
```

setts.Section API will filter out setting keys that are prefixed by ``llrb.``,
it will look-like:

```text
    llrbsetts <==> Settings{"llrb.maxkeylen": 1024, "llrb.maxvallen": 1024}
```

While passing llrbsettings to llrb component, it may not expect the ``llrb.``
prefixed to its settings key. To trim them away:

```go
    llrbsettings := setts.Section("llrb.").Trim("llrb.")
```

Now, llrbsettings will just be:

```text
    llrbsetts <==> Settings{"maxkeylen": 1024, "maxvallen": 1024}
```

In case llrb component provide default set of configuration parameter,
to merge these settings to application-level settings object, use
AddPrefix:

```go
    appsetts := llrbsetts.AddPrefix("llrb.")
```

appsetts will look like:

```text
    appsetts <==> Settings{"llrb.maxkeylen": 1024, "llrb.maxvallen": 1024}
```

**Settings from json**

Most often, settings are obtained from JSON text. One of the reason for
using ``map[string]interface{}`` as the underlying data-structure is to keep
it friendly for JSON. To initialize Settings from JSON:

```go
    var setts Settings
    json.Unmarshal(data, &setts)
```

**Accessors**

With ``map[string]interface{}`` settings value are resolved only during run
time.  There are several helper functions that type check and extract the
settings value.

* Bool(key string), return the boolean value for key.
* Float64(key string), return the float64 value for key.
* Int64(key string) return the int64 value for key.
* Uint64(key string) return the uint64 value for key.
* String(key string) return the string value for key.
* Strings(key string) shall parse value as comma separated string items.

To encode large numbers that can fit within int64 and uint64, settings value
can be encoded as decimal strings - ``{"epoch": "1125899906842624"}.

**Panics**

Settings API don't return error, instead it creates panic. This is because
settings are mostly part of bootstrapping process and are expected to be
clean and suggested by developers. If panics become un-avoidable please use
[panic/recover](https://blog.golang.org/defer-panic-and-recover). We are
listing some of the cases when panic can happened.

* When accessing a settings key which is not present in the map.
* When using one of the typed accessors, if underlying value does not match
the accessor type.
