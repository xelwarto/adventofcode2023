package days

import (
	"code/util"
	"fmt"
)

type Day13 int

var day13 = Day13(13)

func (d Day13) Part1() (string, error) {
	s, err := util.File2Array("inputs/day13_part1.txt")
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

func (d Day13) Part2() (string, error) {
	s, err := util.File2Array("inputs/day13_part2.txt")
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

func (d Day13) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day13)
}
