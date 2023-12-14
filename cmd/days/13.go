package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day13 int

var day13 = Day13(13)

func (d Day13) Part1() (string, error) {
	s, err := util.File2Array("inputs/day13_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	tracker := [][]string{}
	lines := []string{}
	for _, x := range s {
		if x != "" {
			lines = append(lines, x)
		} else {
			tracker = append(tracker, lines)
			lines = []string{}
		}
	}

	if len(lines) > 0 {
		tracker = append(tracker, lines)
	}

	for y, t := range tracker {
		f := false

		rf := -1
		for y := 0; y < len(t); y++ {
			if y+1 < len(t) {
				if t[y] == t[y+1] {
					rf = y
					break
				}
			}
		}
		if rf != -1 {
			r2 := rf
			for q := rf; q >= 0; q-- {
				if (r2 + 1) < len(t) {
					r2++
					if t[q] == t[r2] {
						f = true
					} else {
						f = false
						break
					}
				}
			}
		}

		if rf != -1 && !f {
			sf := rf
			rf = -1
			for y := sf + 1; y < len(t); y++ {
				if y+1 < len(t) {
					if t[y] == t[y+1] {
						rf = y
						break
					}
				}
			}
			if rf != -1 {
				r2 := rf
				for q := rf; q >= 0; q-- {
					if (r2 + 1) < len(t) {
						r2++
						if t[q] == t[r2] {
							f = true
						} else {
							f = false
							break
						}
					}
				}
			}
		}

		if f {
			total = total + (100 * (rf + 1))
			// fmt.Printf("H RF Found (%v): %v\n", y, rf)
		}

		if !f {
			rf = -1
			m := [][]string{}
			for _, l := range t {
				data := strings.Split(l, "")
				m = append(m, data)
			}

			for q := 0; q < len(m[0]); q++ {
				d1 := []string{}
				d2 := []string{}
				for z := 0; z < len(m); z++ {
					if q+1 < len(m[0]) {
						d1 = append(d1, m[z][q])
						d2 = append(d2, m[z][q+1])
					}
				}
				if strings.Join(d1, "") == strings.Join(d2, "") {
					rf = q
					break
				}
			}

			if rf != -1 {
				r2 := rf
				for q := rf; q >= 0; q-- {
					if (r2 + 1) < len(m[0]) {
						r2++
						d1 := []string{}
						d2 := []string{}
						for z := 0; z < len(m); z++ {
							d1 = append(d1, m[z][q])
							d2 = append(d2, m[z][r2])
						}
						if strings.Join(d1, "") == strings.Join(d2, "") {
							f = true
						} else {
							f = false
							break
						}
					}
				}
			}

			if rf != -1 && !f {
				sf := rf
				rf = -1
				m := [][]string{}
				for _, l := range t {
					data := strings.Split(l, "")
					m = append(m, data)
				}

				for q := sf + 1; q < len(m[0]); q++ {
					d1 := []string{}
					d2 := []string{}
					for z := 0; z < len(m); z++ {
						if q+1 < len(m[0]) {
							d1 = append(d1, m[z][q])
							d2 = append(d2, m[z][q+1])
						}
					}
					if strings.Join(d1, "") == strings.Join(d2, "") {
						rf = q
						break
					}
				}

				if rf != -1 {
					r2 := rf
					for q := rf; q >= 0; q-- {
						if (r2 + 1) < len(m[0]) {
							r2++
							d1 := []string{}
							d2 := []string{}
							for z := 0; z < len(m); z++ {
								d1 = append(d1, m[z][q])
								d2 = append(d2, m[z][r2])
							}
							if strings.Join(d1, "") == strings.Join(d2, "") {
								f = true
							} else {
								f = false
								break
							}
						}
					}
				}
			}

			if f {
				total = total + (rf + 1)
				// fmt.Printf("V RF Found (%v): %v\n", y, rf)
				// for _, l := range t {
				// 	fmt.Println(l)
				// }
				// fmt.Println()
			}
		}

		if !f {
			fmt.Println(rf)
			fmt.Printf("RF NOT Found (%v)\n", y)
			for _, l := range t {
				fmt.Println(l)
			}
			fmt.Println()
		}

	}
	return fmt.Sprintf("%v", total), nil
}

// #################################################################################

func (d Day13) Part2() (string, error) {
	s, err := util.File2Array("inputs/day13_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	tracker := [][]string{}
	lines := []string{}
	for _, x := range s {
		if x != "" {
			lines = append(lines, x)
		} else {
			tracker = append(tracker, lines)
			lines = []string{}
		}
	}

	if len(lines) > 0 {
		tracker = append(tracker, lines)
	}

	for y, t := range tracker {
		f := false
		e := 0

		rf := -1
		for y := 0; y < len(t); y++ {
			if y+1 < len(t) {
				if t[y] == t[y+1] {
					rf = y
					break
				}
			}
		}
		if rf != -1 {
			r2 := rf
			for q := rf; q >= 0; q-- {
				if (r2 + 1) < len(t) {
					r2++
					if t[q] != t[r2] {
						e++
					}
				}
			}
		}

		if e == 1 {
			f = true
		}

		if rf != -1 && !f {
			sf := rf
			rf = -1
			e = 0
			for y := sf + 1; y < len(t); y++ {
				if y+1 < len(t) {
					if t[y] == t[y+1] {
						rf = y
						break
					}
				}
			}
			if rf != -1 {
				r2 := rf
				for q := rf; q >= 0; q-- {
					if (r2 + 1) < len(t) {
						r2++
						if t[q] != t[r2] {
							e++
						}
					}
				}
			}
			if e == 1 {
				f = true
			}
		}

		if !f {
			rf = -1
			m := [][]string{}
			for _, l := range t {
				data := strings.Split(l, "")
				m = append(m, data)
			}

			for q := 0; q < len(m); q++ {
				e = 0
				for z := 0; z < len(m[q]); z++ {
					if q+1 < len(m) {
						if m[q][z] != m[q+1][z] {
							e++
						}
					}
				}

				if e == 1 {
					rf = q
					break
				}
			}

			if e == 1 {
				if rf != -1 {
					r2 := rf
					for q := rf; q >= 0; q-- {
						if (r2 + 1) < len(t) {
							r2++

							if rf == (r2 - 1) {
								f = true
							} else {
								if t[q] == t[r2] {
									f = true
								} else {
									f = false
									break
								}
							}
						}
					}
				}
			}
		}

		if !f {
			sf := rf
			rf = -1
			m := [][]string{}
			for _, l := range t {
				data := strings.Split(l, "")
				m = append(m, data)
			}

			for q := sf + 1; q < len(m); q++ {
				e = 0
				for z := 0; z < len(m[q]); z++ {
					if q+1 < len(m) {
						if m[q][z] != m[q+1][z] {
							e++
						}
					}
				}

				if e == 1 {
					rf = q
					break
				}
			}

			if e == 1 {
				if rf != -1 {
					r2 := rf
					for q := rf; q >= 0; q-- {
						if (r2 + 1) < len(t) {
							r2++

							if rf == (r2 - 1) {
								f = true
							} else {
								if t[q] == t[r2] {
									f = true
								} else {
									f = false
									break
								}
							}
						}
					}
				}
			}
		}

		if f {
			total = total + (100 * (rf + 1))
			fmt.Printf("H RF Found (%v): %v\n", y+1, rf)
		}

		// ############################################################################################################

		if !f {
			rf = -1
			m := [][]string{}
			for _, l := range t {
				data := strings.Split(l, "")
				m = append(m, data)
			}

			for q := 0; q < len(m[0]); q++ {
				d1 := []string{}
				d2 := []string{}
				for z := 0; z < len(m); z++ {
					if q+1 < len(m[0]) {
						d1 = append(d1, m[z][q])
						d2 = append(d2, m[z][q+1])
					}
				}
				if strings.Join(d1, "") == strings.Join(d2, "") {
					rf = q
					break
				}
			}

			if rf != -1 {
				e = 0
				r2 := rf
				for q := rf; q >= 0; q-- {
					if (r2 + 1) < len(m[0]) {
						r2++
						d1 := []string{}
						d2 := []string{}
						for z := 0; z < len(m); z++ {
							d1 = append(d1, m[z][q])
							d2 = append(d2, m[z][r2])
						}
						// fmt.Println(q)
						// fmt.Println(r2)
						// fmt.Println(d1)
						// fmt.Println(d2)
						// fmt.Println()
						if strings.Join(d1, "") != strings.Join(d2, "") {
							e++
						}
					}
				}
			}

			if e == 1 {
				f = true
			}

			if rf != -1 && !f {
				sf := rf
				rf = -1
				m := [][]string{}
				for _, l := range t {
					data := strings.Split(l, "")
					m = append(m, data)
				}

				for q := sf + 1; q < len(m[0]); q++ {
					d1 := []string{}
					d2 := []string{}
					for z := 0; z < len(m); z++ {
						if q+1 < len(m[0]) {
							d1 = append(d1, m[z][q])
							d2 = append(d2, m[z][q+1])
						}
					}
					if strings.Join(d1, "") == strings.Join(d2, "") {
						rf = q
						break
					}
				}

				if rf != -1 {
					e = 0
					r2 := rf
					for q := rf; q >= 0; q-- {
						if (r2 + 1) < len(m[0]) {
							r2++
							d1 := []string{}
							d2 := []string{}
							for z := 0; z < len(m); z++ {
								d1 = append(d1, m[z][q])
								d2 = append(d2, m[z][r2])
							}
							if strings.Join(d1, "") != strings.Join(d2, "") {
								e++
							}
						}
					}
				}
			}

			if e == 1 {
				f = true
			}

			if rf != -1 && !f {
				sf := rf
				rf = -1
				m := [][]string{}
				for _, l := range t {
					data := strings.Split(l, "")
					m = append(m, data)
				}

				for q := sf + 1; q < len(m[0]); q++ {
					d1 := []string{}
					d2 := []string{}
					for z := 0; z < len(m); z++ {
						if q+1 < len(m[0]) {
							d1 = append(d1, m[z][q])
							d2 = append(d2, m[z][q+1])
						}
					}
					if strings.Join(d1, "") == strings.Join(d2, "") {
						rf = q
						break
					}
				}

				if rf != -1 {
					e = 0
					r2 := rf
					for q := rf; q >= 0; q-- {
						if (r2 + 1) < len(m[0]) {
							r2++
							d1 := []string{}
							d2 := []string{}
							for z := 0; z < len(m); z++ {
								d1 = append(d1, m[z][q])
								d2 = append(d2, m[z][r2])
							}
							if strings.Join(d1, "") != strings.Join(d2, "") {
								e++
							}
						}
					}
				}
			}

			if e == 1 {
				f = true
			}

			if !f {
				rf = -1
				m := [][]string{}
				for _, l := range t {
					data := strings.Split(l, "")
					m = append(m, data)
				}

				for q := 0; q < len(m[0]); q++ {
					e = 0
					for z := 0; z < len(m); z++ {
						if q+1 < len(m[0]) {
							if m[z][q] != m[z][q+1] {
								e++
							}
						}
					}
					if e == 1 {
						rf = q
						f = true
						break
					}
				}
			}

			if f {
				total = total + (rf + 1)
				fmt.Printf("V RF Found (%v): %v\n", y+1, rf)
				for _, l := range t {
					fmt.Println(l)
				}
				fmt.Println()
			}
		}

		if !f {
			fmt.Printf("RF NOT Found (%v): %v\n", y+1, rf)
			for _, l := range t {
				fmt.Println(l)
			}
			fmt.Println()
		}

	}
	return fmt.Sprintf("%v", total), nil
}

func (d Day13) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day13)
}
