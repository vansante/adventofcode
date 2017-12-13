package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var layers = make(map[int]int)
var maxDepth = 0

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc13.2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for i := range lines {
		addLayerFromLine(lines[i])
	}

	curPos := 0
	delay := 0
	for curPos <= maxDepth {
		layer, ok := layers[curPos]
		if !ok {
			curPos++
			continue
		}

		offset := delay + curPos
		position := offset % (layer * 2 - 2)
		if position == 0 {
			fmt.Printf("Delay %d, reached %d of %d\n", delay, curPos, maxDepth)
			curPos = -1
			delay++
		}
		curPos++
	}

	fmt.Printf("Picoseconds wait: %d\n", delay)
}

func addLayerFromLine(line string) {
	parts := strings.Split(line, ": ")
	if len(parts) < 2 {
		panic("Invalid data")
	}

	depth, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		panic(err)
	}

	height, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic(err)
	}

	layers[depth] = height
	if depth > maxDepth {
		maxDepth = depth
	}
	return
}
