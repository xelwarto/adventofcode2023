package days

import (
	"code/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day1 int

var day1 = Day1(1)

func (d Day1) Part1() (string, error) {
	s, err := util.File2Array("inputs/day1_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {
			var num []string
			data := strings.Split(x, "")
			for _, d := range data {
				_, err := strconv.ParseInt(d, 10, 64)
				if err == nil {
					num = append(num, d)
				}
			}

			val := fmt.Sprintf("%v%v", num[0], num[len(num)-1])
			n, _ := strconv.ParseInt(val, 10, 64)
			total = total + int(n)
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day1) Part2() (string, error) {
	s, err := util.File2Array("inputs/day1_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {

			var numS []string
			var num []string
			data := strings.Split(x, "")

			for i := 0; i < len(data); i++ {
				dig, _ := regexp.MatchString(`\d`, data[i])
				if dig {
					numS = nil
					_, err := strconv.ParseInt(data[i], 10, 64)
					if err == nil {
						num = append(num, data[i])
					}
				} else {
					numS = append(numS, data[i])

					s := strings.Join(numS, "")
					one, _ := regexp.MatchString(`one$`, s)
					if one {
						num = append(num, "1")
					}
					two, _ := regexp.MatchString(`two$`, s)
					if two {
						num = append(num, "2")
					}
					three, _ := regexp.MatchString(`three$`, s)
					if three {
						num = append(num, "3")
					}
					four, _ := regexp.MatchString(`four$`, s)
					if four {
						num = append(num, "4")
					}
					five, _ := regexp.MatchString(`five$`, s)
					if five {
						num = append(num, "5")
					}
					six, _ := regexp.MatchString(`six$`, s)
					if six {
						num = append(num, "6")
					}
					seven, _ := regexp.MatchString(`seven$`, s)
					if seven {
						num = append(num, "7")
					}
					eight, _ := regexp.MatchString(`eight$`, s)
					if eight {
						num = append(num, "8")
					}
					nine, _ := regexp.MatchString(`nine$`, s)
					if nine {
						num = append(num, "9")
					}
				}
			}
			val := fmt.Sprintf("%v%v", num[0], num[len(num)-1])
			n, _ := strconv.ParseInt(val, 10, 64)
			total = total + int(n)
		}
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day1) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day1)
}
