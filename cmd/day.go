package cmd

import (
	"code/cmd/days"

	"github.com/spf13/cobra"
)

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Advent of Code 2023",
}

func init() {
	dayCmd.AddCommand(days.Day1Cmd)
}
