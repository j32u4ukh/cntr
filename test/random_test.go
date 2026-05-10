package test

import (
	"testing"

	"github.com/j32u4ukh/cntr"
)

func TestRandInt_range(t *testing.T) {
	const n = 50
	for i := 0; i < 1000; i++ {
		v := cntr.RandInt(n)
		if v < 0 || v >= n {
			t.Fatalf("RandInt(%d) out of range: %d", n, v)
		}
	}
}

func TestDirectRandom_forwardsRandInt(t *testing.T) {
	orig := cntr.RandInt
	defer func() { cntr.RandInt = orig }()
	want := 3
	wantN := false
	cntr.RandInt = func(n int) int {
		wantN = (n == 10)
		return want
	}
	d := &cntr.DirectRandom{}
	if got := d.Next(10); got != want || !wantN {
		t.Fatalf("DirectRandom.Next: got %d wantN=%v", got, wantN)
	}
}

func TestBufferedRandom_valuesInRange(t *testing.T) {
	buf := cntr.NewBufferedRandom[int](8)
	buf.Init(0, 37)
	for i := 0; i < 64; i++ {
		v := buf.Next()
		if v < 0 || v >= 37 {
			t.Fatalf("BufferedRandom.Next out of range: %d (upper=37)", v)
		}
	}
}

func TestBufferedRandom_singleValueRange(t *testing.T) {
	buf := cntr.NewBufferedRandom[int](5)
	buf.Init(0, 1)
	for i := 0; i < 40; i++ {
		if buf.Next() != 0 {
			t.Fatal("Init(0,1) should always yield 0")
		}
	}
}
