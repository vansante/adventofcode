package assignment

import (
	"fmt"
	"strings"
)

type Day08 struct{}

type d08Digit struct {
	chars      string
	candidates []int
}

func (d *d08Digit) setCandidates() {
	switch len(d.chars) {
	case 2:
		d.candidates = []int{1}
	case 3:
		d.candidates = []int{7}
	case 4:
		d.candidates = []int{4}
	case 5:
		d.candidates = []int{2, 3, 5}
	case 6:
		d.candidates = []int{0, 6, 9}
	case 7:
		d.candidates = []int{8}
	}
}

func (d *Day08) digitsFromStr(s []string) []d08Digit {
	dig := make([]d08Digit, len(s))
	for i := range s {
		dig[i] = d08Digit{
			chars: s[i],
		}
		dig[i].setCandidates()
	}
	return dig
}

type d08Line struct {
	input   []d08Digit
	output  []d08Digit
	dispMap map[string][]string
}

func (l *d08Line) getCandidate(n int) *d08Digit {
	for i := range l.input {
		c := l.input[i].candidates
		if len(c) == 1 && c[0] == n {
			return &l.input[i]
		}
	}
	return nil
}

func (l *d08Line) fillMap() {
	for _, in := range l.input {
		if len(in.candidates) != 1 {
			continue
		}

		chars := strings.Split(in.chars, "")
		switch in.candidates[0] {
		case 1:
			l.addMap("c", chars)
			l.addMap("f", chars)
		case 7:
			l.addMap("a", chars)
			l.addMap("c", chars)
			l.addMap("f", chars)
		case 4:
			l.addMap("b", chars)
			l.addMap("d", chars)
			l.addMap("c", chars)
			l.addMap("f", chars)
		case 8:
		default:
			panic("??")
		}
	}

	for k := range l.dispMap {
		l.dispMap[k] = UniqueStrings(l.dispMap[k])
	}

	l.processMap()
	l.filterDuplicates()
}

func (l *d08Line) processMap() {
	one := l.getCandidate(1)
	four := l.getCandidate(4)
	seven := l.getCandidate(7)
	eight := l.getCandidate(8)

	if one != nil {
		oneChars := strings.Split(one.chars, "")
		l.dispMap["c"] = oneChars
		l.dispMap["f"] = oneChars
	}

	if one != nil && seven != nil {
		chars := StringsIntersect(strings.Split(one.chars, ""), strings.Split(seven.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		char := StringsDiff(strings.Split(seven.chars, ""), strings.Split(one.chars, ""))
		if len(char) != 1 {
			panic(char)
		}
		l.dispMap["a"] = StringsIntersect(chars, l.dispMap["a"])
	}

	if one != nil && four != nil {
		chars := StringsIntersect(strings.Split(four.chars, ""), strings.Split(one.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		chars = StringsIntersect(strings.Split(one.chars, ""), strings.Split(four.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])

		chars = StringsDiff(strings.Split(four.chars, ""), strings.Split(one.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])
	}

	if seven != nil && four != nil {
		chars := StringsIntersect(strings.Split(four.chars, ""), strings.Split(seven.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		chars = StringsIntersect(strings.Split(seven.chars, ""), strings.Split(four.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])

		char := StringsDiff(strings.Split(seven.chars, ""), strings.Split(four.chars, ""))
		if len(char) != 1 {
			panic(char)
		}
		l.dispMap["a"] = StringsIntersect(chars, l.dispMap["a"])
	}

	if eight != nil && one != nil {
		chars := StringsIntersect(strings.Split(eight.chars, ""), strings.Split(one.chars, ""))
		if len(chars) != 2 {
			panic(chars)
		}
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		chars = StringsDiff(strings.Split(eight.chars, ""), strings.Split(one.chars, ""))
		if len(chars) != 5 {
			panic(chars)
		}
		l.dispMap["a"] = StringsIntersect(chars, l.dispMap["a"])
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])
		l.dispMap["e"] = StringsIntersect(chars, l.dispMap["e"])
		l.dispMap["g"] = StringsIntersect(chars, l.dispMap["g"])
	}

	if eight != nil && four != nil {
		chars := StringsIntersect(strings.Split(eight.chars, ""), strings.Split(four.chars, ""))
		if len(chars) != 4 {
			panic(chars)
		}
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		chars = StringsDiff(strings.Split(eight.chars, ""), strings.Split(four.chars, ""))
		if len(chars) != 3 {
			panic(chars)
		}
		l.dispMap["a"] = StringsIntersect(chars, l.dispMap["a"])
		l.dispMap["e"] = StringsIntersect(chars, l.dispMap["e"])
		l.dispMap["g"] = StringsIntersect(chars, l.dispMap["g"])
	}

	if eight != nil && seven != nil {
		chars := StringsIntersect(strings.Split(eight.chars, ""), strings.Split(seven.chars, ""))
		if len(chars) != 3 {
			panic(chars)
		}
		l.dispMap["a"] = StringsIntersect(chars, l.dispMap["a"])
		l.dispMap["c"] = StringsIntersect(chars, l.dispMap["c"])
		l.dispMap["f"] = StringsIntersect(chars, l.dispMap["f"])

		chars = StringsDiff(strings.Split(eight.chars, ""), strings.Split(seven.chars, ""))
		if len(chars) != 4 {
			panic(chars)
		}
		l.dispMap["b"] = StringsIntersect(chars, l.dispMap["b"])
		l.dispMap["d"] = StringsIntersect(chars, l.dispMap["d"])
		l.dispMap["e"] = StringsIntersect(chars, l.dispMap["e"])
		l.dispMap["g"] = StringsIntersect(chars, l.dispMap["g"])
	}
}

func (l *d08Line) addMap(char string, chars []string) {
	l.dispMap[char] = append(l.dispMap[char], chars...)
}

func (l *d08Line) filterDuplicates() {
	remove := func(char string) bool {
		for c := range l.dispMap {
			if char == l.dispMap[c][0] {
				continue
			}
			leng := len(l.dispMap[c])
			l.dispMap[c] = StringsRemove(l.dispMap[c], []string{char})
			if leng != len(l.dispMap[c]) {
				return true
			}
		}
		return false
	}

	removed := true
	for removed {
		removed = false
		for char := range l.dispMap {
			list := l.dispMap[char]
			if len(list) != 1 {
				continue
			}

			removed = remove(list[0])
			if removed {
				break
			}
		}
	}
}

func (d *Day08) GetInput(input string) []d08Line {
	in := SplitLines(input)

	lines := make([]d08Line, 0, len(in))
	for i := range in {
		split := strings.Split(in[i], " | ")
		if len(split) != 2 {
			panic("invalid input")
		}
		signs := strings.Split(split[0], " ")
		out := strings.Split(split[1], " ")
		if len(out) != 4 {
			panic("invalid input")
		}

		lines = append(lines, d08Line{
			input:   d.digitsFromStr(signs),
			output:  d.digitsFromStr(out),
			dispMap: make(map[string][]string),
		})
	}
	return lines
}

func (d *Day08) SolveI(input string) int64 {
	lines := d.GetInput(input)

	var count int64
	for _, line := range lines {
		for i := range line.output {
			out := line.output[i]

			switch out.candidates[0] {
			case 1, 4, 7, 8:
				count++
			}
		}
	}
	return count
}

func (d *Day08) SolveII(input string) int64 {
	lines := d.GetInput(input)

	for i := range lines {
		l := lines[i]
		l.fillMap()
		fmt.Println(l.dispMap)
	}

	return 0
}
