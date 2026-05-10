package cntr_test

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func ExampleDirectRandom_Next() {
	orig := cntr.RandInt
	cntr.RandInt = func(n int) int { return n - 1 }
	defer func() { cntr.RandInt = orig }()

	strat := &cntr.DirectRandom{}
	fmt.Println(strat.Next(50))
	// Output:
	// 49
}

func ExampleNewBufferedShuffle() {
	orig := cntr.RandInt
	cntr.RandInt = func(n int) int { return 1 }
	defer func() { cntr.RandInt = orig }()

	buf := cntr.NewBufferedShuffle(2)
	fmt.Println(buf.Next(10), buf.Next(10))
	// Output:
	// 1 1
}
