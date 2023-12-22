package days

import (
	"code/util"
	"fmt"
)

type Day24 int

var day24 = Day24(24)

func (d Day24) Part1() (string, error) {
	s, err := util.File2Array("inputs/day24_part1.txt")
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

func (d Day24) Part2() (string, error) {
	s, err := util.File2Array("inputs/day24_part2.txt")
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

func (d Day24) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day24)
}
