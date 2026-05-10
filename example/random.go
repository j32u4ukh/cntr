package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

func registerRandomCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Examples for cntr.RandInt and random strategies.",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "rand-int",
			Short: "Print several cntr.RandInt samples in [0, bound).",
			Run: func(*cobra.Command, []string) {
				const bound = 20
				fmt.Printf("RandInt(%d):", bound)
				for i := 0; i < 8; i++ {
					fmt.Printf(" %d", cntr.RandInt(bound))
				}
				fmt.Println()
			},
		},
		&cobra.Command{
			Use:   "buffered-shuffle",
			Short: "BufferedShuffle pre-fills then serves shuffled ints per Next.",
			Run: func(*cobra.Command, []string) {
				buf := cntr.NewBufferedShuffle(8)
				tw := 100
				fmt.Printf("Next(%d):", tw)
				for i := 0; i < 16; i++ {
					fmt.Printf(" %d", buf.Next(tw))
				}
				fmt.Println()
			},
		},
		&cobra.Command{
			Use:   "direct-strategy",
			Short: "DirectRandom used as cntr.RandomStrategy.",
			Run: func(*cobra.Command, []string) {
				strat := &cntr.DirectRandom{}
				fmt.Printf("DirectRandom.Next(50) x 5:")
				for i := 0; i < 5; i++ {
					fmt.Printf(" %d", strat.Next(50))
				}
				fmt.Println()
			},
		},
	)
	root.AddCommand(cmd)
}
