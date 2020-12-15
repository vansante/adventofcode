package main

import "fmt"

const (
	turnsPtI  = 2020
	turnsPtII = 30000000
)

var (
	input            = []int{9, 6, 0, 10, 18, 2, 1}
	lastSpoken       = -1
	lastSpokenTurn   = make(map[int]int)
	beforeSpokenTurn = make(map[int]int)
	spokenFrequency  = make(map[int]int)
)

func reset() {
	lastSpoken = -1
	lastSpokenTurn = make(map[int]int)
	beforeSpokenTurn = make(map[int]int)
	spokenFrequency = make(map[int]int)
}

func speakNumber(turn, num int) {
	//fmt.Printf("Speak number: %d\n", num)
	lastSpoken = num
	beforeSpokenTurn[num] = lastSpokenTurn[num]
	lastSpokenTurn[num] = turn
	spokenFrequency[num]++
}

func takeTurns(turns int) {
	for i := 0; i < turns; i++ {
		if i < len(input) {
			speakNumber(i, input[i])
			continue
		}
		if spokenFrequency[lastSpoken] == 1 {
			speakNumber(i, 0)
			continue
		}
		speakNumber(i, lastSpokenTurn[lastSpoken]-beforeSpokenTurn[lastSpoken])
	}
}

func main() {
	takeTurns(turnsPtI)
	fmt.Printf("The last number spoken was: %d\n\n", lastSpoken)

	reset()
	takeTurns(turnsPtII)
	fmt.Printf("The last number spoken was: %d\n\n", lastSpoken)
}
