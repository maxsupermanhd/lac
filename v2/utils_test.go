package lac

import (
	"testing"
)

func TestCopyMap(t *testing.T) {
	m := map[string]any{
		"hello": "world",
		"nested": map[string]any{
			"key":        "value",
			"answer":     42,
			"digits":     []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			"interfaces": []any{"test1", "test2"},
		},
		"funny": []any{69, 420},
	}
	cfg := NewConf()
	cfg.CopyTree(m)
	digitsAi, ok := cfg.Get("nested", "digits")
	if !ok {
		t.Fatal("Failed to get digitsA")
	}
	digitsA, ok := digitsAi.([]any)
	if !ok {
		t.Fatal("DigitsA have wrong type")
	}

	digitsA[3] = 999

	digitsBi, ok := cfg.Get("nested", "digits")
	if !ok {
		t.Fatal("Failed to get digitsB")
	}
	digitsB, ok := digitsBi.([]any)
	if !ok {
		t.Fatal("DigitsB have wrong type")
	}

	digitsTarget := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range digitsB {
		if digitsB[i] != digitsTarget[i] {
			t.Fatal("Digits changed without Set")
		}
	}
}
