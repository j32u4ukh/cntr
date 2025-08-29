package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func main() {
	// BinaryDataDemo()
	EmptyArrayBinaryDataDemo()
	// NestedMapDemo()
}

func BinaryDataDemo() {
	bd := cntr.NewBinaryData()
	bd.AddFloat32(3.2)
	bd.AddString("Hello, world!")
	bd.AddInt32(32)
	f32, _ := bd.PopFloat32()
	fmt.Printf("f32: %f\n", f32)
	bd.AddFloat64(6.4)
	bd.AddFloat64Array([]float64{0.618, 1.414, 1.618, 2.71828, 3.1415926})
	s, _ := bd.PopString()
	fmt.Printf("s: %s\n", s)
	i32, _ := bd.PopInt32()
	fmt.Printf("i32: %d\n", i32)
	f64, _ := bd.PopFloat64()
	fmt.Printf("f64: %f\n", f64)
	farray, _ := bd.PopFloat64Array()
	fmt.Printf("farray: %+v\n", farray)
}

func EmptyArrayBinaryDataDemo() {
	bd := cntr.NewBinaryData()
	bd.AddByteArray(nil)
	data, err := bd.PopByteArray()
	if err != nil {
		return
	}
	fmt.Printf("data: %+v\n", data)
}

func NestedMapDemo() {
	nm := cntr.NewNestedMap[int, int, int]()
	nm.Set(2, 2, 4)
	nm.Set(2, 3, 6)
	v2, err := nm.Get(2, 3)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("value: %d\n", v2)
}
