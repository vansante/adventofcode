package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	is5_2 := true

	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc5/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")
	var data []int

	for i := range lines {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}

		data = append(data, num)
	}

	total := 0

	idx := 0
	for idx < len(data) && idx >= 0 {
		oldIdx := idx
		idx += data[idx]

		if is5_2 && data[oldIdx] >= 3 {
			data[oldIdx]--
		} else {
			data[oldIdx]++
		}

		total++

		if total > 100000000 {
			panic("Too many iterations")
		}
	}

	fmt.Printf("Total: %d", total)
}
