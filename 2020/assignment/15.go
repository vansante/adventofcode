package assignment

import (
	"strings"
)

type Day15 struct {
	lastSpoken       int
	lastSpokenTurn   map[int]int
	beforeSpokenTurn map[int]int
	spokenFrequency  map[int]int
}

const (
	d15TurnsPtI  = 2020
	d15TurnsPtII = 30_000_000
)

func (d *Day15) init() {
	d.lastSpoken = -1
	d.lastSpokenTurn = make(map[int]int)
	d.beforeSpokenTurn = make(map[int]int)
	d.spokenFrequency = make(map[int]int)
}

func (d *Day15) speakNumber(turn, num int) {
	d.lastSpoken = num
	d.beforeSpokenTurn[num] = d.lastSpokenTurn[num]
	d.lastSpokenTurn[num] = turn
	d.spokenFrequency[num]++
}

func (d *Day15) takeTurns(input []int, turns int) {
	for i := 0; i < turns; i++ {
		if i < len(input) {
			d.speakNumber(i, input[i])
			continue
		}
		if d.spokenFrequency[d.lastSpoken] == 1 {
			d.speakNumber(i, 0)
			continue
		}
		d.speakNumber(i, d.lastSpokenTurn[d.lastSpoken]-d.beforeSpokenTurn[d.lastSpoken])
	}
}

func (d *Day15) SolveI(input string) int64 {
	nums := MakeInts(strings.Split(input, ","))
	d.init()
	d.takeTurns(nums, d15TurnsPtI)
	return int64(d.lastSpoken)
}

func (d *Day15) SolveII(input string) int64 {
	nums := MakeInts(strings.Split(input, ","))
	d.init()
	d.takeTurns(nums, d15TurnsPtII)
	return int64(d.lastSpoken)
}
