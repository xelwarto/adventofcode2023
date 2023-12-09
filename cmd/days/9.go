package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day9 int

var day9 = Day9(9)

func (d Day9) Part1() (string, error) {
	s, err := util.File2Array("inputs/day9_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		data := [][]int{}
		if x != "" {
			d := strings.Split(x, " ")
			n := []int{}
			for _, d1 := range d {
				i, _ := strconv.Atoi(d1)
				n = append(n, i)
			}
			data = append(data, n)

			for {
				n1 := data[len(data)-1]
				n2 := []int{}

				for q := 0; q < len(n1)-1; q++ {
					n2 = append(n2, n1[q+1]-n1[q])
				}
				data = append(data, n2)

				v := 0
				for _, i := range n2 {
					v = v + i
				}
				if v == 0 {
					break
				}
			}

			vals := []int{}
			for q := len(data) - 1; q >= 0; q-- {
				d := data[q]
				vals = append(vals, d[len(d)-1])
			}

			v := 0
			for _, i := range vals {
				v = v + i
			}
			total = total + v
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day9) Part2() (string, error) {
	s, err := util.File2Array("inputs/day9_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		data := [][]int{}
		if x != "" {
			d := strings.Split(x, " ")
			n := []int{}
			for _, d1 := range d {
				i, _ := strconv.Atoi(d1)
				n = append(n, i)
			}
			data = append(data, n)

			for {
				n1 := data[len(data)-1]
				n2 := []int{}

				for q := 0; q < len(n1)-1; q++ {
					n2 = append(n2, n1[q+1]-n1[q])
				}
				data = append(data, n2)

				v := 0
				for _, i := range n2 {
					v = v + i
				}
				if v == 0 {
					break
				}
			}

			vals := []int{}
			for q := len(data) - 1; q >= 0; q-- {
				d := data[q]
				vals = append(vals, d[0])
			}

			v := 0
			for _, i := range vals {
				v = i - v
			}
			total = total + v
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day9) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day9)
}
