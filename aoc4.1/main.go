package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc4.1/input.txt")
	if err != nil {
		panic(err)
	}

	// Set to false for 4.1, true for 4.2
	dupeDetect := true

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")
	data := make([][]string, len(lines))

	for i := range lines {
		data[i] = strings.Split(lines[i], " ")

		if dupeDetect {
			for j := range data[i] {

				s := strings.Split(data[i][j], "")
				sort.Strings(s)
				data[i][j] = strings.Join(s, "")
			}
		}
	}

	total := 0
	for i := 0; i < len(data); i++ {
		found := false

	DuplicateLoop:
		for j := 0; j < len(data[i]); j++ {
			for a := len(data[i]) - 1; a >= 0; a-- {
				if a == j {
					continue
				}

				if data[i][j] == data[i][a] {
					found = true
					break DuplicateLoop
				}
			}
		}

		if !found {
			total++
		}
	}

	fmt.Printf("Total: %d", total)
}
