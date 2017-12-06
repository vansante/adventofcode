package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var history []string

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc6.1/input.txt")
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

	total := 0

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

		total++
		if existsInHistory(data) {
			break
		}
		addToHistory(data)
	}

	fmt.Printf("Total: %d", total)
}

func formatHistory(seq []int) string {
	var nums []string
	for i := range seq {
		nums = append(nums, fmt.Sprintf("%d", seq[i]))
	}

	return strings.Join(nums, "_")
}

func addToHistory(seq []int) {
	history = append(history, formatHistory(seq))
}

func existsInHistory(seq []int) bool {
	seqStr := formatHistory(seq)

	for i := range history {
		if seqStr == history[i] {
			return true
		}
	}
	return false
}
