package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Day15 int

var day15 = Day15(15)

func (d Day15) Part1() (string, error) {
	s, err := util.File2Array("inputs/day15_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {
			data := strings.Split(x, ",")
			for _, d := range data {
				cur := 0
				code := strings.Split(d, "")
				for _, c := range code {
					r, _ := utf8.DecodeRuneInString(c)
					cur = cur + int(r)
					cur = cur * 17
					cur = cur % 256
				}
				total = total + cur
			}
		}
	}
	return fmt.Sprintf("%v", total), nil
}

func Box(s string) int {
	box := 0
	code := strings.Split(s, "")
	for _, c := range code {
		r, _ := utf8.DecodeRuneInString(c)
		box = box + int(r)
		box = box * 17
		box = box % 256
	}
	return box
}

func (d Day15) Part2() (string, error) {
	s, err := util.File2Array("inputs/day15_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	boxes := [][]map[string]int{}
	for q := 0; q < 256; q++ {
		boxes = append(boxes, []map[string]int{})
	}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, ",")
			for _, d := range data {

				if strings.Contains(d, "=") {
					data2 := strings.Split(d, "=")
					lens, _ := strconv.Atoi(data2[1])
					label := make(map[string]int)
					label[data2[0]] = lens
					b := Box(data2[0])
					box := boxes[b]

					f := false
					for q := 0; q < len(box); q++ {
						b := box[q]
						if b[data2[0]] != 0 {
							b[data2[0]] = lens
							f = true
						}
					}
					if !f {
						box = append(box, label)
					}
					boxes[b] = box

				}

				if strings.Contains(d, "-") {
					data2 := strings.Split(d, "-")
					b := Box(data2[0])
					box := boxes[b]

					nbox := []map[string]int{}
					for q := 0; q < len(box); q++ {
						b := box[q]
						if b[data2[0]] == 0 {
							nbox = append(nbox, b)
						}
					}
					boxes[b] = nbox
				}

			}
		}
	}

	for y, box := range boxes {
		for q := 0; q < len(box); q++ {
			for _, v := range box[q] {
				total = total + ((y + 1) * (q + 1) * v)
			}
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day15) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day15)
}
