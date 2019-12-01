package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc10.1/input.txt")
	if err != nil {
		panic(err)
	}

	var list []int
	for i := 0; i <= 255; i++ {
		list = append(list, i)
	}

	numbers := strings.Split(string(input), ",")

	var lengths []int
	for i := range numbers {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic(err)
		}
		lengths = append(lengths, num)
	}

	currentPosition := 0
	skipSize := 0
	listLength := len(list)
	for i := range lengths {
		curLength := lengths[i]
		startIdx := currentPosition
		endIdx := (currentPosition + curLength - 1) % listLength
		//fmt.Printf("CurLength %d. CurPos %d, StartIdx %d, EndIdx %d\n", curLength, currentPosition, startIdx, endIdx)

		for j := 0; j < curLength/2; j++ {
			//fmt.Printf("Swap %d <=> %d\n", startIdx, endIdx)
			list[startIdx], list[endIdx] = list[endIdx], list[startIdx]

			startIdx = (startIdx + 1) % listLength
			endIdx--
			if endIdx < 0 {
				endIdx += listLength
			}
		}

		currentPosition = (currentPosition + curLength + skipSize) % listLength
		skipSize++
	}

	fmt.Printf("List: %#v\n", list)
	fmt.Printf("Result: %d\n", list[0]*list[1])
}
