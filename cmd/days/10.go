package days

import (
	"code/util"
	"fmt"
)

type Day10 int

var day10 = Day10(10)

func (d Day10) Part1() (string, error) {
	s, err := util.File2Array("inputs/day10_part1.txt")
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

func (d Day10) Part2() (string, error) {
	s, err := util.File2Array("inputs/day10_part2.txt")
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

func (d Day10) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day10)
}
