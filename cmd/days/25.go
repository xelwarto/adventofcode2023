package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day25 int

var day25 = Day25(25)

var comps map[string][]string

func ResolveComp(comp string, list *[]string, data map[string][]string) {
	if data[comp] != nil {
		f := false
		for _, l := range *list {
			if l == comp {
				f = true
			}
		}

		if !f {
			*list = append(*list, comp)
		}

		for _, c := range data[comp] {
			f := false
			for _, l := range *list {
				if l == c {
					f = true
				}
			}

			if !f {
				ResolveComp(c, list, data)
			}
		}
	}
}

func CutWire(s []string, data map[string][]string) map[string][]string {
	c := data[s[0]]
	list := []string{}
	for _, n := range c {
		if n != s[1] {
			list = append(list, n)
		}
	}
	data[s[0]] = list

	c = data[s[1]]
	list = []string{}
	for _, n := range c {
		if n != s[0] {
			list = append(list, n)
		}
	}
	data[s[1]] = list

	return data
}

func CompsMatch(list1 []string, list2 []string) bool {
	f := false
	for _, l1 := range list1 {
		if !f {
			for _, l2 := range list2 {
				if l1 == l2 {
					f = true
					break
				}
			}
		}
	}
	return f
}

func CopyMap(m map[string][]string) map[string][]string {
	n := make(map[string][]string)

	for k, v := range m {
		n[k] = v
	}
	return n
}

func (d Day25) Part1() (string, error) {
	s, err := util.File2Array("inputs/day25_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	comps = make(map[string][]string)
	links := [][]string{}
	for _, x := range s {
		if x != "" {
			data1 := strings.Split(x, ": ")
			if comps[data1[0]] == nil {
				comps[data1[0]] = []string{}
			}

			data2 := strings.Split(data1[1], " ")
			for _, d := range data2 {
				comps[data1[0]] = append(comps[data1[0]], d)

				if comps[d] == nil {
					comps[d] = []string{}
				}
				comps[d] = append(comps[d], data1[0])

				links = append(links, []string{data1[0], d})
			}
		}
	}

	// fmt.Println(links)

	tracker := make(map[string]int)
	ch := 0
	for k, _ := range comps {
		tracker[k] = 0
		for _, l := range links {
			if k == l[0] || k == l[1] {
				tracker[k]++
			}
		}

		if tracker[k] > ch {
			ch = tracker[k]
		}
	}

	for z := 0; z < len(links); z++ {
		fmt.Println(z)
		p1 := links[z]
		if z+1 < len(links) {
			for x := z + 1; x < len(links); x++ {
				p2 := links[x]
				if x+1 < len(links) {
					for y := x + 1; y < len(links); y++ {
						p3 := links[y]

						// fmt.Println(p1)
						// fmt.Println(p2)
						// fmt.Println(p3)

						ncomps := CopyMap(comps)
						ncomps = CutWire([]string{p1[0], p1[1]}, ncomps)
						ncomps = CutWire([]string{p2[0], p2[1]}, ncomps)
						ncomps = CutWire([]string{p3[0], p3[1]}, ncomps)

						list1 := []string{}
						ResolveComp(p1[0], &list1, ncomps)

						list2 := []string{}
						ResolveComp(p1[1], &list2, ncomps)

						// fmt.Println(len(list1) + len(list2))
						// fmt.Println(time.Now())

						if !CompsMatch(list1, list2) {
							// fmt.Println(len(list1) + len(list2))
							i := len(list1) * len(list2)
							fmt.Printf("\nANSWER: %v\n\n", i)
						}
					}
				}
			}
		}
	}

	// ncomps := comps
	// // hfx/pzl, bvb/cmg, nvd/jqt,
	// ncomps = CutWire([]string{"pzl", "hfx"}, ncomps)
	// ncomps = CutWire([]string{"bvb", "cmg"}, ncomps)
	// ncomps = CutWire([]string{"nvd", "jqt"}, ncomps)

	// list1 := []string{}
	// ResolveComp("bvb", &list1, ncomps)
	// fmt.Println(list1)

	// list2 := []string{}
	// ResolveComp("cmg", &list2, ncomps)
	// fmt.Println(list2)

	// fmt.Println(CompsMatch(list1, list2))

	return fmt.Sprintf("%v", total), nil
}

func (d Day25) Part2() (string, error) {
	s, err := util.File2Array("inputs/day25_part2.txt")
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

func (d Day25) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day25)
}
