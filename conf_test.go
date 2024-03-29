package lac

import (
	"testing"
)

func TestWalk(t *testing.T) {
	cfg := NewConf()
	cfg.Set("world!", "Hello")
	cfg.Set("test", "foo", "bar", "baz")
	cfg.Walk(func(k []string, v any) {
		t.Logf("%#v %#v", k, v)
	})
}

func TestGetSet(t *testing.T) {
	cfg := NewConf()
	cfg.Set("awd", "dwa")
	v, ok := cfg.Get("dwa")
	if !ok {
		t.Fatal("Can not get key")
	}
	if v != "awd" {
		t.Fatal("Value is incorrect")
	}
}

func TestSetSetGet(t *testing.T) {
	cfg := NewConf()
	cfg.Set("awd", "dwa")
	cfg.Set("fff", "dwa")
	v, ok := cfg.Get("dwa")
	if !ok {
		t.Fatal("Can not get key")
	}
	if v != "fff" {
		t.Fatal("Value is incorrect")
	}
}

func TestGetSetMapOverrite(t *testing.T) {
	cfg := NewConf()
	cfg.Set("test", "foo", "bar")
	v, ok := cfg.Get("foo")
	if !ok {
		t.Fatal("Can not get map key")
	}
	vv, ok := v.(map[string]any)
	if !ok {
		t.Fatal("Value has incorrect type")
	}
	vv["bar"] = "failed"
	v, ok = cfg.Get("foo", "bar")
	if !ok {
		t.Fatal("Can not get key")
	}
	if v != "test" {
		t.Fatal("Value changed inside the tree without Set")
	}
}

func TestGetKeys(t *testing.T) {
	cfg := NewConf()
	cfg.Set("test", "foo")
	cfg.Set(42, "bar")
	cfg.Set(420, "baz")
	cfg.Set(false, "example", "something")
	k, ok := cfg.GetKeys()
	if !ok {
		t.Fatal("GetKeys failed")
	}
	if len(k) != 4 {
		t.Fatal("Wrong number of keys")
	}
	t.Log(k)
	t.Log(cfg.GetKeys("example"))
	t.Log(cfg.GetKeys("wrong"))
}
