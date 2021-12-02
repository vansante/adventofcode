package assignment

import (
	"strconv"
	"strings"
)

type Day02 struct {}

const (
	d02Up = iota + 1
	d02Down
	d02Forward
)

type D02Vector struct {
	Direction int
	Amount int
}

func (d *Day02) SplitInput(lines []string) []D02Vector {
	vs := make([]D02Vector, 0, len(lines))
	for _, line := range lines {
		v := D02Vector{}
		spl := strings.Split(line, " ")
		switch spl[0] {
		case "up":
			v.Direction = d02Up
		case "down":
			v.Direction = d02Down
		case "forward":
			v.Direction = d02Forward
		default:
			panic("unknown direction")
		}
		var err error
		v.Amount, err = strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}

		vs = append(vs, v)
	}
	return vs
}

func (d *Day02) SolveI(input string) int64 {
	vs := d.SplitInput(SplitLines(input))

	var horizontal, vertical int64
	for _, v := range vs {
		switch v.Direction {
		case d02Up:
			vertical -= int64(v.Amount)
		case d02Down:
			vertical += int64(v.Amount)
		case d02Forward:
			horizontal += int64(v.Amount)
		}
	}

	return horizontal * vertical
}

func (d *Day02) SolveII(input string) int64 {
	vs := d.SplitInput(SplitLines(input))

	var horizontal, vertical, aim int64
	for _, v := range vs {
		switch v.Direction {
		case d02Up:
			aim -= int64(v.Amount)
		case d02Down:
			aim += int64(v.Amount)
		case d02Forward:
			horizontal += int64(v.Amount)
			vertical += int64(v.Amount) * aim
		}
	}
	return horizontal * vertical
}