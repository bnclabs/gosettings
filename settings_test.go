package gosettings

import "testing"
import _ "fmt"
import "reflect"

func TestSection(t *testing.T) {
	setts := Settings{
		"section1.param1": 10,
		"section1.param2": 20,
		"section2.param1": 30,
		"section2.param2": 40,
	}
	ref := Settings{
		"section1.param1": 10,
		"section1.param2": 20,
	}
	section := setts.Section("section1")
	if !reflect.DeepEqual(ref, section) {
		t.Fatalf("expected %v, got %v", ref, section)
	}
}

func TestAddPrefix(t *testing.T) {
	setts := Settings{
		"param1": 10,
		"param2": 20,
	}
	ref := Settings{
		"section.param1": 10,
		"section.param2": 20,
	}
	section := setts.AddPrefix("section.")
	if !reflect.DeepEqual(ref, section) {
		t.Fatalf("expected %v, got %v", ref, section)
	}
}

func TestTrim(t *testing.T) {
	setts := Settings{
		"section1.param1": 10,
		"section1.param2": 20,
		"section2.param1": 30,
		"section2.param2": 40,
	}
	ref := Settings{
		"param1": 10,
		"param2": 20,
	}
	trimmed := setts.Section("section1").Trim("section1.")
	if !reflect.DeepEqual(ref, trimmed) {
		t.Fatalf("expected %v, got %v", ref, trimmed)
	}
}

func TestFilter(t *testing.T) {
	setts := Settings{
		"section1.param1": 10,
		"section1.param2": 20,
		"section2.param1": 30,
		"section2.param2": 40,
	}
	ref := Settings{
		"section1.param1": 10,
		"section2.param1": 30,
	}
	filtered := setts.Filter("param1")
	if !reflect.DeepEqual(ref, filtered) {
		t.Fatalf("expected %v, got %v", ref, filtered)
	}
}

func TestMixin(t *testing.T) {
	setts1 := Settings{"section1.param1": 10}
	setts2 := map[string]interface{}{"section1.param2": 20}
	setts3 := Settings{"section2.param1": 30}
	setts4 := Settings{"section2.param2": 40}
	setts := make(Settings).Mixin(setts1, setts2, setts3, setts4)
	ref := Settings{
		"section1.param1": 10,
		"section1.param2": 20,
		"section2.param1": 30,
		"section2.param2": 40,
	}
	if !reflect.DeepEqual(ref, setts) {
		t.Fatalf("expected %v, got %v", ref, setts)
	}
}

func TestBool(t *testing.T) {
	setts := Settings{"param1": true, "param2": false, "param3": "string"}
	if v := setts.Bool("param1"); v != true {
		t.Fatalf("expected %v, got %v", true, v)
	} else if v := setts.Bool("param2"); v != false {
		t.Fatalf("expected %v, got %v", false, v)
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.Bool(key)
	}
	checkpanic("param3")
	checkpanic("param4")
}

func TestInt64(t *testing.T) {
	setts := Settings{
		"float64": float64(10), "float32": float32(10),
		"uint": uint(10), "uint64": uint64(10), "uint32": uint32(10),
		"uint16": uint16(10), "uint8": uint8(10),
		"int": int(10), "int64": int64(10), "int32": int32(10),
		"int16": int16(10), "int8": int8(10), "string": "10",
	}
	ref := int64(10)
	for key := range setts {
		if v := setts.Int64(key); v != ref {
			t.Fatalf("for key %v, expected %v, got %v", key, ref, v)
		}
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.Int64(key)
	}
	setts = setts.Mixin(map[string]interface{}{"notnum": "string"})
	checkpanic("notnum")
	checkpanic("missing")
}

func TestUint64(t *testing.T) {
	setts := Settings{
		"float64": float64(10), "float32": float32(10),
		"uint": uint(10), "uint64": uint64(10), "uint32": uint32(10),
		"uint16": uint16(10), "uint8": uint8(10),
		"int": int(10), "int64": int64(10), "int32": int32(10),
		"int16": int16(10), "int8": int8(10), "string": "10",
	}
	ref := uint64(10)
	for key := range setts {
		if v := setts.Uint64(key); v != ref {
			t.Fatalf("for key %v, expected %v, got %v", key, ref, v)
		}
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.Uint64(key)
	}
	setts = setts.Mixin(map[string]interface{}{"notnum": "string"})
	checkpanic("notnum")
	checkpanic("missing")
}

func TestFloat64(t *testing.T) {
	setts := Settings{
		"float64": float64(10), "float32": float32(10),
		"uint": uint(10), "uint64": uint64(10), "uint32": uint32(10),
		"uint16": uint16(10), "uint8": uint8(10),
		"int": int(10), "int64": int64(10), "int32": int32(10),
		"int16": int16(10), "int8": int8(10),
	}
	ref := float64(10)
	for key := range setts {
		if v := setts.Float64(key); v != ref {
			t.Fatalf("for key %v, expected %v, got %v", key, ref, v)
		}
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.Float64(key)
	}
	setts = setts.Mixin(map[string]interface{}{"notnum": "string"})
	checkpanic("notnum")
	checkpanic("missing")
}

func TestString(t *testing.T) {
	setts := Settings{"param": "value"}
	if v := setts.String("param"); v != "value" {
		t.Fatalf("for key %v, expected %v, got %v", "param", "value", v)
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.String(key)
	}
	setts = setts.Mixin(map[string]interface{}{"notstr": 10})
	checkpanic("notstr")
	checkpanic("missing")
}

func TestStrings(t *testing.T) {
	setts := Settings{"param": "value1, ,value2"}
	if v := setts.Strings("param"); len(v) != 2 {
		t.Fatalf("expected 2, got %v", len(v))
	} else if v[0] != "value1" {
		t.Fatalf("for key %v, expected %v, got %v", "param", "value1", v)
	} else if v[1] != "value2" {
		t.Fatalf("for key %v, expected %v, got %v", "param", "value2", v)
	}

	checkpanic := func(key string) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		setts.Strings(key)
	}
	setts = setts.Mixin(map[string]interface{}{"notstr": 10})
	checkpanic("notstr")
	checkpanic("missing")
}
