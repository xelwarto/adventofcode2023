package days

import (
	"code/util"
	"fmt"
)

type Day7 int

var day7 = Day7(7)

func (d Day7) Part1() (string, error) {
	s, err := util.File2Array("inputs/day7_part1.txt")
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
func (d Day7) Part2() (string, error) {
	s, err := util.File2Array("inputs/day7_part2.txt")
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

func (d Day7) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day7)
}
