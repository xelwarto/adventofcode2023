package days

import (
	"code/util"
	"fmt"
)

type Day19 int

var day19 = Day19(19)

func (d Day19) Part1() (string, error) {
	s, err := util.File2Array("inputs/day19_part1.txt")
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

func (d Day19) Part2() (string, error) {
	s, err := util.File2Array("inputs/day19_part2.txt")
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

func (d Day19) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day19)
}
