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

func ExampleNewBufferedRandom() {
	buf := cntr.NewBufferedRandom[int](2)
	buf.Init(0, 1)
	fmt.Println(buf.Next(), buf.Next())
	// Output:
	// 0 0
}
