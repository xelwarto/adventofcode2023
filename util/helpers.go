package util

import (
	"bufio"
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
