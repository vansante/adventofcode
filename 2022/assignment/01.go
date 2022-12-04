package assignment

import (
	"math"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day01 struct{}

func (d *Day01) getElveCalories(input string) [][]int64 {
	elves := strings.Split(input, "\n\n")
	elvesCalories := make([][]int64, len(elves))
	for i := range elves {
		elvesCalories[i] = util.ParseInt64s(util.SplitLines(elves[i]))
	}
	return elvesCalories
}

func (d *Day01) SolveI(input string) int64 {
	elves := d.getElveCalories(input)
	max := int64(math.MinInt64)
	for _, elve := range elves {
		elveCals := util.SumSlice(elve)
		if elveCals > max {
			max = elveCals
		}
	}
	return max
}

func (d *Day01) SolveII(input string) int64 {
	elves := d.getElveCalories(input)
	calories := make([]int64, len(elves))
	for i, elve := range elves {
		calories[i] = util.SumSlice(elve)
	}
	util.SliceSort(calories, false)
	return calories[0] + calories[1] + calories[2]
}
