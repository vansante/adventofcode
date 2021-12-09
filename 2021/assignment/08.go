package assignment

import (
	"sort"
	"strconv"
	"strings"
)

type Day08 struct{}

type d08Digit []string

func (d *Day08) digitsFromStr(s []string) []d08Digit {
	dig := make([]d08Digit, len(s))
	for i := range s {
		dig[i] = strings.Split(s[i], "")
		sort.Strings(dig[i])
	}
	return dig
}

type d08Line struct {
	input     []d08Digit
	output    []d08Digit
	oneChars  []string
	fourChars []string
}

func (l *d08Line) mapLetters() {
	for i := range l.input {
		in := l.input[i]
		switch len(in) {
		case 2:
			l.oneChars = in
		case 4:
			l.fourChars = in
		}
		if len(l.oneChars) > 0 && len(l.fourChars) > 0 {
			break
		}
	}
}

func (l *d08Line) Output() string {
	v := ""
	for i := range l.output {
		out := l.output[i]

		switch len(out) {
		case 2:
			v += "1"
		case 3:
			v += "7"
		case 4:
			v += "4"
		case 5:
			if len(StringsIntersect(out, l.oneChars)) == 2 {
				v += "3"
				break
			}
			if len(StringsIntersect(out, l.fourChars)) == 3 {
				v += "5"
				break
			}
			v += "2"
		case 6:
			if len(StringsIntersect(out, l.oneChars)) == 1 {
				v += "6"
				break
			}
			if len(StringsIntersect(out, l.fourChars)) == 4 {
				v += "9"
				break
			}
			v += "0"
		case 7:
			v += "8"
		}
	}
	return v
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
			input:  d.digitsFromStr(signs),
			output: d.digitsFromStr(out),
		})
	}
	return lines
}

func (d *Day08) SolveI(input string) int64 {
	lines := d.GetInput(input)

	var count int64
	for _, line := range lines {
		for i := range line.output {
			switch len(line.output[i]) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func (d *Day08) SolveII(input string) int64 {
	lines := d.GetInput(input)

	var sum int64
	for i := range lines {
		l := lines[i]
		l.mapLetters()

		num, err := strconv.ParseInt(l.Output(), 10, 32)
		CheckErr(err)

		sum += num
	}
	return sum
}
