package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day18 int

var day18 = Day18(18)

type LagoonStep struct {
	Direction string
	Length    int
	Color     string
}

func (d Day18) Part1() (string, error) {
	s, err := util.File2Array("inputs/day18_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	steps := []LagoonStep{}
	for _, x := range s {
		if x != "" {
			data := strings.Split(x, " ")

			if len(data) == 3 {
				i, _ := strconv.Atoi(data[1])
				line := LagoonStep{
					Direction: data[0],
					Length:    i,
					Color:     data[2],
				}
				steps = append(steps, line)
			}
		}
	}

	xMax := 0
	xMin := 0
	yMax := 0
	yMin := 0

	xc := 1
	yc := 1
	for _, s := range steps {
		if s.Direction == "R" {
			xc = xc + s.Length
		}
		if s.Direction == "L" {
			xc = xc - s.Length
		}
		if s.Direction == "D" {
			yc = yc + s.Length
		}
		if s.Direction == "U" {
			yc = yc - s.Length
		}

		if xc > 0 && xc > xMax {
			xMax = xc
		}
		if yc > 0 && yc > yMax {
			yMax = yc
		}
		if xc < 0 && xc < xMin {
			xMin = xc
		}
		if yc < 0 && yc < yMin {
			yMin = yc
		}
	}

	gy := 1 - (yMin)
	gx := 1 - (xMin)

	yMax = (yMax - (yMin)) + 2
	xMax = (xMax - (xMin)) + 2

	grid := [][]string{}

	for q := 0; q < yMax; q++ {
		s := []string{}
		for w := 0; w < xMax; w++ {
			s = append(s, ".")
		}
		grid = append(grid, s)
	}

	grid[gy][gx] = "#"
	for _, s := range steps {
		if s.Direction == "R" {
			for q := 1; q <= s.Length; q++ {
				gx++
				grid[gy][gx] = "#"
			}
		}
		if s.Direction == "L" {
			for q := 1; q <= s.Length; q++ {
				gx--
				grid[gy][gx] = "#"
			}
		}
		if s.Direction == "D" {
			for q := 1; q <= s.Length; q++ {
				gy++
				grid[gy][gx] = "#"
			}
		}
		if s.Direction == "U" {
			for q := 1; q <= s.Length; q++ {
				gy--
				grid[gy][gx] = "#"
			}
		}
	}

	for y, line := range grid {
		for q := 0; q < len(line); q++ {
			if line[q] == "." {
				sy := y
				sx := q
				c := 0

				for {
					if (sy+1) >= len(grid) || (sx+1) >= len(line) {
						if c%2 != 0 {
							grid[y][q] = "*"
						}
						break
					}
					sy = (sy + 1)
					sx = (sx + 1)
					// 7L
					if grid[sy][sx] == "#" {
						if grid[sy+1][sx] != "#" || grid[sy][sx-1] != "#" {
							if grid[sy-1][sx] != "#" || grid[sy][sx+1] != "#" {
								c++
							}
						}
					}
				}
			}
		}
	}

	// for y, line := range tracker {
	// 	for q := 0; q < len(line); q++ {
	// 		if line[q] == "*" {
	// 			sy := y
	// 			sx := q
	// 			c := 0
	// 			for {
	// 				if (sy+1) >= len(tracker) || (sx+1) >= len(line) {
	// 					if c%2 != 0 {
	// 						total++
	// 					}
	// 					break
	// 				}
	// 				sy = (sy + 1)
	// 				sx = (sx + 1)
	// 				// -|FJ
	// 				if tracker[sy][sx] == "-" || tracker[sy][sx] == "|" || tracker[sy][sx] == "F" || tracker[sy][sx] == "J" {
	// 					c++
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	for _, g := range grid {
		for q := 0; q < len(g); q++ {
			if g[q] == "#" || g[q] == "*" {
				total++
			}
		}
		fmt.Println(g)
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day18) Part2() (string, error) {
	s, err := util.File2Array("inputs/day18_part2.txt")
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

func (d Day18) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day18)
}
