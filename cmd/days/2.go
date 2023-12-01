package days

import (
	"code/util"
	"fmt"
)

type Day2 int

var day2 = Day2(2)

func (d Day2) Part1() (string, error) {
	s, err := util.File2Array("inputs/day2_part1.txt")
	if err != nil {
		return "", err
	}

	for _, x := range s {
		if x != "" {
			fmt.Println(x)
		}
	}

	return "Day2", nil
}

func (d Day2) Part2() (string, error) {
	s, err := util.File2Array("inputs/day2_part2.txt")
	if err != nil {
		return "", err
	}

	for _, x := range s {
		if x != "" {
			fmt.Println(x)
		}
	}

	return "Day2", nil
}

func (d Day2) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day2)
}
