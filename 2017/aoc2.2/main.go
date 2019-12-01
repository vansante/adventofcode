package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc2.2/input.txt")
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

	DivideLoop:
		for j := 0; j < len(data[i]); j++ {
			for a := len(data[i]) - 1; a >= 0; a-- {
				if a == j {
					continue
				}

				if data[i][j] % data[i][a] == 0 {
					total += data[i][j] / data[i][a]
					break DivideLoop
				}

				if data[i][a] % data[i][j] == 0 {
					total += data[i][a] / data[i][j]
					break DivideLoop
				}
			}
		}

	}

	fmt.Printf("Total: %d", total)
}
