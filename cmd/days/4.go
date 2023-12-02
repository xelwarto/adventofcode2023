package days

import (
	"code/util"
	"fmt"
)

type Day4 int

var day4 = Day4(4)

func (d Day4) Part1() (string, error) {
	s, err := util.File2Array("inputs/day4_part1.txt")
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

func (d Day4) Part2() (string, error) {
	s, err := util.File2Array("inputs/day4_part2.txt")
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

func (d Day4) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day4)
}
