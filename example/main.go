package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	bm := cntr.NewBikeyMap[int, string, float32]()
	bm.Add(1, "a", 1)
	bm.Add(2, "b", 4)
	bm.Add(3, "c", 9)
	iter := bm.GetIterator()
	for iter.HasNext() {
		value := iter.Next()
		fmt.Printf("value: %+v\n", value)
	}
}
