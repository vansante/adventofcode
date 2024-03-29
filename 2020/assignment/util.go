package assignment

import (
	"math"
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

func MakeInts(input []string) []int {
	ints := make([]int, 0, len(input))
	for i := range input {
		num, err := strconv.ParseInt(input[i], 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, int(num))
	}
	return ints
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func UniqueStrings(sl []string) []string {
	m := make(map[string]struct{}, len(sl))
	for i := range sl {
		m[sl[i]] = struct{}{}
	}

	nw := make([]string, 0, len(m))
	for s := range m {
		nw = append(nw, s)
	}
	return nw
}

func StringsContains(sl []string, s string) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func IntsContains(sl []int, n int) bool {
	for i := range sl {
		if sl[i] == n {
			return true
		}
	}
	return false
}

func IntsMin(sl []int) int {
	min := math.MaxInt
	for i := range sl {
		if sl[i] < min {
			min = sl[i]
		}
	}
	return min
}

func IntsMax(sl []int) int {
	max := math.MinInt
	for i := range sl {
		if sl[i] > max {
			max = sl[i]
		}
	}
	return max
}

func IntegersMin(sl []int64) int64 {
	min := int64(math.MaxInt64)
	for i := range sl {
		if sl[i] < min {
			min = sl[i]
		}
	}
	return min
}

func IntegersMax(sl []int64) int64 {
	max := int64(math.MinInt64)
	for i := range sl {
		if sl[i] > max {
			max = sl[i]
		}
	}
	return max
}
