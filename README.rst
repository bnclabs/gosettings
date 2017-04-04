[![Build Status](https://travis-ci.org/prataprc/gosettings.png)](https://travis-ci.org/prataprc/gosettings)
[![Coverage Status](https://coveralls.io/repos/prataprc/gosettings/badge.png?branch=master&service=github)](https://coveralls.io/github/prataprc/gosettings?branch=master)
[![GoDoc](https://godoc.org/github.com/prataprc/gosettings?status.png)](https://godoc.org/github.com/prataprc/gosettings)

Configure golang libraries and applications by supplying settings defined and
implemented under this package.

* Settings object is represented as map[string]interface{} object.
* Settings can be marshalled to JSON or unmarshalled from JSON using
  golang's stdlib encoding/json.
* All methods exported on settings object are immutable, except Mixin.

**How to use Mixin**

.. code-block:: go

    setts = make(Settings).Mixin(settings1, settings2)
