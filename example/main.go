package main

import (
	"encoding/binary"
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	bs := cntr.NumberToBytes2(int32(20), binary.LittleEndian)
	fmt.Printf("bs: %+v\n", bs)
}
