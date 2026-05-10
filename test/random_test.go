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

func TestBufferedShuffle_valuesInRange(t *testing.T) {
	buf := cntr.NewBufferedShuffle(8)
	tw := 37
	for i := 0; i < 64; i++ {
		v := buf.Next(tw)
		if v < 0 || v >= tw {
			t.Fatalf("BufferedShuffle.Next out of range: %d (tw=%d)", v, tw)
		}
	}
}

func TestBufferedShuffle_refillUsesRandInt(t *testing.T) {
	orig := cntr.RandInt
	defer func() { cntr.RandInt = orig }()
	var calls int
	cntr.RandInt = func(n int) int {
		calls++
		if n != 7 {
			t.Fatalf("unexpected totalWeight n=%d", n)
		}
		return 2
	}
	buf := cntr.NewBufferedShuffle(3)
	for i := 0; i < 3; i++ {
		if buf.Next(7) != 2 {
			t.Fatal("expected patched RandInt value")
		}
	}
	if calls != 3 {
		t.Fatalf("first refill: RandInt calls=%d", calls)
	}
	buf.Next(7)
	if calls != 6 {
		t.Fatalf("after second refill: RandInt calls=%d", calls)
	}
}
