package days

import (
	"code/util"
	"fmt"
)

type Day25 int

var day25 = Day25(25)

func (d Day25) Part1() (string, error) {
	s, err := util.File2Array("inputs/day25_part1.txt")
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

func (d Day25) Part2() (string, error) {
	s, err := util.File2Array("inputs/day25_part2.txt")
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

func (d Day25) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day25)
}
