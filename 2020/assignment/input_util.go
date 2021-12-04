package assignment

import (
	"strconv"
	"strings"
)

func SplitLines(input string) []string {
	split := strings.Split(input, "\n")

	var lines []string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func MakeIntegers(input []string) []int64 {
	var ints []int64
	for i := range input {
		num, err := strconv.ParseInt(input[i], 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, num)
	}
	return ints
}
