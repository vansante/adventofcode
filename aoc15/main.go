package main

import (
	"math"
	"fmt"
)

type Generator struct {
	factor      uint64
	value       uint64
	criterion   uint64
	judgeValues []uint64
}

const divider = 2147483647

func (gen *Generator) Next() uint64 {
	next := gen.value * gen.factor % divider
	gen.value = next
	if next%gen.criterion == 0 {
		gen.judgeValues = append(gen.judgeValues, next)
	}
	return next
}

func (gen *Generator) JudgeValue(index int) uint64 {
	if index < len(gen.judgeValues) {
		return gen.judgeValues[index]
	}
	return 0
}

func main() {
	genA := &Generator{
		factor: 16807,
		//value:  65,
		value:     618,
		criterion: 4,
	}
	genB := &Generator{
		factor: 48271,
		//value:  8921,
		value:     814,
		criterion: 8,
	}

	matchCount, selectiveMatchCount := 0, 0
	currentMatch := 0
	mask := math.MaxUint64 & uint64(math.MaxUint16)
	for i := 0; ; i++ {
		genA.Next()
		genB.Next()

		if i < 40*1000*1000 && genA.value&mask == genB.value&mask {
			matchCount++
		}

		valA := uint64(genA.JudgeValue(currentMatch))
		valB := uint64(genB.JudgeValue(currentMatch))
		if valA > 0 && valB > 0 {
			currentMatch++
			if valA&mask == valB&mask {
				selectiveMatchCount++
			}
		}

		if currentMatch == 5*1000*1000 {
			break
		}
	}

	fmt.Printf("Count: %d, Selective count: %d", matchCount, selectiveMatchCount)
}
