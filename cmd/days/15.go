package days

import (
	"code/util"
	"fmt"
)

type Day15 int

var day15 = Day15(15)

func (d Day15) Part1() (string, error) {
	s, err := util.File2Array("inputs/day15_part1.txt")
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

func (d Day15) Part2() (string, error) {
	s, err := util.File2Array("inputs/day15_part2.txt")
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

func (d Day15) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day15)
}
