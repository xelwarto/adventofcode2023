package days

import (
	"code/util"
	"fmt"
)

type Day12 int

var day12 = Day12(12)

func (d Day12) Part1() (string, error) {
	s, err := util.File2Array("inputs/day12_part1.txt")
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

func (d Day12) Part2() (string, error) {
	s, err := util.File2Array("inputs/day12_part2.txt")
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

func (d Day12) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day12)
}
