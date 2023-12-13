package days

import (
	"code/util"
	"fmt"
)

type Day14 int

var day14 = Day14(14)

func (d Day14) Part1() (string, error) {
	s, err := util.File2Array("inputs/day14_part1.txt")
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

func (d Day14) Part2() (string, error) {
	s, err := util.File2Array("inputs/day14_part2.txt")
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

func (d Day14) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day14)
}
