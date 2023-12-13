package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day12 int

var day12 = Day12(12)

func sData(s string) string {
	res := []string{}
	spring := 0
	data := strings.Split(s, "")

	for _, d := range data {
		if d == "#" {
			spring++
		} else {
			if spring > 0 {
				sp := strconv.Itoa(spring)
				res = append(res, sp)
				spring = 0
			}
		}
	}
	if spring > 0 {
		sp := strconv.Itoa(spring)
		res = append(res, sp)
	}
	return strings.Join(res, ",")
}

func perms(s []string, b []int, i int) []string {
	// fmt.Println(i)
	res := []string{}
	if i < len(b) {
		p := b[i]
		s[p] = "#"
		res = append(res, strings.Join(s, ""))

		for q := i + 1; q < len(b); q++ {
			for _, r := range perms(s, b, q) {
				res = append(res, r)
			}
		}

		s[p] = "."
		res = append(res, strings.Join(s, ""))

		for q := i + 1; q < len(b); q++ {
			for _, r := range perms(s, b, q) {
				res = append(res, r)
			}
		}
	}
	return res
}

func (d Day12) Part1() (string, error) {
	s, err := util.File2Array("inputs/day12_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for y, x := range s {
		fmt.Println(y + 1)
		if x != "" {
			line := strings.Split(x, " ")
			val := line[1]
			data := strings.Split(line[0], "")

			b := []int{}
			for y, d := range data {
				if d == "?" {
					b = append(b, y)
				}
			}

			for _, i := range b {
				data[i] = "."
			}
			res1 := perms(data, b, 0)
			fmt.Println(len(res1))

			c := 0
			match := []string{}
			for _, r := range res1 {
				// fmt.Println(r)
				if sData(r) == val {
					f := false
					for _, m := range match {
						if r == m {
							f = true
						}
					}
					if !f {
						match = append(match, r)
						c++
					}

				}
			}
			total = total + c
		}

	}

	return fmt.Sprintf("%v", total), nil
}

// ##########################################################################################

func perms2(s []string, b []int, i int) []string {

	res := []string{}
	res = append(res, strings.Join(s, ""))
	s[b[i]] = "#"

	if i+1 < len(b) {
		for q := i + 1; q < len(b); q++ {
			for _, r := range perms2(s, b, q) {
				res = append(res, r)
			}
		}
	}
	fmt.Println(i)
	fmt.Println(res)

	// s[b[i]] = "."
	// if i+1 < len(b) {
	// 	for q := i + 1; q < len(b); q++ {
	// 		for _, r := range perms2(s, b, q) {
	// 			res = append(res, r)
	// 		}
	// 	}
	// }

	// s[b[i]] = "."
	// res = append(res, strings.Join(s, ""))

	// s[b[i]] = "#"
	// if i < len(b) {
	// 	res = perms(s, b, i+1)

	// 	// s[b[i]] = "."
	// 	// res = append(res, strings.Join(s, ""))

	// 	// for q := i + 1; q < len(b); q++ {
	// 	// 	for _, r := range perms(s, b, q) {
	// 	// 		res = append(res, r)
	// 	// 	}
	// 	// }
	// }

	// res = append(res, strings.Join(s, ""))
	return res
}

// ....###
// #...###
// ##..###
// ###.###
// #.#.###
// .#..###
// .##.###
// ..#.###

// #...###
// ##..###
// ###.###
// ##..###
// #...###
// #.#.###
// #...###
// #.#.###
// #...###

func (d Day12) Part2() (string, error) {
	s, err := util.File2Array("inputs/day12_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		// fmt.Println(y + 1)
		if x != "" {
			line := strings.Split(x, " ")
			val := line[1]
			data := strings.Split(line[0], "")

			b := []int{}
			for y, d := range data {
				if d == "?" {
					b = append(b, y)
				}
			}

			for _, i := range b {
				data[i] = "."
			}
			res1 := perms2(data, b, 0)
			// fmt.Println(len(res1))

			c := 0
			for _, r := range res1 {
				// fmt.Println(r)
				if sData(r) == val {
					c++
				}
			}
			total = total + c
		}

	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day12) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day12)
}
