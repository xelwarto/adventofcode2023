package days

func init() {
	d := addDay(1)

	d.Part1 = func() (string, error) {
		return "1", nil
	}

	d.Part2 = func() (string, error) {
		return "2", nil
	}
}
