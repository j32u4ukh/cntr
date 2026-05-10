package cntr_test

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func ExampleWeightedRandom_Select() {
	orig := cntr.RandInt
	cntr.RandInt = func(n int) int { return 0 }
	defer func() { cntr.RandInt = orig }()

	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{"pick": 10})
	fmt.Println(w.Select())
	// Output:
	// pick
}
