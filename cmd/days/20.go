package days

import (
	"code/util"
	"fmt"
)

type Day20 int

var day20 = Day20(20)

func (d Day20) Part1() (string, error) {
	s, err := util.File2Array("inputs/day20_part1.txt")
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

func (d Day20) Part2() (string, error) {
	s, err := util.File2Array("inputs/day20_part2.txt")
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

func (d Day20) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day20)
}
