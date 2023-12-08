package days

import (
	"code/util"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Day8 int

var day8 = Day8(8)

type Step struct {
	Left  string
	Right string
}

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(GCD(a, b)))
}

func (d Day8) Part1() (string, error) {
	guideInput := `LRRLRRRLRRLLLRLLRRLRRLLRRRLRRLLRLRRRLRLRRLRLRRRLRLRLRRLLRLRLRRLRRRLRRRLRRRLRLRRLLLLRLLRLLRRLRRRLLLRLRRRLRLRRRLRLRRLRRRLRRRLRLRLLRRRLLRLLRLRLRLRLLRRLRRLRRRLRRLRLRLRLRLRRLRRRLLRRRLLRLLLRRRLLRRRLRRRLRRLRLRRLRLLRRLLRRLRLRLRRLRLRRLLRRRLLRRRLLRLRRRLRLRRRLRLRRRLRRRLRRLRRLRRLLRRRLRRRLLLRRRR`
	guide := strings.Split(guideInput, "")
	steps := make(map[string]Step)
	s, err := util.File2Array("inputs/day8_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {
			x1 := strings.ReplaceAll(x, " ", "")
			x1 = strings.ReplaceAll(x1, "(", "")
			x1 = strings.ReplaceAll(x1, ")", "")
			data := strings.Split(x1, "=")
			data2 := strings.Split(data[1], ",")

			step := Step{
				Left:  data2[0],
				Right: data2[1],
			}

			steps[data[0]] = step
		}
	}

	c := 0
	l := "AAA"
	for {
		if c >= len(guide) {
			c = 0
		}
		d := guide[c]
		s := steps[l]

		if d == "L" {
			l = s.Left
		}

		if d == "R" {
			l = s.Right
		}
		total++

		if l == "ZZZ" {
			break
		}
		c++
	}

	return fmt.Sprintf("%v", total), nil
}

// ##################################################################################
func (d Day8) Part2() (string, error) {
	guideInput := `LRRLRRRLRRLLLRLLRRLRRLLRRRLRRLLRLRRRLRLRRLRLRRRLRLRLRRLLRLRLRRLRRRLRRRLRRRLRLRRLLLLRLLRLLRRLRRRLLLRLRRRLRLRRRLRLRRLRRRLRRRLRLRLLRRRLLRLLRLRLRLRLLRRLRRLRRRLRRLRLRLRLRLRRLRRRLLRRRLLRLLLRRRLLRRRLRRRLRRLRLRRLRLLRRLLRRLRLRLRRLRLRRLLRRRLLRRRLLRLRRRLRLRRRLRLRRRLRRRLRRLRRLRRLLRRRLRRRLLLRRRR`
	guide := strings.Split(guideInput, "")
	steps := make(map[string]Step)
	s, err := util.File2Array("inputs/day8_part2.txt")
	if err != nil {
		return "", err
	}

	l := []string{}
	total := int64(0)
	for _, x := range s {
		if x != "" {
			x1 := strings.ReplaceAll(x, " ", "")
			x1 = strings.ReplaceAll(x1, "(", "")
			x1 = strings.ReplaceAll(x1, ")", "")
			data := strings.Split(x1, "=")
			data2 := strings.Split(data[1], ",")

			step := Step{
				Left:  data2[0],
				Right: data2[1],
			}
			steps[data[0]] = step

			re := regexp.MustCompile(`..A`)
			if re.MatchString(data[0]) {
				l = append(l, data[0])
			}
		}
	}

	c := 0
	totals := []int64{}
	for _, x := range l {
		for {
			if c >= len(guide) {
				c = 0
			}
			d := guide[c]
			s := steps[x]

			if d == "L" {
				x = s.Left
			}

			if d == "R" {
				x = s.Right
			}
			total++

			re := regexp.MustCompile(`..Z`)
			if re.MatchString(x) {
				break
			}
			c++
		}
		totals = append(totals, total)
		total = 0
		c = 0
	}

	t := int64(totals[0])
	for q := 1; q < len(totals); q++ {
		t = Lcm(totals[q], t)
	}

	return fmt.Sprintf("%v", t), nil
}

func (d Day8) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day8)
}
