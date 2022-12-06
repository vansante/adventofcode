package assignment

import (
	"github.com/vansante/adventofcode/2022/util"
)

type Day06 struct{}

func (d *Day06) findMarker(input string, amount int) int {
	for i := amount; i <= len(input); i++ {
		if i < amount {
			continue
		}
		cur := input[i-amount : i]

		l := util.RemoveSliceDuplicates([]rune(cur))
		if len(l) == amount {
			return i
		}
	}
	panic("not found")
}

func (d *Day06) SolveI(input string) any {
	return d.findMarker(input, 4)
}

func (d *Day06) SolveII(input string) any {
	return d.findMarker(input, 14)
}
