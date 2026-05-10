package cntr_test

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func ExampleWeightedRandom_Select() {
	w := cntr.NewWeightedRandom[int]()
	w.Init(map[string]int{"pick": 10})
	fmt.Println(w.Select())
	// Output:
	// pick
}
