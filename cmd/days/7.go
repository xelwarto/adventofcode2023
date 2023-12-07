package days

import (
	"code/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day7 int

var day7 = Day7(7)

// var cards = []string{`2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `T`, `J`, `Q`, `K`, `A`}

var cards = []string{`J`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `T`, `Q`, `K`, `A`}

func cardToInt(s string) int {
	for x := 0; x < len(cards); x++ {
		if s == cards[x] {
			return x
		}
	}
	return -1
}

func orderHands(r []string, h string) []string {
	res := []string{}
	data1 := strings.Split(h, " ")
	hand := strings.Split(data1[0], "")

	f := false
	for _, x := range r {
		data := strings.Split(x, " ")
		y := strings.Split(data[0], "")

		if !f {
			for q := 0; q < len(y); q++ {
				i1 := cardToInt(hand[q])
				i2 := cardToInt(y[q])

				if i1 == i2 {
					continue
				}

				if i1 > i2 {
					res = append(res, h)
					res = append(res, x)
					f = true
					break
				} else {
					res = append(res, x)
					break
				}
			}
		} else {
			res = append(res, x)
		}
	}
	if !f {
		res = append(res, h)
	}
	return res
}

func (d Day7) Part1() (string, error) {
	s, err := util.File2Array("inputs/day7_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0

	r6 := []string{} //Five of a kind
	r5 := []string{} //Four of a kind
	r4 := []string{} //Full house
	r3 := []string{} //Three of a kind
	r2 := []string{} //Two pair
	r1 := []string{} //One pair
	r0 := []string{}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, " ")

			f := false
			// Five of a kind
			for q := 0; q < len(cards); q++ {
				re := regexp.MustCompile(cards[q])
				if len(re.FindAllString(data[0], -1)) == 5 {
					r6 = orderHands(r6, x)
					f = true
					break
				}
			}

			if !f {
				// Four of a kind
				for q := 0; q < len(cards); q++ {
					re := regexp.MustCompile(cards[q])
					if len(re.FindAllString(data[0], -1)) == 4 {
						r5 = orderHands(r5, x)
						f = true
						break
					}
				}
			}

			if !f {
				//Full house
				c3 := false
				c2 := false
				for q := 0; q < len(cards); q++ {
					re := regexp.MustCompile(cards[q])
					if len(re.FindAllString(data[0], -1)) == 3 {
						c3 = true
					}
					if len(re.FindAllString(data[0], -1)) == 2 {
						c2 = true
					}
				}
				if c2 && c3 {
					r4 = orderHands(r4, x)
					f = true
				}
			}

			if !f {
				//Three of a kind
				for q := 0; q < len(cards); q++ {
					re := regexp.MustCompile(cards[q])
					if len(re.FindAllString(data[0], -1)) == 3 {
						r3 = orderHands(r3, x)
						f = true
						break
					}
				}
			}

			if !f {
				//Two pair
				c := 0
				for q := 0; q < len(cards); q++ {
					re := regexp.MustCompile(cards[q])
					if len(re.FindAllString(data[0], -1)) == 2 {
						c++
					}
				}
				if c == 2 {
					r2 = orderHands(r2, x)
					f = true
				}
			}

			if !f {
				//One pair
				for q := 0; q < len(cards); q++ {
					re := regexp.MustCompile(cards[q])
					if len(re.FindAllString(data[0], -1)) == 2 {
						r1 = orderHands(r1, x)
						f = true
						break
					}
				}
			}

			if !f {
				r0 = orderHands(r0, x)
			}
		}
	}

	res := []string{}
	for _, x6 := range r6 {
		res = append(res, x6)
	}
	for _, x5 := range r5 {
		res = append(res, x5)
	}
	for _, x4 := range r4 {
		res = append(res, x4)
	}
	for _, x3 := range r3 {
		res = append(res, x3)
	}
	for _, x2 := range r2 {
		res = append(res, x2)
	}
	for _, x1 := range r1 {
		res = append(res, x1)
	}
	for _, x0 := range r0 {
		res = append(res, x0)
	}

	for q := 0; q < len(res); q++ {
		data := strings.Split(res[q], " ")
		i, _ := strconv.Atoi(data[1])
		total = total + (i * (len(res) - q))
	}

	return fmt.Sprintf("%v", total), nil
}

// ##################################################################################
func (d Day7) Part2() (string, error) {
	s, err := util.File2Array("inputs/day7_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0

	r6 := []string{} //Five of a kind
	r5 := []string{} //Four of a kind
	r4 := []string{} //Full house
	r3 := []string{} //Three of a kind
	r2 := []string{} //Two pair
	r1 := []string{} //One pair
	r0 := []string{}

	for _, x := range s {
		if x != "" {
			data := strings.Split(x, " ")
			re := regexp.MustCompile(`J`)
			hand := re.ReplaceAllString(data[0], "*")

			f := false
			// Five of a kind
			for q := 0; q < len(cards); q++ {

				exp := fmt.Sprintf("[%s*]", cards[q])
				re := regexp.MustCompile(exp)
				if len(re.FindAllString(hand, -1)) == 5 {
					r6 = orderHands(r6, x)
					f = true
					break
				}
			}

			if !f {
				// Four of a kind
				for q := 0; q < len(cards); q++ {
					exp := fmt.Sprintf("[%s*]", cards[q])
					re := regexp.MustCompile(exp)
					if len(re.FindAllString(hand, -1)) == 4 {
						r5 = orderHands(r5, x)
						f = true
						break
					}
				}
			}

			if !f {
				//Full house

				if !strings.Contains(hand, "*") {
					c3 := false
					c2 := false
					for q := 0; q < len(cards); q++ {
						exp := fmt.Sprintf("[%s*]", cards[q])
						re := regexp.MustCompile(exp)
						if len(re.FindAllString(hand, -1)) == 3 {
							c3 = true
						}
						if len(re.FindAllString(hand, -1)) == 2 {
							c2 = true
						}
					}
					if c2 && c3 {
						r4 = orderHands(r4, x)
						f = true
					}
				}

				if strings.Contains(hand, "*") {
					re := regexp.MustCompile(`[*]`)
					if len(re.FindAllString(hand, -1)) == 1 {
						c3 := false
						c2 := false
						for q := 0; q < len(cards); q++ {
							re := regexp.MustCompile(cards[q])
							if len(re.FindAllString(hand, -1)) == 2 {
								if !c3 {
									c3 = true
								} else {
									c2 = true
								}
							}
						}
						if c2 && c3 {
							r4 = orderHands(r4, x)
							f = true
						}
					}

					if len(re.FindAllString(hand, -1)) == 2 {
						for q := 0; q < len(cards); q++ {
							re := regexp.MustCompile(cards[q])
							if len(re.FindAllString(hand, -1)) == 2 {
								r4 = orderHands(r4, x)
								f = true
								break
							}
						}
					}
				}
			}

			if !f {
				//Three of a kind
				for q := 0; q < len(cards); q++ {
					exp := fmt.Sprintf("[%s*]", cards[q])
					re := regexp.MustCompile(exp)
					if len(re.FindAllString(hand, -1)) == 3 {
						r3 = orderHands(r3, x)
						f = true
						break
					}
				}
			}

			if !f {
				//Two pair
				c := 0
				for q := 0; q < len(cards); q++ {
					exp := fmt.Sprintf("[%s*]", cards[q])
					re := regexp.MustCompile(exp)
					if len(re.FindAllString(hand, -1)) == 2 {
						c++
					}
				}
				if c == 2 {
					r2 = orderHands(r2, x)
					f = true
				}
			}

			if !f {
				//One pair
				for q := 0; q < len(cards); q++ {
					exp := fmt.Sprintf("[%s*]", cards[q])
					re := regexp.MustCompile(exp)
					if len(re.FindAllString(hand, -1)) == 2 {
						r1 = orderHands(r1, x)
						f = true
						break
					}
				}
			}

			if !f {
				r0 = orderHands(r0, x)
			}
		}
	}

	res := []string{}
	for _, x6 := range r6 {
		res = append(res, x6)
	}
	for _, x5 := range r5 {
		res = append(res, x5)
	}
	for _, x4 := range r4 {
		res = append(res, x4)
	}
	for _, x3 := range r3 {
		res = append(res, x3)
	}
	for _, x2 := range r2 {
		res = append(res, x2)
	}
	for _, x1 := range r1 {
		res = append(res, x1)
	}
	for _, x0 := range r0 {
		res = append(res, x0)
	}

	for q := 0; q < len(res); q++ {
		data := strings.Split(res[q], " ")
		i, _ := strconv.Atoi(data[1])
		total = total + (i * (len(res) - q))
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day7) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day7)
}
