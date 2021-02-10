Settings and configurations for golang application
--------------------------------------------------

[![Build Status](https://travis-ci.org/bnclabs/gosettings.png)](https://travis-ci.org/bnclabs/gosettings)
[![Coverage Status](https://coveralls.io/repos/bnclabs/gosettings/badge.png?branch=master&service=github)](https://coveralls.io/github/bnclabs/gosettings?branch=master)
[![GoDoc](https://godoc.org/github.com/bnclabs/gosettings?status.png)](https://godoc.org/github.com/bnclabs/gosettings)
[![GitPitch](https://gitpitch.com/assets/badge.svg)](https://gitpitch.com/bnclabs/gosettings/master?grs=github&t=white)

* Easy to learn and easy to use settings for golang libraries and applications.
* Settings object is represented as `map[string]interface{}` object.
* Settings can be marshalled to JSON or un-marshalled from JSON.
* Possible to add more formats for marshalling and un-marshalling settings.
* All methods exported on settings object are immutable, except `Mixin`.
* Stable APIs, existing APIs are not going to change.

Settings as Key value pairs
---------------------------

Golang map is the chosen structure for passing around settings, holding them
within components and interpret them when ever needed. To create new
settings object:

```go
setts := make(Settings)
```

Subsequently `setts` can be populated with {key,value} pairs, where each
{key,value} pair correspond to a single configuration parameter. To make
most out of `Settings` it is better to avoid nested values, as in:

```go
// not recommended
setts["llrb"] = Settings{
    "maxkeylen": 1024,
    "maxvallen": 1024,
}
```

Instead compose settings as:

```go
    // recommended
    setts := make(Settings)
    setts["llrb.maxkeylen"] = 1024
    setts["llrb.maxvallen"] = 1024
```

Although the former style is quite natural to manage a tree of settings at
component and sub-component level, it can quickly become complex when
we start passing settings object around the application. While the later
style encourages flat map of {key,value} pairs, we can still preserve the
topology of settings by using the dot-separated namespaces. There are three
APIs available to filter out component level settings and merge it with
container namespace, `Section`, `Trim`, and `AddPrefix`:

```go
    setts := make(Settings)
    setts["numvbuckets"] = 8
    setts["memalloc"] = 1000000
    setts["llrb.maxkeylen"] = 1024
    setts["llrb.maxvallen"] = 1024

    llrbsetts := setts.Section("llrb.") // Section
```

setts.Section API will filter out setting keys that are prefixed by `llrb.`,
it will look-like:

```text
llrbsetts <==> Settings{"llrb.maxkeylen": 1024, "llrb.maxvallen": 1024}
```

While passing llrbsettings to the llrb component, it may not expect the
`llrb.` prefixed to its settings key-name. To trim them away:

```go
llrbsettings := setts.Section("llrb.").Trim("llrb.") // Trim away "llrb."
```

Now, llrbsettings will just be:

```text
llrbsetts <==> Settings{"maxkeylen": 1024, "maxvallen": 1024}
```

In case, the llrb component is going to provide default set of
configuration parameters, and we want to merge them with our
application-level settings object, then, use AddPrefix:

```go
appsetts := llrbsetts.AddPrefix("llrb.")
```

appsetts will look like:

```text
appsetts <==> Settings{"llrb.maxkeylen": 1024, "llrb.maxvallen": 1024}
```

**Settings from json**

Most often, settings are obtained from JSON text. One of the reason for
using `map[string]interface{}` as the underlying data-structure is to keep
it friendly for JSON. To initialize Settings from JSON:

```go
var setts Settings
json.Unmarshal(data, &setts)
```

Can't get simpler than that !

**Accessors**

With `map[string]interface{}`, settings value are resolved only during run
time.  There are several helper functions that type check and extract the
settings value.

* Bool(key string), return the boolean value for key.
* Float64(key string), return the float64 value for key.
* Int64(key string) return the int64 value for key.
* Uint64(key string) return the uint64 value for key.
* String(key string) return the string value for key.
* Strings(key string) shall parse value as comma separated string items.

**NOTE**: To encode large numbers that can fit within int64 and uint64,
settings value can be encoded as decimal strings
Eg: `{"epoch": "1125899906842624"}`.

Panic and Recovery
------------------

Settings API don't return error, instead it creates panic. This is because
settings are mostly part of bootstrapping process and are expected to be
clean and suggested by developers. If panics become un-avoidable please use
[panic/recover](https://blog.golang.org/defer-panic-and-recover). We are
listing some of the cases when panic can happened.

* When accessing a settings key which is not present in the map.
* When using one of the typed accessors, if underlying value does not match
  the accessor type.

How to contribute
-----------------

* Pick an issue, or create an new issue. Provide adequate documentation for
  the issue.
* Assign the issue or get it assigned.
* Work on the code, once finished, raise a pull request.
* Gosettings is written in [golang](https://golang.org/), hence expected to
  follow the global guidelines for writing go programs.
* As of now, branch `master` is the development branch.
