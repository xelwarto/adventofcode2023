package days

import (
	"code/util"
	"fmt"
	"strings"
)

type Day20 int

var day20 = Day20(20)

var l_pulse = 0
var h_pulse = 0
var b_push = int64(0)

func RecordPulse(p string) {
	if p == "low" {
		l_pulse++
	}

	if p == "high" {
		h_pulse++
	}
}

func DummyPulse(q []string) {
	// fmt.Printf("%v -%v-> %v\n", q[2], q[1], q[0])
	RecordPulse(q[1])
}

var queue [][]string

func DeQueue() {
	for len(queue) > 0 {
		q := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = [][]string{}
		}

		if mods[q[0]] != nil {
			m := mods[q[0]]
			m.Pulse(q[1], q[2])
		} else {
			DummyPulse(q)
		}
	}
}

var mods map[string]Module

type Module interface {
	GetName() string
	GetType() string
	AddOutput(string)
	AddState(string)
	Pulse(string, string)
	GetPulses() int64
}

// Broadcaster Module
// There is a single broadcast module (named broadcaster).
// When it receives a pulse, it sends the same pulse to all of its destination modules.

type Broadcaster struct {
	Name     string
	Type     string
	Output   []string
	Receiver string
}

func NewBroadcaster(name string) Broadcaster {
	mod := Broadcaster{
		Name:   name,
		Type:   "Broadcaster",
		Output: []string{},
	}
	return mod
}

func (mod Broadcaster) GetName() string {
	return mod.Name
}

func (mod Broadcaster) GetType() string {
	return mod.Type
}

func (mod Broadcaster) GetPulses() int64 {
	return 0
}

func (mod Broadcaster) AddOutput(m string) {
	mod.Output = append(mod.Output, m)
	mods[mod.GetName()] = mod
}

func (mod Broadcaster) AddState(from string) {
}

func (mod Broadcaster) Pulse(p string, from string) {
	// fmt.Printf("%v -%v-> %v\n", from, p, mod.GetName())
	mod.Receiver = p
	RecordPulse(p)
	mods[mod.GetName()] = mod

	for _, m := range mod.Output {
		q := []string{m, p, mod.GetName()}
		queue = append(queue, q)
	}
	DeQueue()
}

// FlipFlop Module
// Flip-flop modules (prefix %) are either on or off; they are initially off.
// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
// However, if a flip-flop module receives a low pulse, it flips between on and off.
// If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.
type FlipFlop struct {
	Name     string
	Type     string
	On       bool
	Output   []string
	Receiver string
}

func NewFlipFlop(name string) FlipFlop {
	mod := FlipFlop{
		Name:   name,
		Type:   "FlipFlop",
		On:     false,
		Output: []string{},
	}
	return mod
}

func (mod FlipFlop) GetName() string {
	return mod.Name
}

func (mod FlipFlop) GetType() string {
	return mod.Type
}

func (mod FlipFlop) GetPulses() int64 {
	return 0
}

func (mod FlipFlop) AddOutput(m string) {
	mod.Output = append(mod.Output, m)
	mods[mod.GetName()] = mod
}

func (mod FlipFlop) AddState(from string) {
}

func (mod FlipFlop) Pulse(p string, from string) {
	// fmt.Printf("%v -%v-> %v\n", from, p, mod.GetName())
	mod.Receiver = p
	RecordPulse(p)
	if p == "low" {
		send := ""
		if mod.On {
			mod.On = false
			send = "low"
		} else {
			mod.On = true
			send = "high"
		}

		for _, m := range mod.Output {
			q := []string{m, send, mod.GetName()}
			queue = append(queue, q)
		}
	}
	mods[mod.GetName()] = mod
}

// Conjunction Module
// Conjunction modules (prefix &) remember the type of the most recent pulse received from
// each of their connected input modules; they initially default to remembering a low pulse for each input.
// When a pulse is received, the conjunction module first updates its memory for that input.
// Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.
type Conjunction struct {
	Name   string
	Type   string
	On     bool
	Output []string
	State  map[string]string
	Pulses int64
}

func NewConjunction(name string) Conjunction {
	mod := Conjunction{
		Name:   name,
		Type:   "Conjunction",
		On:     false,
		Output: []string{},
		State:  make(map[string]string),
		Pulses: 0,
	}
	return mod
}

func (mod Conjunction) GetName() string {
	return mod.Name
}

func (mod Conjunction) GetType() string {
	return mod.Type
}

func (mod Conjunction) GetPulses() int64 {
	return mod.Pulses
}

func (mod Conjunction) AddState(from string) {
	mod.State[from] = "low"
	mods[mod.GetName()] = mod
}

func (mod Conjunction) AddOutput(m string) {
	mod.Output = append(mod.Output, m)
	mods[mod.GetName()] = mod
}

func (mod Conjunction) Pulse(p string, from string) {
	// fmt.Printf("%v -%v-> %v\n", from, p, mod.GetName())
	RecordPulse(p)
	mod.State[from] = p

	send := "high"
	h := 0
	for _, s := range mod.State {
		if s == "high" {
			h++
		}
	}
	if h == len(mod.State) {
		send = "low"
	}

	if send == "high" {
		mod.Pulses = b_push
	}

	for _, m := range mod.Output {
		q := []string{m, send, mod.GetName()}
		queue = append(queue, q)
	}
	mods[mod.GetName()] = mod
}

func (d Day20) Part1() (string, error) {
	s, err := util.File2Array("inputs/day20_part1.txt")
	if err != nil {
		return "", err
	}

	total := 0
	mods = make(map[string]Module)
	for _, x := range s {
		if x != "" {
			x = strings.ReplaceAll(x, " ", "")
			data1 := strings.Split(x, "->")

			if data1[0] == "broadcaster" {
				m := NewBroadcaster("broadcaster")
				mods["broadcaster"] = m
			}

			if strings.Contains(data1[0], "%") {
				data1[0] = strings.ReplaceAll(data1[0], "%", "")
				m := NewFlipFlop(data1[0])
				mods[data1[0]] = m
			}

			if strings.Contains(data1[0], "&") {
				data1[0] = strings.ReplaceAll(data1[0], "&", "")
				m := NewConjunction(data1[0])
				mods[data1[0]] = m
			}
		}
	}

	for _, x := range s {
		if x != "" {
			x = strings.ReplaceAll(x, " ", "")
			data1 := strings.Split(x, "->")
			data2 := strings.Split(data1[1], ",")

			if strings.Contains(data1[0], "%") {
				data1[0] = strings.ReplaceAll(data1[0], "%", "")
			}

			if strings.Contains(data1[0], "&") {
				data1[0] = strings.ReplaceAll(data1[0], "&", "")
			}

			for _, d := range data2 {
				mods[data1[0]].AddOutput(d)

				if mods[d] != nil {
					if mods[d].GetType() == "Conjunction" {
						mods[d].AddState(data1[0])
					}
				}
			}

		}
	}

	for q := 0; q < 10000; q++ {
		mods["broadcaster"].Pulse("low", "button")
	}

	fmt.Printf("\nHigh: %v / Low: %v\n", h_pulse, l_pulse)
	total = h_pulse * l_pulse

	return fmt.Sprintf("%v", total), nil
}

func (d Day20) Part2() (string, error) {
	s, err := util.File2Array("inputs/day20_part2.txt")
	if err != nil {
		return "", err
	}

	// total := 0
	mods = make(map[string]Module)
	for _, x := range s {
		if x != "" {
			x = strings.ReplaceAll(x, " ", "")
			data1 := strings.Split(x, "->")

			if data1[0] == "broadcaster" {
				m := NewBroadcaster("broadcaster")
				mods["broadcaster"] = m
			}

			if strings.Contains(data1[0], "%") {
				data1[0] = strings.ReplaceAll(data1[0], "%", "")
				m := NewFlipFlop(data1[0])
				mods[data1[0]] = m
			}

			if strings.Contains(data1[0], "&") {
				data1[0] = strings.ReplaceAll(data1[0], "&", "")
				m := NewConjunction(data1[0])
				mods[data1[0]] = m
			}
		}
	}

	for _, x := range s {
		if x != "" {
			x = strings.ReplaceAll(x, " ", "")
			data1 := strings.Split(x, "->")
			data2 := strings.Split(data1[1], ",")

			if strings.Contains(data1[0], "%") {
				data1[0] = strings.ReplaceAll(data1[0], "%", "")
			}

			if strings.Contains(data1[0], "&") {
				data1[0] = strings.ReplaceAll(data1[0], "&", "")
			}

			for _, d := range data2 {
				mods[data1[0]].AddOutput(d)

				if mods[d] != nil {
					if mods[d].GetType() == "Conjunction" {
						mods[d].AddState(data1[0])
					}
				}
			}

		}
	}

	for {
		b_push++
		mods["broadcaster"].Pulse("low", "button")
		// fmt.Println(mods["pr"].GetPulses())
		if b_push > 4000 {
			break
		}
	}

	total := util.LCM(
		util.LCM(
			util.LCM(
				mods["fv"].GetPulses(),
				mods["pr"].GetPulses()),
			mods["bt"].GetPulses()),
		mods["rd"].GetPulses())

	return fmt.Sprintf("%v", total), nil
}

func (d Day20) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day20)
}
