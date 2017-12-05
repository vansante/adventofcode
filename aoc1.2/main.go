package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc1.2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	var data []int

	for i := 0; i < len(inputStr); i++ {
		curDigit, err := strconv.Atoi(inputStr[i:i+1])
		if err != nil {
			panic(err)
		}
		data = append(data, curDigit)
	}

	half := len(data) / 2

	total := 0
	for i := 0; i < len(data); i++ {
		idx := i + half
		if idx > len(data) - 1 {
			idx = idx % len(data)
		}

		if data[i] == data[idx] {
			total += data[i]
		}
	}

	fmt.Printf("Total: %d", total)
}
