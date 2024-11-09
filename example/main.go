package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	nm := cntr.NewNestedMap[int, int, int]()
	nm.Set(2, 2, 4)
	nm.Set(2, 3, 6)
	v2, err := nm.Get(2, 3)
	if err != nil{
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("value: %d\n",v2)
}
