package assignment

import (
	"strconv"
	"strings"
)

func SplitLines(input string) []string {
	split := strings.Split(input, "\n")

	lines := make([]string, 0, len(input))
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
	ints := make([]int64, 0, len(input))
	for i := range input {
		num, err := strconv.ParseInt(input[i], 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, num)
	}
	return ints
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
