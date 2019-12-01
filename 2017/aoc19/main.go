package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"unicode"
)

type Direction int

const (
	right Direction = iota
	up
	left
	down
)

var directions = []Direction{up, right, down, left}

func (dir Direction) applyDirection(x, y int) (newX, newY int) {
	newX, newY = x, y

	switch dir {
	case up:
		newY--
	case right:
		newX++
	case down:
		newY++
	case left:
		newX--
	default:
		panic(dir)
	}
	return
}

func (dir Direction) getAtDirection(grid [][]string, x, y int) (result string) {
	switch dir {
	case up:
		result = grid[y-1][x]
	case right:
		result = grid[y][x+1]
	case down:
		result = grid[y+1][x]
	case left:
		result = grid[y][x-1]
	default:
		panic(dir)
	}
	return
}

func (dir Direction) getOpposing() (opposing Direction) {
	switch dir {
	case up:
		opposing = down
	case right:
		opposing = left
	case down:
		opposing = up
	case left:
		opposing = right
	default:
		panic(dir)
	}
	return
}

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc19/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	grid := make([][]string, len(lines))
	for i := range lines {
		grid[i] = strings.Split(lines[i], "")
	}

	x, y := 0, 0
	dir := down
	for i := range grid[0] {
		if grid[0][i] == "|" {
			x = i
			break
		}
	}

	letters := ""
	steps := 0

WalkLoop:
	for {
		if grid[y][x] == " " {
			break WalkLoop
		}

		char := []rune(grid[y][x])[0]

		if unicode.IsLetter(char) {
			letters += grid[y][x]
		} else {
			switch grid[y][x] {
			case "+":
				found := false
				for _, posDir := range directions {
					if posDir != dir.getOpposing() && posDir.getAtDirection(grid, x, y) != " " {
						dir = posDir
						found = true
						break
					}
				}
				if !found {
					break WalkLoop
				}
			case "-":
			case "|":
			default:
				panic(fmt.Sprintf("%d, %d: %s", x, y, grid[y][x]))
			}
		}
		x, y = dir.applyDirection(x, y)
		steps++
	}

	fmt.Printf("Letters: %s, Steps: %d\n", letters, steps)
}
