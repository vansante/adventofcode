package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var pipes = make(map[int][]int)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc12/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for i := range lines {
		addPipesFromLine(lines[i])
	}

	connected := make(map[int]int)
	for pipe := range pipes {
		connected[pipe] = -1
	}

	for i := 0; i < len(pipes); i++ {
		for pipe := range pipes {
			if connected[pipe] <= pipe {
				if connected[pipe] < 0 {
					connected[pipe] = pipe
				}
				connections := pipes[pipe]
				for j := range connections {
					if connected[connections[j]] > connected[pipe] {
						connected[connections[j]] = connected[pipe]
					}
				}
			}
		}
	}

	totalZero := 0
	groups := make(map[int]bool)
	for pipe := range connected {
		groups[connected[pipe]] = true
		if connected[pipe] == 0 {
			totalZero++
		}
	}

	fmt.Printf("Connected to group zero: %d, Total groups: %d\n", totalZero, len(groups))
}

func addPipesFromLine(line string) {
	parts := strings.Split(line, " <-> ")
	if len(parts) < 2 {
		panic("Invalid data")
	}

	fromParts := strings.Split(parts[0], ",")
	toParts := strings.Split(parts[1], ",")
	var from []int
	var to []int

	for i := range fromParts {
		pipe, err := strconv.Atoi(strings.TrimSpace(fromParts[i]))
		if err != nil {
			panic(err)
		}
		from = append(from, pipe)
	}

	for i := range toParts {
		pipe, err := strconv.Atoi(strings.TrimSpace(toParts[i]))
		if err != nil {
			panic(err)
		}
		to = append(to, pipe)
	}

	for i := range from {
		for j := range to {
			curFrom := from[i]
			curTo := to[j]

			pipes[curFrom] = append(pipes[curFrom], curTo)
			pipes[curTo] = append(pipes[curTo], curFrom)
		}
	}
	return
}
