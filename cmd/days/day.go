package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

type Day interface {
	Part1() (string, error)
	Part2() (string, error)
	DayInt() int
}

var days []Day

var DayCmd = &cobra.Command{
	Use:   "day [day] [all|p1|p2]",
	Short: "Advent of Code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(days) > 0 {
			part := "all"
			if len(args) > 1 {
				part = args[1]
			}
			i, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			found := false
			for _, d := range days {
				if i == d.DayInt() {
					found = true
					fmt.Printf("Advent of Code 2023 - Day %v\n\n", d.DayInt())

					if part == "all" || part == "p1" {
						p1, err := d.Part1()
						if err != nil {
							log.Fatalf("Error: %v", err)
						}
						fmt.Printf("Part1 Answer: %s\n", p1)
					}

					if part == "all" || part == "p2" {
						p2, err := d.Part2()
						if err != nil {
							log.Fatalf("Error: %v", err)
						}
						fmt.Printf("Part2 Answer: %s\n", p2)
					}
				}
			}

			if !found {
				log.Fatalf("Error: requested day not found")
			}
		} else {
			cmd.Help()
		}
	},
}

func NewDay(d Day) {
	days = append(days, d)
}
