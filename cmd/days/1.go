package days

type Day1 int

var day1 = Day1(1)

func (d Day1) Part1() (string, error) {
	return "Day1", nil
}

func (d Day1) Part2() (string, error) {
	return "Day1", nil
}

func (d Day1) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day1)
}
