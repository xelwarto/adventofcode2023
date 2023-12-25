package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day23 int

var day23 = Day23(23)

type HikePath struct {
	X     int
	PX    int
	Y     int
	PY    int
	Count int
	Track [][]int
}

func (d Day23) Part1() (string, error) {
	s, err := util.File2Array("inputs/day23_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]string{}
	start := []int{0, 0}
	for y, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			grid = append(grid, data)

			for n := range data {
				if data[n] == "S" {
					start = []int{y, n}
				}
			}
		}
	}

	h1 := HikePath{
		Y:     start[0],
		PY:    0,
		X:     start[1],
		PX:    0,
		Count: 0,
	}
	hikes := []HikePath{h1}
	fhikes := []HikePath{}

	for len(hikes) > 0 {
		newhikes := []HikePath{}
		for n := range hikes {
			hike := hikes[n]

			hikey := hike.Y
			hikepy := hike.PY
			hikex := hike.X
			hikepx := hike.PX
			hikec := hike.Count

			if hikey+1 == len(grid) {
				fhikes = append(fhikes, hike)
			} else {

				if hikey+1 != hikepy {
					if grid[hikey+1][hikex] == "." || grid[hikey+1][hikex] == "v" {
						p := HikePath{
							X:     hikex,
							Y:     hikey + 1,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
						}
						newhikes = append(newhikes, p)
					}
				}

				if hikex+1 != hikepx {
					if grid[hikey][hikex+1] == "." || grid[hikey][hikex+1] == ">" {
						p := HikePath{
							X:     hikex + 1,
							Y:     hikey,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
						}
						newhikes = append(newhikes, p)
					}
				}

				if hikex-1 != hikepx {
					if grid[hikey][hikex-1] == "." {
						p := HikePath{
							X:     hikex - 1,
							Y:     hikey,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
						}
						newhikes = append(newhikes, p)
					}
				}

				if grid[hikey][hikex] != "S" {
					if hikey-1 != hikepy {
						if grid[hikey-1][hikex] == "." {
							p := HikePath{
								X:     hikex,
								Y:     hikey - 1,
								PX:    hikex,
								PY:    hikey,
								Count: hikec + 1,
							}
							newhikes = append(newhikes, p)
						}
					}
				}
			}
		}
		hikes = newhikes
	}

	// for _, line := range grid {
	// 	fmt.Println(line)
	// }

	for _, f := range fhikes {
		if total == 0 {
			total = f.Count
		} else {
			if f.Count > total {
				total = f.Count
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##################################################################################

func (d Day23) Part2() (string, error) {
	s, err := util.File2Array("inputs/day23_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	grid := [][]string{}
	start := []int{0, 0}
	for y, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			grid = append(grid, data)

			for n := range data {
				if data[n] == "S" {
					start = []int{y, n}
				}
			}
		}
	}

	h1 := HikePath{
		Y:     start[0],
		PY:    0,
		X:     start[1],
		PX:    0,
		Count: 0,
		Track: [][]int{},
	}
	// h1.Track = append(h1.Track, []int{start[0], start[1]})
	hikes := []HikePath{h1}
	fhikes := []HikePath{}

	for len(hikes) > 0 {
		// for q := 0; q < 200; q++ {
		// fmt.Println(len(hikes))
		newhikes := []HikePath{}
		for n := range hikes {
			hike := hikes[n]
			if hike.Y+1 == len(grid) {
				fhikes = append(fhikes, hike)
			} else {
				// if n >= 100 {
				// 	newhikes = append(newhikes, hike)
				// } else {
				hikey := hike.Y
				hikepy := hike.PY
				hikex := hike.X
				hikepx := hike.PX
				hikec := hike.Count
				hiket := hike.Track

				d := false
				for m := range hike.Track {
					t := hike.Track[m]
					if hikey == t[0] && hikex == t[1] {
						d = true
					}
				}
				if d {
					continue
				}

				hiket = append(hiket, []int{hikey, hikex})

				if hikey+1 != hikepy {
					if grid[hikey+1][hikex] != "#" {
						p := HikePath{
							X:     hikex,
							Y:     hikey + 1,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
							Track: hiket,
						}
						newhikes = append(newhikes, p)
					}
				}

				if hikex+1 != hikepx {
					if grid[hikey][hikex+1] != "#" {
						p := HikePath{
							X:     hikex + 1,
							Y:     hikey,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
							Track: hiket,
						}
						newhikes = append(newhikes, p)
					}
				}

				if hikex-1 != hikepx {
					if grid[hikey][hikex-1] != "#" {
						p := HikePath{
							X:     hikex - 1,
							Y:     hikey,
							PX:    hikex,
							PY:    hikey,
							Count: hikec + 1,
							Track: hiket,
						}
						newhikes = append(newhikes, p)
					}
				}

				if grid[hikey][hikex] != "S" {
					if hikey-1 != hikepy {
						if grid[hikey-1][hikex] != "#" {
							p := HikePath{
								X:     hikex,
								Y:     hikey - 1,
								PX:    hikex,
								PY:    hikey,
								Count: hikec + 1,
								Track: hiket,
							}
							newhikes = append(newhikes, p)
						}
					}
					// }
				}
			}
		}
		hikes = newhikes

		for _, f := range fhikes {
			if total == 0 {
				total = f.Count
			} else {
				if f.Count > total {
					total = f.Count
				}
			}
		}

		fmt.Println(total)

		fmt.Println(hikes)
	}

	// for _, line := range grid {
	// 	fmt.Println(line)
	// }

	// for _, f := range fhikes {
	// 	if total == 0 {
	// 		total = f.Count
	// 	} else {
	// 		if f.Count > total {
	// 			total = f.Count
	// 		}
	// 	}
	// }

	return fmt.Sprintf("%v", total), nil
}

func (d Day23) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day23)
}
