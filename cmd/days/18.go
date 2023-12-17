package days

import (
	"code/util"
	"fmt"
)

type Day18 int

var day18 = Day18(18)

func (d Day18) Part1() (string, error) {
	s, err := util.File2Array("inputs/day18_part1.txt")
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

func (d Day18) Part2() (string, error) {
	s, err := util.File2Array("inputs/day18_part2.txt")
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

func (d Day18) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day18)
}
