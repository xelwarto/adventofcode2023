package days

import (
	"code/util"
	"fmt"
	"log"
	"strings"
)

type Day10 int

var day10 = Day10(10)

type PipeMove struct {
	East  int
	West  int
	North int
	South int
}

var types = make(map[string]PipeMove)

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position

func init() {
	types["|"] = PipeMove{
		East:  0,
		West:  0,
		North: 1,
		South: 1,
	}
	types["-"] = PipeMove{
		East:  1,
		West:  1,
		North: 0,
		South: 0,
	}
	types["L"] = PipeMove{
		East:  1,
		West:  0,
		North: 1,
		South: 0,
	}
	types["J"] = PipeMove{
		East:  0,
		West:  1,
		North: 1,
		South: 0,
	}
	types["7"] = PipeMove{
		East:  0,
		West:  1,
		North: 0,
		South: 1,
	}
	types["F"] = PipeMove{
		East:  1,
		West:  0,
		North: 0,
		South: 1,
	}
}

func (d Day10) Part1() (string, error) {
	s, err := util.File2Array("inputs/day10_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	pipes := [][]string{}
	start := []int{}
	for y, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			pipes = append(pipes, data)

			s := strings.Index(x, "S")
			if s >= 0 {
				start = append(start, y)
				start = append(start, s)
			}
		}
	}

	py := start[0]
	px := start[1]

	sy := py
	sx := px + 1
	c := 0
	for {
		c++
		pipe := pipes[sy][sx]
		// fmt.Println(pipe)

		if pipe == "." {
			log.Fatal("GROUND")
		}

		if pipe == "S" {
			break
		}

		dir := types[pipe]
		hy := sy
		hx := sx

		if pipe == "|" || pipe == "-" {
			if (sy - dir.North) != py {
				sy = sy - dir.North
			} else if (sy + dir.South) != py {
				sy = sy + dir.South
			}

			if (sx + dir.East) != px {
				sx = sx + dir.East
			} else if (sx - dir.West) != px {
				sx = sx - dir.West
			}
		} else {
			if (sy - dir.North) != py {
				sy = sy - dir.North
			}

			if (sy + dir.South) != py {
				sy = sy + dir.South
			}

			if (sx + dir.East) != px {
				sx = sx + dir.East
			}

			if (sx - dir.West) != px {
				sx = sx - dir.West
			}
		}

		py = hy
		px = hx
	}

	total = c / 2

	return fmt.Sprintf("%v", total), nil
}

// ###################################################################
func (d Day10) Part2() (string, error) {
	s, err := util.File2Array("inputs/day10_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	pipes := [][]string{}
	tracker := [][]string{}
	start := []int{}
	for y, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			pipes = append(pipes, data)

			t := []string{}
			for q := 0; q < len(data); q++ {
				t = append(t, "*")
			}
			tracker = append(tracker, t)

			s := strings.Index(x, "S")
			if s >= 0 {
				start = append(start, y)
				start = append(start, s)
			}
		}
	}

	py := start[0]
	px := start[1]

	sy := py
	sx := px + 1
	for {
		pipe := pipes[sy][sx]
		tracker[sy][sx] = pipe
		// fmt.Println(pipe)

		if pipe == "." {
			log.Fatal("GROUND")
		}

		if pipe == "S" {
			break
		}

		dir := types[pipe]
		hy := sy
		hx := sx

		if pipe == "|" || pipe == "-" {
			if (sy - dir.North) != py {
				sy = sy - dir.North
			} else if (sy + dir.South) != py {
				sy = sy + dir.South
			}

			if (sx + dir.East) != px {
				sx = sx + dir.East
			} else if (sx - dir.West) != px {
				sx = sx - dir.West
			}
		} else {
			if (sy - dir.North) != py {
				sy = sy - dir.North
			}

			if (sy + dir.South) != py {
				sy = sy + dir.South
			}

			if (sx + dir.East) != px {
				sx = sx + dir.East
			}

			if (sx - dir.West) != px {
				sx = sx - dir.West
			}
		}

		py = hy
		px = hx
	}

	tracker2 := [][]string{}
	for _, line := range tracker {
		if len(tracker2) == 0 {
			for q := 0; q < len(line); q++ {
				t := []string{line[q]}
				tracker2 = append(tracker2, t)
			}
		} else {
			for q := 0; q < len(line); q++ {
				t := tracker2[q]
				t = append(t, line[q])
				tracker2[q] = t
			}
		}
	}

	for _, line := range tracker2 {
		fmt.Println(line)

	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day10) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day10)
}

// [\|J7]\*+[\|FL]
