package days

import (
	"code/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day3 int

var day3 = Day3(3)

func (d Day3) Part1() (string, error) {
	s, err := util.File2Array("inputs/day3_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	foundNums := []string{}
	for y, x := range s {
		if x != "" {
			data := strings.Split(x, "")
			re3 := regexp.MustCompile(`\d+`)
			nums3 := re3.FindAllStringIndex(x, -1)

			for _, x1 := range nums3 {
				f := false
				b := []string{}
				for a := x1[0]; a < x1[1]; a++ {
					b = append(b, data[a])
				}
				c := strings.Join(b, "")

				if x1[0] > 0 {

					re1 := regexp.MustCompile(`[@#$%&*\-+=/]`)
					nums1 := re1.MatchString(data[x1[0]-1])
					if nums1 {
						foundNums = append(foundNums, c)
						f = true
					}
				}

				if !f && x1[1] != 140 {
					re1 := regexp.MustCompile(`[@#$%&*\-+=/]`)
					nums1 := re1.MatchString(data[x1[1]])
					if nums1 {
						foundNums = append(foundNums, c)
						f = true
					}
				}

				if y > 0 && !f {
					for q := x1[0] - 1; q < (x1[1] + 1); q++ {
						if q >= 0 && q < len(data) {
							w := s[y-1]
							data1 := strings.Split(w, "")
							re := regexp.MustCompile(`[@#$%&*\-+=/]`)
							if re.MatchString(data1[q]) {
								foundNums = append(foundNums, c)
								f = true
								// fmt.Printf("%s ", c1)
								break
							}
						}
					}
				}

				if y < len(s)-1 && !f {
					for q := x1[0] - 1; q < (x1[1] + 1); q++ {
						if q >= 0 && q < len(data) {
							w := s[y+1]
							data1 := strings.Split(w, "")
							re := regexp.MustCompile(`[@#$%&*\-+=/]`)
							if re.MatchString(data1[q]) {
								foundNums = append(foundNums, c)
								// fmt.Printf("%s ", c1)
								break
							}
						}
					}
				}
			}
			// fmt.Println()
		}
	}

	for _, n := range foundNums {
		re := regexp.MustCompile(`[@#$%&*\-+=/]`)
		r := re.ReplaceAllString(n, "")
		// fmt.Println(r)

		i, err := strconv.Atoi(r)
		if err != nil {
			return "", err
		}
		total = total + i
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day3) Part2() (string, error) {
	s, err := util.File2Array("inputs/day3_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0
	foundNums := []int{}
	for y, x := range s {
		if x != "" {
			re := regexp.MustCompile(`\*`)
			gears := re.FindAllStringIndex(x, -1)

			dataLine := strings.Split(x, "")
			reLine := regexp.MustCompile(`\d+`)
			numsLine := reLine.FindAllStringIndex(x, -1)

			for _, g := range gears {
				var nums []int
				gear := g[0]

				// up
				if (y - 1) >= 0 {
					lineUp := s[y-1]
					dataUp := strings.Split(lineUp, "")
					reUp := regexp.MustCompile(`\d+`)
					numsUp := reUp.FindAllStringIndex(lineUp, -1)

					for _, n := range numsUp {
						if gear >= (n[0]-1) && gear < (n[1]+1) {
							u := []string{}
							for a := n[0]; a < n[1]; a++ {
								u = append(u, dataUp[a])
							}
							v, _ := strconv.Atoi(strings.Join(u, ""))
							nums = append(nums, v)
							// fmt.Println(v)
						}
					}
				}

				// right
				if (gear - 1) >= 0 {
					for _, n := range numsLine {
						if (gear) == (n[1]) {
							u := []string{}
							for a := n[0]; a < n[1]; a++ {
								u = append(u, dataLine[a])
							}
							v, _ := strconv.Atoi(strings.Join(u, ""))
							nums = append(nums, v)
						}
					}
				}

				// left
				if (gear + 1) < len(dataLine) {
					for _, n := range numsLine {
						if (gear + 1) == (n[0]) {
							u := []string{}
							for a := n[0]; a < n[1]; a++ {
								u = append(u, dataLine[a])
							}
							v, _ := strconv.Atoi(strings.Join(u, ""))
							nums = append(nums, v)
							// fmt.Println(v)
						}
					}
				}

				// down
				if (y + 1) < len(s) {
					lineDown := s[y+1]
					dataDown := strings.Split(lineDown, "")
					reDown := regexp.MustCompile(`\d+`)
					numsDown := reDown.FindAllStringIndex(lineDown, -1)

					for _, n := range numsDown {
						if gear >= (n[0]-1) && gear < (n[1]+1) {
							u := []string{}
							for a := n[0]; a < n[1]; a++ {
								u = append(u, dataDown[a])
							}
							v, _ := strconv.Atoi(strings.Join(u, ""))
							nums = append(nums, v)
							// fmt.Println(v)
						}
					}
				}

				if len(nums) == 2 {
					// fmt.Println(nums)
					foundNums = append(foundNums, (nums[0] * nums[1]))
				}

			}
		}
	}

	for _, n := range foundNums {
		// fmt.Println(n)
		total = total + n
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day3) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day3)
}
