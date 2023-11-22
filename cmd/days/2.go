package days

type Day2 int

var day2 = Day2(2)

func (d Day2) Part1() (string, error) {
	return "Day2", nil
}

func (d Day2) Part2() (string, error) {
	return "Day2", nil
}

func (d Day2) DayInt() int {
	return int(d)
}

func init() {
	NewDay(&day2)
}
