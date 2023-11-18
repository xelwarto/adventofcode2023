package cmd

import (
	"code/cmd/demo"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Root Command
var rootCmd = &cobra.Command{
	Use:   "main.go",
	Short: "Advent of Code 2023",
}

func init() {
	// Load demo commands for testing code
	rootCmd.AddCommand(demo.DemoCmd)

	// Load event day commands
	rootCmd.AddCommand(dayCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
