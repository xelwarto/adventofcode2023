package days

import (
	"code/util"
	"fmt"
)

type Day21 int

var day21 = Day21(21)

func (d Day21) Part1() (string, error) {
	s, err := util.File2Array("inputs/day21_part1.txt")
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
