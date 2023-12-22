package days

import (
	"code/util"
	"fmt"
)

type Day22 int

var day22 = Day22(22)

func (d Day22) Part1() (string, error) {
	s, err := util.File2Array("inputs/day22_part1.txt")
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

func (d Day22) Part2() (string, error) {
	s, err := util.File2Array("inputs/day22_part2.txt")
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

func (d Day22) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day22)
}
