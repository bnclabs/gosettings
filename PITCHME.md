@title[gosettings]

@snap[midpoint slide1]
<h1>gosettings</h1>
@size[80%](For applications and libraries)
@snapend

@snap[south-east author-box]
@fa[envelope](prataprc@gmail.com - R Pratap Chakravarthy) <br/>
@fa[github](https://github.com/bnclabs/gosettings) <br/>
@snapend

---

In a nutshell
=============

@ul

- Easy to learn and easy to use settings for your libraries and applications.
- Settings object is represented as @color[blue](map[string]interface{}) object.
- Settings can be marshalled to ``JSON`` or un-marshalled from ``JSON``.
- Possible to add more formats for marshalling and un-marshalling settings.
- All methods exported on settings object are immutable, except @color[blue](Mixin).
- Stable APIs, existing APIs are not going to change.

@ulend

---

{key,value} pairs
=================

```go
example := Settings{
    "level":      "info",
    "flags":      "",
    "file":       "",
    "timeformat": "2006-01-02T15:04:05",
    "prefix":     "[%v]",
}
```

@ul

- Keys are strings, identify the settings name (aka: configuration parameter).
- Values can be @color[blue](nil, bool, number, string, array, property-map).
- Nested settings are not allowed, as in, if value is a property-map it shall not be interpreted as {key,value} pairs of settings.

@ulend


---

Accessing Values
================

```golang
setts.Bool("mvcc.enable")
setts.Int64("minkeysize")
setts.Uint64("maxvalsize")
setts.String("log.file")
setts.Strings("services")

setts := {"epoch": "1125899906842624"}
fmt.Println("%d", setts.Uint64("epoch"))
```

@[1](return @color[yellow](mvcc.enable) as boolean value)
@[2](return @color[yellow](minkeysize) as int64 value)
@[3](return @color[yellow](maxvalsize) as uint64 value)
@[4](return @color[yellow](log.file) as string value)
@[5](parse @color[yellow](services) for comma seprated text, and return array of string)

@[7](To encode large numbers that can fit within int64 and uint64, settings value can be encoded as decimal strings)
@[8](stdout: @color[yellow](1125899906842624))

---

A topology of settings
======================

```go
setts := make(Settings)
setts["storage.llrb.maxkeylen"] = 1024
setts["storage.llrb.maxvallen"] = 1024
```

@ul[mt20]

- While building applications that use several components, and each component allows itself to be configured via predefined set parameters, it becomes imperative that these components and sub-components need to be organized at application level. This can be done by organizing the settings keys in namespace format.
- Component and sub-components are separated by a ``.`` dot.

@ulend

---

Filtering settings
==================

Say we have a component @color[blue](storage.llrb) and we need to filter
out parameters relevant for storage.llrb, use the **Section()** API.

```go
setts := make(Settings)
setts["storage.numvbuckets"] = 8
setts["storage.memalloc"] = 1000000
setts["storage.llrb.maxkeylen"] = 1024
setts["storage.llrb.maxvallen"] = 1024

llrbsetts := setts.Section("llrb.") // Section
```

@[7](@color[yellow](llrbsetts) will be @color[cyan](Settings{"storage.llrb.maxkeylen": 1024, "storage.llrb.maxvallen": 1024}))

---

Trimming settings
================

While passing llrbsetts to the llrb component, it may not
expect the @color[blue](storage.llrb.) prefixed to its settings key-name.

```go
llrbsetts = setts.Section("storage.llrb.").Trim("storage.llrb.")
```

@[1](Now, @color[yellow](llrbsetts) will just be @color[cyan](Settings{"maxkeylen": 1024, "maxvallen": 1024}))

---

From JSON
=========

Most often, settings are obtained from JSON text. One of the reason for
using @color[blue](map[string]interface{}) as the underlying data-structure is
to keep it friendly for JSON. To initialize Settings from JSON:

```go
var setts Settings
json.Unmarshal(data, &setts)
```

Can't get simpler than that !

---

Addprefix and Mixin
===================

When default settings from different components need to be consolidated into
application level settings.

```go
import comp1
import comp2

comp1setts := comp1.Defaultsettings()
comp2setts := comp1.Defaultsettings()
appsetts := make(Settings).Mixin(
    comp1setts.AddPrefix("comp1."), comp2setts.AddPrefix("comp2."),
)
```

---

Thank you
=========

If gosettings sounds useful please check out the following links.

@fa[book] [Project README](https://github.com/bnclabs/gosettings). <br/>
@fa[code] [API doc](https://godoc.org/github.com/bnclabs/gosettings). <br/>
@fa[github] [Please contribute](https://github.com/bnclabs/gosettings/issues). <br/>
