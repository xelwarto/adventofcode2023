package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day11 int

var day11 = Day11(11)

func (d Day11) Part1() (string, error) {
	s, err := util.File2Array("inputs/day11_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	space1 := []string{}

	for _, x := range s {
		if x != "" {
			if !strings.Contains(x, "#") {
				space1 = append(space1, x)
			}
			space1 = append(space1, x)
		}
	}

	tracker := [][]string{}
	for _, line := range space1 {
		tracker = append(tracker, strings.Split(line, ""))
	}

	e := []int{}
	for q := 0; q < len(tracker[0]); q++ {
		f := false
		for x := 0; x < len(tracker); x++ {
			if tracker[x][q] == "#" {
				f = true
			}
		}
		if !f {
			e = append(e, q)
		}
	}

	for y, line := range tracker {
		nline := []string{}
		for q := 0; q < len(line); q++ {
			nline = append(nline, line[q])
			for x := 0; x < len(e); x++ {
				if e[x] == q {
					nline = append(nline, ".")
				}
			}
		}
		tracker[y] = nline
	}

	g := [][]int{}
	for y, line := range tracker {
		for q := 0; q < len(line); q++ {
			if line[q] == "#" {
				g = append(g, []int{y, q})
			}
		}
	}

	c := 0
	for q := 0; q < len(g); q++ {
		gax := g[q]
		for x := q + 1; x < len(g); x++ {
			gax2 := g[x]

			x := 0
			y := 0

			if gax[1] <= gax2[1] {
				y = gax2[0] - gax[0]
				x = (gax2[1] - gax[1])
			} else {
				y = gax2[0] - gax[0]
				x = (gax[1] - gax2[1])
			}
			c++
			// fmt.Printf("(%v) %v: y: %v x:%v\n", c, q, y, x)
			total = total + y + x
		}
	}

	// for _, line := range tracker {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%v", total), nil
}

// #####################################################################

func (d Day11) Part2() (string, error) {
	s, err := util.File2Array("inputs/day11_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	space1 := []string{}
	mul := 1000000

	ey := []int{}
	for y, x := range s {
		if x != "" {
			if !strings.Contains(x, "#") {
				ey = append(ey, y)
			}
			space1 = append(space1, x)
		}
	}

	tracker := [][]string{}
	for _, line := range space1 {
		tracker = append(tracker, strings.Split(line, ""))
	}

	ex := []int{}
	for q := 0; q < len(tracker[0]); q++ {
		f := false
		for x := 0; x < len(tracker); x++ {
			if tracker[x][q] == "#" {
				f = true
			}
		}
		if !f {
			ex = append(ex, q)
		}
	}

	// for y, line := range tracker {
	// 	nline := []string{}
	// 	for q := 0; q < len(line); q++ {
	// 		nline = append(nline, line[q])
	// 		for x := 0; x < len(ex); x++ {
	// 			if ex[x] == q {
	// 				for z := 0; z < mul-1; z++ {
	// 					nline = append(nline, ".")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	tracker[y] = nline
	// }

	g := [][]int{}
	for y, line := range tracker {
		for q := 0; q < len(line); q++ {
			if line[q] == "#" {
				g = append(g, []int{y, q})
			}
		}
	}

	c := 0
	for q := 0; q < len(g); q++ {
		gax := g[q]
		for x := q + 1; x < len(g); x++ {
			gax2 := g[x]

			x := 0
			y := 0

			if gax[1] <= gax2[1] {
				cy := 0
				for z := 0; z < len(ey); z++ {
					if gax2[0] > ey[z] && gax[0] < ey[z] {
						cy++
					}
				}
				y = (gax2[0] + (cy * (mul - 1))) - gax[0]

				cx := 0
				for z := 0; z < len(ex); z++ {
					if gax[1] < ex[z] && gax2[1] > ex[z] {
						cx++
					}
				}
				x = (gax2[1] + (cx * (mul - 1))) - gax[1]
			} else {
				cy := 0
				for z := 0; z < len(ey); z++ {
					if gax2[0] > ey[z] && gax[0] < ey[z] {
						cy++
					}
				}
				y = (gax2[0] + (cy * (mul - 1))) - gax[0]

				cx := 0
				for z := 0; z < len(ex); z++ {
					if gax2[1] < ex[z] && gax[1] > ex[z] {
						cx++
					}
				}
				x = (gax[1] + (cx * (mul - 1))) - gax2[1]
			}
			c++
			// fmt.Printf("(%v) %v: y: %v x:%v\n", c, q, y, x)
			total = total + y + x
		}
	}

	// for _, line := range tracker {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%v", total), nil
}

func (d Day11) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day11)
}
