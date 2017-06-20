Gosettings: for applications and libraries
==========================================

R Pratap Chakravarthy <br/>
prataprc@gmail.com <br/>
[https://github.com/prataprc/gosettings](https://github.com/prataprc/gosettings)

---

In a nutshell
=============

* Easy to learn and easy to use settings for your libraries and applications.
* Settings object is represented as ``map[string]interface{}`` object.
* Settings can be marshalled to JSON or un-marshalled from JSON.
* Possible to add more formats for marshalling and un-marshalling settings.
* All methods exported on settings object are immutable, except ``Mixin``.
* Stable APIs, existing APIs are not going to change.

---

Settings are key,value pairs
============================

* Keys are strings, identify the settings name (aka: configuration parameter).
* Values can be nil, bool, number, string, array, property-map.
* Nested settings are not allowed, as in, if value is a property-map it
shall not be interpreted as {key,value} pairs of settings.

```go
setts := Settings{
    "level":      "info",
    "flags":      "",
    "file":       "",
    "timeformat": "2006-01-02T15:04:05",
    "prefix":     "[%v]",
}
```

Accessing values
================

There are several accessor methods available on ``Settings`` to get the
concrete value mapped to a settings key.

```go
setts.Bool("mvcc.enable")       // return the boolean value for key.
setts.Float64("memutilization") // return the float64 value for key.
setts.Int64("minkeysize") // return the int64 value for key.
setts.Uint64("maxvalzie") // return the uint64 value for key.
setts.String("log.file")  // return the string value for key.
setts.Strings("services") // shall parse value as comma separated string items.
```

**NOTE**: To encode large numbers that can fit within int64 and uint64,
settings value can be encoded as decimal strings:

```go
setts := {"epoch": "1125899906842624"}
fmt.Println("%d", setts.Uint64("epoch"))
```

Namespace: for topology of settings
===================================

While building applications that use several components, and each component
allows itself to be configured via predefined set settings parameters,
it becomes imperative that these components and sub-components need to be
organized at application level.

This can be done by organizing the settings keys in namespace format.

```go
setts := make(Settings)
setts["storage.llrb.maxkeylen"] = 1024
setts["storage.llrb.maxvallen"] = 1024
```

Component and sub-components are separated by ``.``

Filtering settings
==================

To filter out settings that are only relevant for ``llrb``:

```go
setts := make(Settings)
setts["storage.numvbuckets"] = 8
setts["storage.memalloc"] = 1000000
setts["storage.llrb.maxkeylen"] = 1024
setts["storage.llrb.maxvallen"] = 1024

llrbsetts := setts.Section("llrb.") // Section
```

Use the Section() API, llrbsetts will be:

```text
Settings{"storage.llrb.maxkeylen": 1024, "storage.llrb.maxvallen": 1024}
```

Triming settings
================

While passing llrbsetts to the llrb component, it may not expect the
``storage.llrb.`` prefixed to its settings key-name. To trim them away:

```go
llrbsetts = setts.Section("storage.llrb.").Trim("storage.llrb.")
```

Now, llrbsettings will just be:

```text
Settings{"maxkeylen": 1024, "maxvallen": 1024}
```

Thank you
=========

If gosettings sounds useful please check out the following links.

[Project README](https://github.com/prataprc/gosettings). <br/>
[API doc](https://godoc.org/github.com/prataprc/gosettings). <br/>
[Please contribute](https://github.com/prataprc/gosettings/issues). <br/>
