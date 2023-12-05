package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day5 int

var day5 = Day5(5)

// var seeds = []int{79, 14, 55, 13}

var seeds = []int{3037945983, 743948277, 2623786093, 391282324, 195281306, 62641412, 769611781, 377903357, 2392990228, 144218002, 1179463071, 45174621, 2129467491, 226193957, 1994898626, 92402726, 1555863421, 340215202, 426882817, 207194644}

func makeMap(s []string, sr int) int {
	for _, x := range s {
		if x != "" && !strings.Contains(x, "#") {
			data := strings.Split(x, " ")

			if len(data) == 3 {
				dst, _ := strconv.Atoi(data[0])
				src, _ := strconv.Atoi(data[1])
				rng, _ := strconv.Atoi(data[2])

				if sr >= src && sr <= (src+rng) {
					for y := 0; y < rng; y++ {
						if sr == (src + y) {
							return dst + y
						}
					}
				}
			}
		}
	}
	return sr
}

func (d Day5) Part1() (string, error) {
	total := 0
	totals := []int{}

	s1, _ := util.File2Array("inputs/day5_1.txt")
	s2, _ := util.File2Array("inputs/day5_2.txt")
	s3, _ := util.File2Array("inputs/day5_3.txt")
	s4, _ := util.File2Array("inputs/day5_4.txt")
	s5, _ := util.File2Array("inputs/day5_5.txt")
	s6, _ := util.File2Array("inputs/day5_6.txt")
	s7, _ := util.File2Array("inputs/day5_7.txt")

	for _, seed := range seeds {
		i1 := makeMap(s1, seed)
		i2 := makeMap(s2, i1)
		i3 := makeMap(s3, i2)
		i4 := makeMap(s4, i3)
		i5 := makeMap(s5, i4)
		i6 := makeMap(s6, i5)
		i7 := makeMap(s7, i6)

		totals = append(totals, i7)
	}

	for _, t := range totals {
		if total == 0 {
			total = t
		} else {
			if t < total {
				total = t
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##################################################################################
func (d Day5) Part2() (string, error) {
	total := 0

	return fmt.Sprintf("%v", total), nil
}

func (d Day5) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day5)
}
