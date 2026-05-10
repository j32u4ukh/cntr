package test

import (
	"testing"

	"github.com/j32u4ukh/cntr"
)

func TestWeightedRandom_Select_empty(t *testing.T) {
	w := cntr.NewWeightedRandom[int]()
	w.Init(nil)
	if w.Select() != "" {
		t.Fatal("empty init should select empty string")
	}
	w.Init(map[string]int{"x": 0})
	if w.Select() != "" {
		t.Fatal("only zero weights -> empty")
	}
}

func TestWeightedRandom_Select_singleKey(t *testing.T) {
	w := cntr.NewWeightedRandom[int]()
	w.Init(map[string]int{"only": 100})
	for i := 0; i < 50; i++ {
		if got := w.Select(); got != "only" {
			t.Fatalf("want only, got %q", got)
		}
	}
}

func TestWeightedRandom_Select_twoKeys_inRange(t *testing.T) {
	w := cntr.NewWeightedRandom[int]()
	w.Init(map[string]int{"a": 30, "b": 70})
	for i := 0; i < 200; i++ {
		k := w.Select()
		if k != "a" && k != "b" {
			t.Fatalf("unexpected key %q", k)
		}
	}
}

func TestWeightedRandom_Select_keySortOrder_bucketBoundary(t *testing.T) {
	w := cntr.NewWeightedRandom[int]()
	w.Init(map[string]int{"z": 10, "m": 10})
	for i := 0; i < 300; i++ {
		k := w.Select()
		if k != "m" && k != "z" {
			t.Fatalf("unexpected key %q", k)
		}
	}
}
