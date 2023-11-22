package cmd

import (
	"code/cmd/days"
	"code/cmd/demo"
	"log"

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
	rootCmd.AddCommand(days.DayCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
