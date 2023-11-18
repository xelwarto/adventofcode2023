package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of Code 2023 - Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day 1")
	},
}
