package days

import (
	"code/util"
	"fmt"
)

type Day9 int

var day9 = Day9(9)

func (d Day9) Part1() (string, error) {
	s, err := util.File2Array("inputs/day9_part1.txt")
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

func (d Day9) Part2() (string, error) {
	s, err := util.File2Array("inputs/day9_part2.txt")
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

func (d Day9) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day9)
}
