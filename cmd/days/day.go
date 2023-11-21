package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var DayCmd = &cobra.Command{
	Use:   "day",
	Short: "Advent of Code 2023",
}

type Day struct {
	Command *cobra.Command
	Part1   func() (string, error)
	Part2   func() (string, error)
}

func addDay(i int) *Day {
	s := strconv.Itoa(i)
	d := Day{}
	d.Command = &cobra.Command{
		Use:   s,
		Short: fmt.Sprintf("Advent of Code 2023 - Day %s", s),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Advent of Code 2023 - Day %s\n\n", s)

			p1, err := d.Part1()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Part 1 Answer: %v\n", p1)

			p2, err := d.Part2()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Part 2 Answer: %v\n", p2)
		},
	}
	d.Part1 = func() (string, error) {
		return "", nil
	}
	d.Part2 = func() (string, error) {
		return "", nil
	}
	DayCmd.AddCommand(d.Command)
	return &d
}
