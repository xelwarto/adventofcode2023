package days

import (
	"fmt"
)

type Day6 int

var day6 = Day6(6)

func (d Day6) Part1() (string, error) {
	total := 0
	inputs := [][]int{{51, 377}, {69, 1171}, {98, 1224}, {78, 1505}}

	totals := []int{}
	for _, i := range inputs {
		c := 0
		t := i[0]
		d := i[1]

		for x := 0; x <= t; x++ {
			r := 1 * x
			y := t - x

			dist := r * y
			if dist > d {
				c++
			}
		}
		totals = append(totals, c)
	}

	for _, t := range totals {
		if total == 0 {
			total = t
		} else {
			total = total * t
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##################################################################################
func (d Day6) Part2() (string, error) {
	total := 0
	inputs := [][]int{{51699878, 377117112241505}}

	totals := []int{}
	for _, i := range inputs {
		c := 0
		t := i[0]
		d := i[1]

		for x := 0; x <= t; x++ {
			r := 1 * x
			y := t - x

			dist := r * y
			if dist > d {
				c++
			}
		}
		totals = append(totals, c)
	}

	for _, t := range totals {
		if total == 0 {
			total = t
		} else {
			total = total * t
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
