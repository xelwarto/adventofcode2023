package days

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Day1Cmd = &cobra.Command{
	Use:   "1",
	Short: "Advent of Code 2023 - Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Advent of Code 2023 - Day 1\n\n")
	},
}