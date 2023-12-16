package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day16 int

var day16 = Day16(16)

type Beam struct {
	X  int
	Y  int
	DY int
	DX int
	On bool
}

func findTotal(b Beam, grid [][]string) int {
	c2 := 0
	total := 0

	eng := [][]int{}

	for _, g := range grid {
		i := []int{}
		for q := 0; q < len(g); q++ {
			i = append(i, 0)
		}
		eng = append(eng, i)
	}

	tracker := []Beam{}
	tracker = append(tracker, b)

	for w := 0; w < 10000; w++ {
		end := true
		for z := range tracker {
			beam := tracker[z]
			if beam.On {
				beam.X = beam.X + (beam.DX)
				beam.Y = beam.Y + (beam.DY)

				if beam.X < len(grid[0]) && beam.Y < len(grid) {
					if beam.X >= 0 && beam.Y >= 0 {
						end = false
						g := grid[beam.Y][beam.X]
						eng[beam.Y][beam.X] = 1

						if g == "|" && beam.DX != 0 {
							beam.DX = 0
							beam.DY = 1

							b := Beam{
								Y:  beam.Y,
								X:  beam.X,
								DX: 0,
								DY: -1,
								On: true,
							}
							tracker = append(tracker, b)
						}

						if g == "-" && beam.DY != 0 {
							beam.DX = 1
							beam.DY = 0

							b := Beam{
								Y:  beam.Y,
								X:  beam.X,
								DX: -1,
								DY: 0,
								On: true,
							}
							tracker = append(tracker, b)
						}

						if g == "/" {
							f := false
							if beam.DX > 0 && !f {
								beam.DX = 0
								beam.DY = -1
								f = true
							}

							if beam.DX < 0 && !f {
								beam.DX = 0
								beam.DY = 1
								f = true
							}

							if beam.DY > 0 && !f {
								beam.DX = -1
								beam.DY = 0
								f = true
							}

							if beam.DY < 0 && !f {
								beam.DX = 1
								beam.DY = 0
								f = true
							}
						}

						if g == "\\" {
							f := false
							if beam.DX > 0 && !f {
								beam.DX = 0
								beam.DY = 1
								f = true
							}

							if beam.DX < 0 && !f {
								beam.DX = 0
								beam.DY = -1
								f = true
							}

							if beam.DY > 0 && !f {
								beam.DX = 1
								beam.DY = 0
								f = true
							}

							if beam.DY < 0 && !f {
								beam.DX = -1
								beam.DY = 0
								f = true
							}
						}
					} else {
						beam.On = false
					}
				} else {
					beam.On = false
				}
				tracker[z] = beam
			}
		}

		if end {
			break
		}

		c := 0
		for _, e := range eng {
			for q := 0; q < len(e); q++ {
				if e[q] == 1 {
					c++
				}
			}
		}
		if c > total {
			total = c
			c2 = 0
		} else {
			c2++
		}

		if c2 > 10 {
			break
		}
	}
	return total
}

func (d Day16) Part1() (string, error) {
	s, err := util.File2Array("inputs/day16_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]string{}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			grid = append(grid, data)
		}
	}

	b := Beam{
		Y:  0,
		X:  -1,
		DX: 1,
		DY: 0,
		On: true,
	}
	total = findTotal(b, grid)

	return fmt.Sprintf("%v", total), nil
}

// #################################################################################################

func (d Day16) Part2() (string, error) {
	s, err := util.File2Array("inputs/day16_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]string{}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			grid = append(grid, data)
		}
	}

	for q := 0; q < len(grid[0]); q++ {
		b := Beam{
			Y:  -1,
			X:  q,
			DX: 0,
			DY: 1,
			On: true,
		}
		t := findTotal(b, grid)

		if t > total {
			total = t
		}
	}

	for q := 0; q < len(grid[0]); q++ {
		b := Beam{
			Y:  len(grid),
			X:  q,
			DX: 0,
			DY: -1,
			On: true,
		}
		t := findTotal(b, grid)

		if t > total {
			total = t
		}
	}

	for q := 0; q < len(grid); q++ {
		b := Beam{
			Y:  q,
			X:  -1,
			DX: 1,
			DY: 0,
			On: true,
		}
		t := findTotal(b, grid)

		if t > total {
			total = t
		}
	}

	for q := 0; q < len(grid); q++ {
		b := Beam{
			Y:  q,
			X:  len(grid[0]),
			DX: -1,
			DY: 0,
			On: true,
		}
		t := findTotal(b, grid)

		if t > total {
			total = t
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day16) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day16)
}
