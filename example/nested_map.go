package main

import (
	"fmt"
	"os"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

func registerNestedMapCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "nested-map",
		Short: "Example for cntr.NestedMap[K1, K2, V].",
		Run: func(*cobra.Command, []string) {
			runNestedMapDemo()
		},
	}
	root.AddCommand(cmd)
}

func runNestedMapDemo() {
	nm := cntr.NewNestedMap[int, int, int]()
	nm.Set(2, 2, 4)
	nm.Set(2, 3, 6)
	v2, err := nm.Get(2, 3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}
	fmt.Printf("value: %d\n", v2)
}
