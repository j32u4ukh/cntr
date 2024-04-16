package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	a1 := cntr.NewArray(1, 2, 3, 4, 5)
	a2 := a1.Clone()
	a2.Append(9)
	fmt.Printf("a1: %+v\n", a1)
	fmt.Printf("a2: %+v\n", a2)
	iter := a2.GetIterator()
	for iter.HasNext(){
		e := iter.Next()
		fmt.Printf("e: %v\n", e)
	}
	fmt.Println("End")
}
