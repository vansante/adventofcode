package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc1.1/input.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	inputStr := string(input)
	if len(inputStr) > 0 {
		prevDigit := -1
		for i := 0; i < len(input); i++ {
			curDigit, err := strconv.Atoi(inputStr[i:i+1])
			if err != nil {
				panic(err)
			}

			if curDigit == prevDigit {
				total += curDigit
			}

			prevDigit = curDigit
		}

		firstDigit, err := strconv.Atoi(inputStr[0:1])
		if err != nil {
			panic(err)
		}

		if firstDigit == prevDigit {
			total += prevDigit
		}
	}

	fmt.Printf("Total: %d", total)
}
