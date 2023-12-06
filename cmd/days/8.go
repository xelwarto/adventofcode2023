package days

import (
	"code/util"
	"fmt"
)

type Day8 int

var day8 = Day8(8)

func (d Day8) Part1() (string, error) {
	s, err := util.File2Array("inputs/day8_part1.txt")
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
func (d Day8) Part2() (string, error) {
	s, err := util.File2Array("inputs/day8_part2.txt")
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

func (d Day8) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day8)
}
