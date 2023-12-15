package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day14 int

var day14 = Day14(14)

func (d Day14) Part1() (string, error) {
	s, err := util.File2Array("inputs/day14_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	tracker := [][]string{}
	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			tracker = append(tracker, data)
		}
	}

	for y := 1; y < len(tracker); y++ {
		for q := 0; q < len(tracker[y]); q++ {
			cur := tracker[y][q]
			if cur == "O" {
				for z := y - 1; z >= 0; z-- {
					pre := tracker[z][q]
					if pre == "O" || pre == "#" {
						break
					} else {
						tracker[z][q] = "O"
						if (z + 1) < len(tracker) {
							tracker[z+1][q] = "."
						}
					}
				}
			}
		}
	}

	for y := range tracker {
		// fmt.Println(t)
		for q := 0; q < len(tracker[y]); q++ {
			if tracker[y][q] == "O" {
				total = total + (len(tracker) - y)
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ################################################################################################
func (d Day14) Part2() (string, error) {
	s, err := util.File2Array("inputs/day14_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	tracker := [][]string{}
	c := 1000
	for _, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			tracker = append(tracker, data)
		}
	}

	for w := 1; w <= c; w++ {

		//  north
		for y := 1; y < len(tracker); y++ {
			for q := 0; q < len(tracker[y]); q++ {
				cur := tracker[y][q]
				if cur == "O" {
					for z := y - 1; z >= 0; z-- {
						pre := tracker[z][q]
						if pre == "O" || pre == "#" {
							break
						} else {
							tracker[z][q] = "O"
							if (z + 1) < len(tracker) {
								tracker[z+1][q] = "."
							}
						}
					}
				}
			}
		}

		// fmt.Println("North")
		// for _, t := range tracker {
		// 	fmt.Println(t)
		// }
		// fmt.Println()

		// west
		for q := 1; q < len(tracker[0]); q++ {
			for y := 0; y < len(tracker); y++ {
				cur := tracker[y][q]
				if cur == "O" {
					for z := q - 1; z >= 0; z-- {
						pre := tracker[y][z]
						if pre == "O" || pre == "#" {
							break
						} else {
							tracker[y][z] = "O"
							if (z + 1) < len(tracker[0]) {
								tracker[y][z+1] = "."
							}
						}
					}
				}
			}
		}

		// fmt.Println("West")
		// for _, t := range tracker {
		// 	fmt.Println(t)
		// }
		// fmt.Println()

		//  south
		for y := len(tracker) - 1; y >= 0; y-- {
			for q := 0; q < len(tracker[y]); q++ {
				cur := tracker[y][q]
				if cur == "O" {
					for z := y + 1; z < len(tracker); z++ {
						pre := tracker[z][q]
						if pre == "O" || pre == "#" {
							break
						} else {
							tracker[z][q] = "O"
							if (z - 1) >= 0 {
								tracker[z-1][q] = "."
							}
						}
					}
				}
			}
		}

		// fmt.Println("South")
		// for _, t := range tracker {
		// 	fmt.Println(t)
		// }
		// fmt.Println()

		// east
		for q := len(tracker[0]) - 1; q >= 0; q-- {
			for y := 0; y < len(tracker); y++ {
				cur := tracker[y][q]
				if cur == "O" {
					for z := q + 1; z < len(tracker[0]); z++ {
						pre := tracker[y][z]
						if pre == "O" || pre == "#" {
							break
						} else {
							tracker[y][z] = "O"
							if (z - 1) >= 0 {
								tracker[y][z-1] = "."
							}
						}
					}
				}
			}
		}

		// fmt.Println("East")
		// for _, t := range tracker {
		// 	fmt.Println(t)
		// }
		// fmt.Println()
	}

	for y := range tracker {
		// fmt.Println(t)
		for q := 0; q < len(tracker[y]); q++ {
			if tracker[y][q] == "O" {
				total = total + (len(tracker) - y)
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day14) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day14)
}
