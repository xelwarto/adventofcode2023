package days

import (
	"code/util"
	"fmt"
)

type Day23 int

var day23 = Day23(23)

func (d Day23) Part1() (string, error) {
	s, err := util.File2Array("inputs/day23_part1.txt")
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

func (d Day23) Part2() (string, error) {
	s, err := util.File2Array("inputs/day23_part2.txt")
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

func (d Day23) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day23)
}
