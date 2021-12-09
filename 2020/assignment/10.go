package assignment

import (
	"sort"
)

type Day10 struct{}

func (d *Day10) findDifferences(current int64, numbers []int64) (diff1, diff2, diff3 int) {
	for i := range numbers {
		var diff int64
		switch numbers[i] {
		case current + 1:
			diff1++
			diff = 1
		case current + 2:
			diff2++
			diff = 2
		case current + 3:
			diff3++
			diff = 3
		default:
			continue
		}
		sub1, sub2, sub3 := d.findDifferences(current+diff, numbers)
		diff1 += sub1
		diff2 += sub2
		diff3 += sub3
		break
	}
	return diff1, diff2, diff3
}

func (d *Day10) findCombinations(current int64, numbers []int64, values map[int64]int64) int64 {
	if len(numbers) == 0 {
		return 1
	}
	if val, ok := values[current]; ok {
		return val
	}

	var val int64
	for i := range numbers {
		if numbers[i] <= current || numbers[i] > current+3 {
			continue
		}
		val += d.findCombinations(numbers[i], numbers[i+1:], values)
	}
	values[current] = val
	return val
}

func (d *Day10) SolveI(input string) int64 {
	numbers := MakeIntegers(SplitLines(input))

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	diff1, _, diff3 := d.findDifferences(0, numbers)
	diff3++ // For the device adapter

	return int64(diff1) * int64(diff3)
}

func (d *Day10) SolveII(input string) int64 {
	numbers := MakeIntegers(SplitLines(input))

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	adapters := make(map[int64]int64)
	combinations := d.findCombinations(0, numbers, adapters)

	return combinations
}
