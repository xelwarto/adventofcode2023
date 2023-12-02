package days

import (
	"code/util"
	"fmt"
)

type Day5 int

var day5 = Day5(5)

func (d Day5) Part1() (string, error) {
	s, err := util.File2Array("inputs/day5_part1.txt")
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

func (d Day5) Part2() (string, error) {
	s, err := util.File2Array("inputs/day5_part2.txt")
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

func (d Day5) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day5)
}
