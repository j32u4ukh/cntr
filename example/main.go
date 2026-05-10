package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:           "example",
		Short:         "Run cntr container demos by subcommand.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	registerBinaryDataCommands(rootCmd)
	registerNestedMapCommands(rootCmd)
	registerBinaryTreeCommands(rootCmd)
	registerRandomCommands(rootCmd)
	registerWeightedRandomCommands(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
