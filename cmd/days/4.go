package days

import (
	"code/util"
	"fmt"
	"regexp"
	"strings"
)

type Day4 int

var day4 = Day4(4)

func (d Day4) Part1() (string, error) {
	s, err := util.File2Array("inputs/day4_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	for _, x := range s {
		if x != "" {
			winners := []string{}
			cardData1 := strings.Split(x, ": ")
			cardData2 := strings.Split(cardData1[1], " | ")

			re := regexp.MustCompile(`\d+`)
			wNums := re.FindAllString(cardData2[0], -1)
			myNums := re.FindAllString(cardData2[1], -1)

			for _, m := range myNums {
				for _, w := range wNums {
					if w == m {
						winners = append(winners, m)
					}
				}
			}

			t := 0
			for x := 0; x < len(winners); x++ {
				if t == 0 {
					t = 1
				} else {
					t = t * 2
				}

			}
			total = total + t
		}
	}

	return fmt.Sprintf("%v", total), nil
}

type Card struct {
	Copies  int
	Winners int
}

func (d Day4) Part2() (string, error) {
	s, err := util.File2Array("inputs/day4_part2.txt")
	if err != nil {
		return "", err
	}

	cards := []Card{}
	for x := 0; x < len(s); x++ {
		c := Card{
			Copies:  1,
			Winners: 0,
		}
		cards = append(cards, c)
	}

	total := 0
	for y, x := range s {
		if x != "" {
			winners := []string{}
			cardData1 := strings.Split(x, ": ")
			cardData2 := strings.Split(cardData1[1], " | ")

			re := regexp.MustCompile(`\d+`)
			wNums := re.FindAllString(cardData2[0], -1)
			myNums := re.FindAllString(cardData2[1], -1)

			for _, m := range myNums {
				for _, w := range wNums {
					if w == m {
						winners = append(winners, m)
					}
				}
			}
			cards[y].Winners = len(winners)
		}
	}

	for y, card := range cards {
		for z := 0; z < card.Copies; z++ {

			w := card.Winners
			// fmt.Printf("%v: %v\n", y, w)

			for q := 1; q <= w; q++ {
				cards[y+q].Copies++
			}
		}

	}

	for _, card := range cards {
		// fmt.Printf("%v: %v\n", y, card.Copies)
		total = total + card.Copies
	}

	return fmt.Sprintf("%v", total), nil
}

func (d Day4) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day4)
}
