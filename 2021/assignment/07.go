package assignment

import (
	"math"
	"strings"
)

type Day07 struct{}

func (d *Day07) getPositions(input string) []int64 {
	return MakeIntegers(strings.Split(strings.TrimSpace(input), ","))
}

func (d *Day07) distance(a, b int64) int64 {
	dist := a - b
	if dist < 0 {
		return b - a
	}
	return dist
}

func (d *Day07) cost(a, b int64) int64 {
	dist := a - b
	if dist < 0 {
		dist = b - a
	}
	// https://en.wikipedia.org/wiki/Triangular_number
	return dist * (dist + 1) / 2
}

func (d *Day07) SolveI(input string) int64 {
	positions := d.getPositions(input)

	max := IntegersMax(positions)

	least := int64(math.MaxInt64)
	for i := int64(0); i < max; i++ {
		fuel := int64(0)
		for _, p := range positions {
			fuel += d.distance(p, i)
		}

		if least > fuel {
			least = fuel
		}
	}

	return least
}

func (d *Day07) SolveII(input string) int64 {
	positions := d.getPositions(input)

	max := IntegersMax(positions)

	least := int64(math.MaxInt64)
	for i := int64(0); i < max; i++ {
		fuel := int64(0)
		for _, p := range positions {
			fuel += d.cost(p, i)
		}

		if least > fuel {
			least = fuel
		}
	}
	return least
}
