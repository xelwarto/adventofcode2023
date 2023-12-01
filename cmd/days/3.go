package days

import (
	"code/util"
	"fmt"
)

type Day3 int

var day3 = Day3(3)

func (d Day3) Part1() (string, error) {
	s, err := util.File2Array("inputs/day3_part1.txt")
	if err != nil {
		return "", err
	}

	for _, x := range s {
		if x != "" {
			fmt.Println(x)
		}
	}

	return "Day3", nil
}

func (d Day3) Part2() (string, error) {
	s, err := util.File2Array("inputs/day3_part2.txt")
	if err != nil {
		return "", err
	}

	for _, x := range s {
		if x != "" {
			fmt.Println(x)
		}
	}

	return "Day3", nil
}

func (d Day3) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day3)
}
