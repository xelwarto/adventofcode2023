package days

type Day3 int

var day3 = Day3(3)

func (d Day3) Part1() (string, error) {
	return "Day3", nil
}

func (d Day3) Part2() (string, error) {
	return "Day3", nil
}

func (d Day3) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day3)
}
