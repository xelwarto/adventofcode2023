package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day2 int

var day2 = Day2(2)

func (d Day2) Part1() (string, error) {
	s, err := util.File2Array("inputs/day2_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	for _, x := range s {
		if x != "" {
			f := true
			data := strings.Split(x, ": ")
			gData := strings.Split(data[0], " ")

			id, err := strconv.Atoi(gData[1])
			if err != nil {
				return "", err
			}

			dData := strings.Split(data[1], ", ")
			for _, d := range dData {
				r := strings.Split(d, " ")
				i, err := strconv.Atoi(r[0])
				if err != nil {
					return "", err
				}

				if r[1] == "red" {
					if i > 12 {
						f = false
					}
				}

				if r[1] == "blue" {
					if i > 14 {
						f = false
					}
				}

				if r[1] == "green" {
					if i > 13 {
						f = false
					}
				}
			}
			if f {
				total = total + id
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day2) Part2() (string, error) {
	s, err := util.File2Array("inputs/day2_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	for _, x := range s {
		if x != "" {
			red := 0
			green := 0
			blue := 0

			data := strings.Split(x, ": ")
			dData := strings.Split(data[1], ", ")
			for _, d := range dData {
				r := strings.Split(d, " ")
				i, err := strconv.Atoi(r[0])
				if err != nil {
					return "", err
				}

				if r[1] == "red" {
					if i > red {
						red = i
					}
				}

				if r[1] == "blue" {
					if i > blue {
						blue = i
					}
				}

				if r[1] == "green" {
					if i > green {
						green = i
					}
				}
			}

			total = total + (red * green * blue)

		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day2) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day2)
}
