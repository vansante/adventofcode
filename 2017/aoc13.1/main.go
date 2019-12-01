package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Layer struct {
	depth  int
	height int

	curHeight int
	down      bool
}

var layers = make(map[int]*Layer)
var maxDepth = 0

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc13.1/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for i := range lines {
		addLayerFromLine(lines[i])
	}

	_, severity := travel(0, false)
	fmt.Printf("Severity at picosecond zero: %d\n", severity)
}

func travel(delay int, stopOnCollision bool) (collision bool, severity int) {
	resetLayers()

	severity = 0
	for i := 0; i < delay; i++ {
		progressLayers()
	}

	for i := 0; i <= maxDepth; i++ {
		layer, found := layers[i]
		if found && layer.collision(0) {
			severity += i * layer.height
			collision = true
			if stopOnCollision {
				return
			}
		}
		progressLayers()
	}
	return
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

	layers[depth] = &Layer{
		depth:  depth,
		height: height,
	}
	if depth > maxDepth {
		maxDepth = depth
	}
	return
}

func progressLayers() {
	for i := range layers {
		if layers[i].down {
			layers[i].curHeight++

			if layers[i].curHeight > layers[i].height-1 {
				layers[i].down = false
				layers[i].curHeight -= 2
			}
		} else {
			layers[i].curHeight--

			if layers[i].curHeight < 0 {
				layers[i].down = true
				layers[i].curHeight = 1
			}
		}
	}
}

func resetLayers() {
	for i := range layers {
		layers[i].curHeight = 0
		layers[i].down = true
	}
}

func (l *Layer) collision(height int) bool {
	return l.curHeight == height
}
