package test

import (
	"testing"

	"github.com/j32u4ukh/cntr"
)

func TestRandomElementByte(t *testing.T) {
	original := cntr.RandInt
	defer func() { cntr.RandInt = original }()
	cntr.RandInt = func(n int) int {
		return n - 2
	}
	result := cntr.RandomElement([]byte{1, 2, 3, 4, 5})
	answer := byte(4)
	if result != answer {
		t.Errorf("Expected: %d, result: %d", answer, result)
	}
}

func TestRandomElementString(t *testing.T) {
	original := cntr.RandInt
	defer func() { cntr.RandInt = original }()
	cntr.RandInt = func(n int) int {
		return n - 2
	}
	result := cntr.RandomElement([]string{"a", "b", "c", "d"})
	answer := "c"
	if result != answer {
		t.Errorf("Expected: %s, result: %s", answer, result)
	}
}
