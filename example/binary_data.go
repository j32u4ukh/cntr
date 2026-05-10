package main

import (
	"fmt"
	"os"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

func registerBinaryDataCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "binary-data",
		Short: "Examples for cntr.BinaryData (stack-style read/write).",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "full",
			Short: "Push float, string, ints, float64 array; pop in FIFO order.",
			Run: func(*cobra.Command, []string) {
				runBinaryDataFull()
			},
		},
		&cobra.Command{
			Use:   "empty-array",
			Short: "Serialize and read back a nil byte slice.",
			Run: func(*cobra.Command, []string) {
				runBinaryDataEmptyArray()
			},
		},
	)
	root.AddCommand(cmd)
}

func runBinaryDataFull() {
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

func runBinaryDataEmptyArray() {
	bd := cntr.NewBinaryData()
	bd.AddByteArray(nil)
	data, err := bd.PopByteArray()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}
	fmt.Printf("data: %+v\n", data)
}
