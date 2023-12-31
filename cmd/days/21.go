package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day21 int

var day21 = Day21(21)

func (d Day21) Part1() (string, error) {
	s, err := util.File2Array("inputs/day21_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]string{}
	// start := []int{0, 0}
	for _, l := range s {
		if l != "" {
			data := strings.Split(l, "")
			grid = append(grid, data)

			for x := range data {
				if data[x] == "S" {
					// start = []int{y, x}
				}
			}
		}
	}

	for q := 0; q < 6; q++ {

	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day21) Part2() (string, error) {
	s, err := util.File2Array("inputs/day21_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {
			fmt.Println(x)
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day21) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day21)
}
