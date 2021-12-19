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

// StringsIntersect returns strings contained in both slices
func StringsIntersect(s1, s2 []string) []string {
	if len(s2) == 0 {
		return s1
	}
	nw := make([]string, 0, len(s1))
	for i := range s1 {
		if !StringsContains(s2, s1[i]) {
			continue
		}
		nw = append(nw, s1[i])
	}
	return nw
}

// StringsDiff returns strings from s1 not contained in s2
func StringsDiff(s1, s2 []string) []string {
	nw := make([]string, 0, len(s1))
	for i := range s1 {
		if StringsContains(s2, s1[i]) {
			continue
		}
		nw = append(nw, s1[i])
	}
	return nw
}

func StringsRemove(sl, remove []string) []string {
	nw := make([]string, 0, len(sl))
	for i := range sl {
		if StringsContains(remove, sl[i]) {
			continue
		}
		nw = append(nw, sl[i])
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

func IntsContains(sl []int, s int) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func IntsIntersect(s1, s2 []int) []int {
	if len(s2) == 0 {
		return s1
	}
	nw := make([]int, 0, len(s1))
	for i := range s1 {
		if !IntsContains(s2, s1[i]) {
			continue
		}
		nw = append(nw, s1[i])
	}
	return nw
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
