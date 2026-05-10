package test

import (
	"testing"

	"github.com/j32u4ukh/cntr"
)

type seqStrategy struct {
	vals []int
	i    int
}

func (s *seqStrategy) Next(totalWeight int) int {
	if s.i >= len(s.vals) {
		s.i = 0
	}
	v := s.vals[s.i]
	s.i++
	if v < 0 || v >= totalWeight {
		panic("test strategy: value out of range for totalWeight")
	}
	return v
}

func TestWeightedRandom_Select_empty(t *testing.T) {
	w := cntr.NewWeightedSelector()
	w.Init(nil)
	if w.Select() != "" {
		t.Fatal("empty init should select empty string")
	}
	w.Init(map[string]int{"x": 0})
	if w.Select() != "" {
		t.Fatal("only zero weights -> empty")
	}
}

func TestWeightedRandom_Select_deterministicLowSegment(t *testing.T) {
	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{"a": 30, "b": 70})
	w.SetStrategy(&seqStrategy{vals: []int{0, 5, 29}})
	if got := w.Select(); got != "a" {
		t.Fatalf("target 0 -> a, got %q", got)
	}
	if got := w.Select(); got != "a" {
		t.Fatalf("target 5 -> a, got %q", got)
	}
	if got := w.Select(); got != "a" {
		t.Fatalf("target 29 -> a, got %q", got)
	}
}

func TestWeightedRandom_Select_deterministicHighSegment(t *testing.T) {
	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{"a": 30, "b": 70})
	w.SetStrategy(&seqStrategy{vals: []int{30, 99}})
	if got := w.Select(); got != "b" {
		t.Fatalf("target 30 -> b, got %q", got)
	}
	if got := w.Select(); got != "b" {
		t.Fatalf("target 99 -> b, got %q", got)
	}
}

func TestWeightedRandom_Select_keySortOrder(t *testing.T) {
	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{"z": 10, "m": 10})
	w.SetStrategy(&seqStrategy{vals: []int{5}})
	if got := w.Select(); got != "m" {
		t.Fatalf("sorted keys m before z, cumulative m ends at 10; target 5 -> m, got %q", got)
	}
}

func TestWeightedRandom_Select_internalDirectRandomWhenNil(t *testing.T) {
	orig := cntr.RandInt
	defer func() { cntr.RandInt = orig }()
	cntr.RandInt = func(n int) int {
		if n != 1 {
			t.Fatalf("totalWeight should be 1, got n=%d", n)
		}
		return 0
	}
	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{"only": 1})
	if got := w.Select(); got != "only" {
		t.Fatalf("got %q", got)
	}
}
