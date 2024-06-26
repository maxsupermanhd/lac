package lac

import (
	"slices"
	"testing"
)

func Test_parseJSON(t *testing.T) {
	jin := []byte(`{"A": 2, "B": 2.5, "C": 2147483647, "D": 2147483649}`)
	cfg, err := FromBytesJSON(jin)
	if err != nil {
		t.Fatal(err)
	}
	n, ok := cfg.GetInt("A")
	if !ok {
		t.Fatal("A not ok")
	}
	if n != 2 {
		t.Fatal("A is wrong")
	}
	_, ok = cfg.GetInt("B")
	if ok {
		t.Fatal("B is ok")
	}
	_, ok = cfg.GetInt("C")
	if !ok {
		t.Fatal("C is not ok")
	}
	_, ok = cfg.GetInt("D")
	if !ok {
		t.Fatal("D is not ok")
	}
	n2, ok := cfg.GetInt64("D")
	if !ok {
		t.Fatal("D is not ok")
	}
	if n2 != 2147483649 {
		t.Fatal("D is wrong")
	}
	n3, ok := cfg.GetFloat64("B")
	if !ok {
		t.Fatal("B is not ok")
	}
	if n3 != 2.5 {
		t.Fatal("B is wrong")
	}
}

func Test_JSONArrays(t *testing.T) {
	jin := []byte(`{"A": ["hello", "world"], "B": [1, 2, 3]}`)
	cfg, err := FromBytesJSON(jin)
	if err != nil {
		t.Fatal(err)
	}
	a, ok := cfg.GetSliceString("A")
	if !ok {
		t.Fatal("a is not ok")
	}
	if slices.Compare(a, []string{"hello", "world"}) != 0 {
		t.Fatal("a is wrong")
	}
	b, ok := cfg.GetSliceInt("B")
	if !ok {
		t.Fatal("B is not ok")
	}
	if slices.Compare(b, []int{1, 2, 3}) != 0 {
		t.Fatal("B is wrong")
	}
}
