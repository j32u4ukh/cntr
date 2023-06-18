package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	bm := cntr.NewBikeyMap[string, int, float32]()
	bm.Add("a", 1, 1.0)
	bm.Add("b", 2, 1.414)
	bm.Add("c", 3, 1.71)

	value, _ := bm.GetByKey1("b")
	fmt.Printf("value: %+v", value)
}
