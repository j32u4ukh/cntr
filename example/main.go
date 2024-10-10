package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	bd := cntr.NewBinaryData()
	m := map[string][]byte{
		"A": {1, 2, 3},
		"B": {4, 5, 6, 7},
		"C": {8, 9},
	}
	bd.AddMapStringByteArray(m)
	data, err := bd.PopMapStringByteArray()
	if err != nil {
		fmt.Printf("Failed to read MapStringByteArray")
		return
	}
	fmt.Printf("data: %+v\n", data)
}
