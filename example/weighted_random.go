package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

func registerWeightedRandomCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "weighted-random",
		Short: "Examples for cntr.WeightedRandom (BST over prefix weights).",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "select",
			Short: "Init from map, sample Select multiple times.",
			Run: func(*cobra.Command, []string) {
				runWeightedRandomSelectExample()
			},
		},
	)
	root.AddCommand(cmd)
}

func runWeightedRandomSelectExample() {
	w := cntr.NewWeightedSelector()
	w.Init(map[string]int{
		"apple":    35,
		"banana":   35,
		"coconut":  30,
		"skipped":  0,
		"_ignored": 0,
	})
	fmt.Printf("total implied by Init (apple+banana+coconut sorted by key): cumulative order a,b,c\n")
	for i := 0; i < 12; i++ {
		fmt.Printf("sample %02d -> %s\n", i, w.Select())
	}
}
