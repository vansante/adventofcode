package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var history = make(map[string]int)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc6.2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	nums := strings.Split(inputStr, "\t")
	var data []int

	for i := range nums {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

		data = append(data, num)
	}

	result := 0
	iteration := 0
	for {
		highest := -1
		highestIdx := -1
		for i := range data {
			if data[i] > highest {
				highest = data[i]
				highestIdx = i
			}
		}

		redistr := data[highestIdx]
		data[highestIdx] = 0

		for i := highestIdx + 1; ; i++ {
			if i > len(data)-1 {
				i = 0
			}

			data[i]++
			redistr--

			if redistr == 0 {
				break
			}
		}

		if existsInHistory(data) {
			result = iteration - getPreviousIteration(data)
			break
		}
		addToHistory(iteration, data)

		iteration++
	}

	fmt.Printf("Result: %d", result)
}

func formatHistory(seq []int) string {
	var nums []string
	for i := range seq {
		nums = append(nums, fmt.Sprintf("%d", seq[i]))
	}
	return strings.Join(nums, "_")
}

func addToHistory(iteration int, seq []int) {
	seqStr := formatHistory(seq)
	history[seqStr] = iteration
}

func existsInHistory(seq []int) bool {
	seqStr := formatHistory(seq)
	return history[seqStr] > 0
}

func getPreviousIteration(seq []int) int {
	seqStr := formatHistory(seq)
	return history[seqStr]
}
