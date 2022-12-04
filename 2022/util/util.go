package util

import (
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
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

func ParseInt64s(input []string) []int64 {
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

func ParseInts(input []string) []int {
	ints := make([]int, 0, len(input))
	for i := range input {
		num, err := strconv.ParseInt(input[i], 10, 32)
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

func SliceRemoveAll[T comparable](sl, remove []T) []T {
	nw := make([]T, 0, len(sl))
	for i := range sl {
		if SliceContains(remove, sl[i]) {
			continue
		}
		nw = append(nw, sl[i])
	}
	return nw
}

func SliceContains[T comparable](sl []T, s T) bool {
	for i := range sl {
		if sl[i] == s {
			return true
		}
	}
	return false
}

func RemoveSliceDuplicates[T comparable](sl []T) []T {
	m := make(map[T]struct{}, len(sl))
	for i := range sl {
		m[sl[i]] = struct{}{}
	}

	nw := make([]T, 0, len(m))
	for s := range m {
		nw = append(nw, s)
	}
	return nw
}

// SliceDiff returns items in s1 not contained in s2
func SliceDiff(s1, s2 []string) []string {
	if len(s2) == 0 {
		return s1
	}
	nw := make([]string, 0, len(s1))
	for i := range s1 {
		if SliceContains(s2, s1[i]) {
			continue
		}
		nw = append(nw, s1[i])
	}
	return nw
}

// SliceIntersect returns items from s1 also contained in s2
func SliceIntersect[T comparable](s1, s2 []T) []T {
	if len(s2) == 0 {
		return s1
	}
	nw := make([]T, 0, len(s1))
	for i := range s1 {
		if !SliceContains(s2, s1[i]) {
			continue
		}
		nw = append(nw, s1[i])
	}
	return nw
}

func SliceSort[T constraints.Ordered](s []T, ascending bool) {
	sort.Slice(s, func(i, j int) bool {
		if ascending {
			return s[i] < s[j]
		}
		return s[i] > s[j]
	})
}

func Abs[T constraints.Signed](s T) T {
	if s < 0 {
		return -s
	}
	return s
}

func Max[T constraints.Ordered](s1, s2 T) T {
	if s1 > s2 {
		return s1
	}
	return s2
}

func Min[T constraints.Ordered](s1, s2 T) T {
	if s1 < s2 {
		return s1
	}
	return s2
}

func MaxSlice[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func MinSlice[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func SumSlice[T constraints.Signed](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
