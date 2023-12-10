package days

import (
	"code/util"
	"fmt"
)

type Day11 int

var day11 = Day11(11)

func (d Day11) Part1() (string, error) {
	s, err := util.File2Array("inputs/day11_part1.txt")
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

func (d Day11) Part2() (string, error) {
	s, err := util.File2Array("inputs/day11_part2.txt")
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

func (d Day11) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day11)
}
