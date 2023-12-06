package days

import (
	"code/util"
	"fmt"
)

type Day6 int

var day6 = Day6(6)

func (d Day6) Part1() (string, error) {
	s, err := util.File2Array("inputs/day6_part1.txt")
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

// ##################################################################################
func (d Day6) Part2() (string, error) {
	s, err := util.File2Array("inputs/day6_part2.txt")
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

func (d Day6) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day6)
}
