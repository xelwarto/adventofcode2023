package days

import (
	"code/util"
	"fmt"
)

type Day17 int

var day17 = Day17(17)

func (d Day17) Part1() (string, error) {
	s, err := util.File2Array("inputs/day17_part1.txt")
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

func (d Day17) Part2() (string, error) {
	s, err := util.File2Array("inputs/day17_part2.txt")
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

func (d Day17) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day17)
}
