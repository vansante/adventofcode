package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc9/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := strings.Split(string(input), "")

	lastExclamation := false
	inGarbage := false
	groups := 0
	groupLevel := 0
	groupScore := 0
	garbageCount := 0

	for i := range inputStr {
		character := inputStr[i]
		if inGarbage {
			if lastExclamation {
				lastExclamation = false
				continue
			}
			if character == ">" {
				inGarbage = false
			} else if character != "!" {
				garbageCount++
			}
			lastExclamation = character == "!"
		} else {
			switch character {
			case "{":
				groups++
				groupLevel++
				groupScore += groupLevel
			case "}":
				groupLevel--
			case "<":
				inGarbage = true
			}
		}
	}

	fmt.Printf("Group total: %d, Score: %d, Garbage characters: %d\n", groups, groupScore, garbageCount)
}
