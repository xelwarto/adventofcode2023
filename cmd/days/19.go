package days

import (
	"code/util"
	"fmt"
	"strconv"
	"strings"
)

type Day19 int

var day19 = Day19(19)

// px{a<2006:qkq,m>2090:A,rfg}

type Rule struct {
	Part     string
	Operator string
	Value    int
	Workflow string
}

// {x=787,m=2655,a=1222,s=2876}
type Part struct {
	X        int
	M        int
	A        int
	S        int
	Accepted bool
	Rejected bool
}

var workflows map[string][]Rule

func NextWorkflow(w string, p Part) string {
	rules := workflows[w]
	for _, r := range rules {
		if r.Operator == "" {
			return r.Workflow
		}

		partValue := -1
		if r.Part == "a" {
			partValue = p.A
		}
		if r.Part == "m" {
			partValue = p.M
		}
		if r.Part == "s" {
			partValue = p.S
		}
		if r.Part == "x" {
			partValue = p.X
		}

		if partValue != -1 {
			if r.Operator == ">" {
				if partValue > r.Value {
					return r.Workflow
				}
			}

			if r.Operator == "<" {
				if partValue < r.Value {
					return r.Workflow
				}
			}
		}
	}
	return ""
}

func findEnd(p Part) string {
	end := "in"
	for end != "A" && end != "R" {
		end = NextWorkflow(end, p)
	}
	return end
}

func (d Day19) Part1() (string, error) {
	s, err := util.File2Array("inputs/day19_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0

	readParts := false
	workflows = make(map[string][]Rule)
	parts := []Part{}
	for _, x := range s {
		if x != "" {
			if readParts {
				x = strings.Replace(x, "{", "", -1)
				x = strings.Replace(x, "}", "", -1)
				data := strings.Split(x, ",")

				p := Part{
					Accepted: false,
					Rejected: false,
				}
				for _, d := range data {
					l := strings.Split(d, "=")
					if len(l) == 2 {

						if l[0] == "x" {
							i, _ := strconv.Atoi(l[1])
							p.X = i
						}

						if l[0] == "m" {
							i, _ := strconv.Atoi(l[1])
							p.M = i
						}

						if l[0] == "a" {
							i, _ := strconv.Atoi(l[1])
							p.A = i
						}

						if l[0] == "s" {
							i, _ := strconv.Atoi(l[1])
							p.S = i
						}

					}
				}
				parts = append(parts, p)
			} else {
				// rfg{s<537:gd,x>2440:R,A}
				rules := []Rule{}
				data := strings.Split(x, "{")
				rs := strings.Replace(data[1], "}", "", -1)
				rdata := strings.Split(rs, ",")
				for _, r := range rdata {
					rule := Rule{}

					if strings.Contains(r, ":") {
						data1 := strings.Split(r, ":")
						rule.Workflow = data1[1]
						if strings.Contains(data1[0], ">") {
							rule.Operator = ">"
							data2 := strings.Split(data1[0], ">")
							i, _ := strconv.Atoi(data2[1])
							rule.Value = i
							rule.Part = data2[0]
						}

						if strings.Contains(data1[0], "<") {
							rule.Operator = "<"
							data2 := strings.Split(data1[0], "<")
							i, _ := strconv.Atoi(data2[1])
							rule.Value = i
							rule.Part = data2[0]
						}
					} else {
						rule.Workflow = r
					}

					rules = append(rules, rule)
				}

				workflows[data[0]] = rules
			}
		} else {
			readParts = true
		}
	}

	for _, p := range parts {
		e := findEnd(p)
		if e == "A" {
			total = total + p.A + p.M + p.S + p.X
		}
	}

	return fmt.Sprintf("%v", total), nil
}

// ##############################################################################################
// 167409079868000

func Combos(w string) int {
	val := 0

	return val
}

func (d Day19) Part2() (string, error) {
	s, err := util.File2Array("inputs/day19_part2.txt")
	if err != nil {
		return "", err
	}

	total := 0

	readParts := false
	workflows = make(map[string][]Rule)
	parts := []Part{}
	for _, x := range s {
		if x != "" {
			if readParts {
				x = strings.Replace(x, "{", "", -1)
				x = strings.Replace(x, "}", "", -1)
				data := strings.Split(x, ",")

				p := Part{
					Accepted: false,
					Rejected: false,
				}
				for _, d := range data {
					l := strings.Split(d, "=")
					if len(l) == 2 {

						if l[0] == "x" {
							i, _ := strconv.Atoi(l[1])
							p.X = i
						}

						if l[0] == "m" {
							i, _ := strconv.Atoi(l[1])
							p.M = i
						}

						if l[0] == "a" {
							i, _ := strconv.Atoi(l[1])
							p.A = i
						}

						if l[0] == "s" {
							i, _ := strconv.Atoi(l[1])
							p.S = i
						}

					}
				}
				parts = append(parts, p)
			} else {
				// rfg{s<537:gd,x>2440:R,A}
				rules := []Rule{}
				data := strings.Split(x, "{")
				rs := strings.Replace(data[1], "}", "", -1)
				rdata := strings.Split(rs, ",")
				for _, r := range rdata {
					rule := Rule{}

					if strings.Contains(r, ":") {
						data1 := strings.Split(r, ":")
						rule.Workflow = data1[1]
						if strings.Contains(data1[0], ">") {
							rule.Operator = ">"
							data2 := strings.Split(data1[0], ">")
							i, _ := strconv.Atoi(data2[1])
							rule.Value = i
							rule.Part = data2[0]
						}

						if strings.Contains(data1[0], "<") {
							rule.Operator = "<"
							data2 := strings.Split(data1[0], "<")
							i, _ := strconv.Atoi(data2[1])
							rule.Value = i
							rule.Part = data2[0]
						}
					} else {
						rule.Workflow = r
					}

					rules = append(rules, rule)
				}

				workflows[data[0]] = rules
			}
		} else {
			readParts = true
		}
	}

	total = Combos("px")

	return fmt.Sprintf("%v", total), nil
}

func (d Day19) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day19)
}
