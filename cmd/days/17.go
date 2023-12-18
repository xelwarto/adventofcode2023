package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day17 int

var day17 = Day17(17)

type Path struct {
	X        int
	Y        int
	DY       int
	DX       int
	Forward  int
	Finished bool
	Tracker  [][]int
}

func (d Day17) Part1() (string, error) {
	s, err := util.File2Array("inputs/day17_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]int{}
	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			line := []int{}
			for _, d := range data {
				i, _ := strconv.Atoi(d)
				line = append(line, i)
			}
			grid = append(grid, line)
		}
	}

	start := Path{
		X:        0,
		Y:        0,
		DY:       0,
		DX:       1,
		Forward:  0,
		Finished: false,
		Tracker:  [][]int{},
	}
	tracker := []Path{start}

	for q := 0; q < 20; q++ {
		end := true
		fmt.Println(len(tracker))

		t2 := []Path{}
		for x := range tracker {
			t := tracker[x]
			if !t.Finished {
				end = false

				if t.DX != 0 {
					if (t.Y - 1) >= 0 {
						new := Path{
							X:        t.X,
							Y:        t.Y,
							DX:       0,
							DY:       -1,
							Forward:  0,
							Finished: false,
						}
						f := true
						for _, m := range t.Tracker {
							if new.X == m[1] && new.Y == m[0] {
								f = false
							}
							new.Tracker = append(new.Tracker, m)
						}

						if f {
							t2 = append(t2, new)
						}
					}
					if (t.Y + 1) < len(grid) {
						new := Path{
							X:        t.X,
							Y:        t.Y,
							DX:       0,
							DY:       1,
							Forward:  0,
							Finished: false,
						}
						f := true
						for _, m := range t.Tracker {
							if new.X == m[1] && new.Y == m[0] {
								f = false
							}
							new.Tracker = append(new.Tracker, m)
						}

						if f {
							t2 = append(t2, new)
						}
					}
				}

				if t.DY != 0 {
					if (t.X - 1) >= 0 {
						new := Path{
							X:        t.X,
							Y:        t.Y,
							DX:       -1,
							DY:       0,
							Forward:  0,
							Finished: false,
						}
						f := true
						for _, m := range t.Tracker {
							if new.X == m[1] && new.Y == m[0] {
								f = false
							}
							new.Tracker = append(new.Tracker, m)
						}

						if f {
							t2 = append(t2, new)
						}
					}
					if (t.X + 1) < len(grid[0]) {
						new := Path{
							X:        t.X,
							Y:        t.Y,
							DX:       1,
							DY:       0,
							Forward:  0,
							Finished: false,
						}
						f := true
						for _, m := range t.Tracker {
							if new.X == m[1] && new.Y == m[0] {
								f = false
							}
							new.Tracker = append(new.Tracker, m)
						}

						if f {
							t2 = append(t2, new)
						}
					}
				}

				t.Tracker = append(t.Tracker, []int{t.Y, t.X})

				if (t.X+(t.DX)) < len(grid[0]) && (t.Y+(t.DY)) < len(grid) &&
					(t.X+(t.DX)) >= 0 && (t.Y+(t.DY)) >= 0 {
					t.X = t.X + (t.DX)
					t.Y = t.Y + (t.DY)
					t.Forward++

					if t.Forward > 3 {
						t.Finished = true
					} else {
						t2 = append(t2, t)
					}
				} else {
					t.Finished = true
				}
			}
		}
		tracker = t2
		if end {
			break
		}
	}

	for _, t := range tracker {
		fmt.Println(t)
	}

	// for _, line := range grid {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%v", total), nil
}

func (d Day17) Part2() (string, error) {
	s, err := util.File2Array("inputs/day17_part2.txt")
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

func (d Day17) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day17)
}
