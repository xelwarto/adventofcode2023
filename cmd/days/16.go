package days

import (
	"code/util"
	"fmt"
)

type Day16 int

var day16 = Day16(16)

func (d Day16) Part1() (string, error) {
	s, err := util.File2Array("inputs/day16_part1.txt")
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

func (d Day16) Part2() (string, error) {
	s, err := util.File2Array("inputs/day16_part2.txt")
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

func (d Day16) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day16)
}
