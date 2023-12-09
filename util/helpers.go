package util

import (
	"bufio"
	"math"
	"os"
)

func File2Array(f string) ([]string, error) {
	var s []string

	file, err := os.Open(f)
	if err != nil {
		return s, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		s = append(s, t)
	}

	if err := scanner.Err(); err != nil {
		return s, err
	}

	return s, nil
}

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(GCD(a, b)))
}
