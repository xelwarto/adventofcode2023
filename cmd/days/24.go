package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

// https://www.mathopenref.com/coordintersection.html
// https://www.mathopenref.com/coordequationps.html (point - slope form)

// m = (y - Py) / (x - Px)
// x = ( m1(Px1) - Py1 - m2(Px2) + Py2 ) / (m1 - m2)
// y = m(x - Px) + Py

type Day24 int

var day24 = Day24(24)

type Storm struct {
	Px float64
	X  float64
	Py float64
	Y  float64
	Z  float64
	Vx float64
	Vy float64
	Vz float64
	M  float64
}

var min = float64(200000000000000)
var max = float64(400000000000000)

func (d Day24) Part1() (string, error) {
	s, err := util.File2Array("inputs/day24_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	storms := []Storm{}
	for _, x := range s {
		if x != "" {
			x = strings.Replace(x, " @", ",", -1)
			x = strings.ReplaceAll(x, " ", "")
			data := strings.Split(x, ",")

			dx, _ := strconv.ParseFloat(data[0], 64)
			dy, _ := strconv.ParseFloat(data[1], 64)
			dz, _ := strconv.ParseFloat(data[2], 64)
			dvx, _ := strconv.ParseFloat(data[3], 64)
			dvy, _ := strconv.ParseFloat(data[4], 64)
			dvz, _ := strconv.ParseFloat(data[5], 64)

			s := Storm{
				Px: dx,
				X:  dx + (dvx),
				Py: dy,
				Y:  dy + (dvy),
				Z:  dz,
				Vx: dvx,
				Vy: dvy,
				Vz: dvz,
			}

			storms = append(storms, s)
		}
	}

	for q := 0; q < len(storms); q++ {
		if (q + 1) < len(storms) {
			for w := q + 1; w < len(storms); w++ {
				s1 := storms[q]
				s2 := storms[w]

				s1.M = float64((s1.Y - s1.Py)) / float64((s1.X - s1.Px))
				s2.M = float64((s2.Y - s2.Py)) / float64((s2.X - s2.Px))

				x := float64((float64((s1.M * s1.Px)) - s1.Py - float64((s2.M * s2.Px)) + s2.Py)) / float64((s1.M - (s2.M)))
				y := (float64(s1.M)*float64((x-s1.Px)) + float64(s1.Py))

				// fmt.Println(s1)
				// fmt.Println(s2)

				// fmt.Println(x)
				// fmt.Println(y)
				// fmt.Println()

				f := true

				if s1.Vx < 0 {
					if x > s1.Px {
						f = false
					}
				} else {
					if x < s1.Px {
						f = false
					}
				}

				if s2.Vx < 0 {
					if x > s2.Px {
						f = false
					}
				} else {
					if x < s2.Px {
						f = false
					}
				}

				if s1.Vy < 0 {
					if y > s1.Py {
						f = false
					}
				} else {
					if y < s1.Py {
						f = false
					}
				}

				if s2.Vy < 0 {
					if y > s2.Py {
						f = false
					}
				} else {
					if y < s2.Py {
						f = false
					}
				}

				if (x < min || x > max) || (y < min || y > max) {
					f = false
				}

				if f {
					// fmt.Println(s1)
					// fmt.Println(s2)

					// fmt.Println(x)
					// fmt.Println(y)
					// fmt.Println()
					total++
				}
			}
		}

	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day24) Part2() (string, error) {
	s, err := util.File2Array("inputs/day24_part2.txt")
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

func (d Day24) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day24)
}
