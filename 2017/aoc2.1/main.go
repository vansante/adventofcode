package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc2.1/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	data := make([][]int, len(lines))

	for i := range lines {
		nums := strings.Split(lines[i], "\t")

		for j := range nums {
			num, err := strconv.Atoi(nums[j])
			if err != nil {
				panic(err)
			}
			data[i] = append(data[i], num)
		}
	}

	total := 0
	for i := 0; i < len(data); i++ {
		largest := -1
		smallest := math.MaxInt32

		for j := 0; j < len(data[i]); j++ {
			if data[i][j] < smallest {
				smallest = data[i][j]
			}

			if data[i][j] > largest {
				largest = data[i][j]
			}
		}
		total += largest - smallest
	}

	fmt.Printf("Total: %d", total)
}
