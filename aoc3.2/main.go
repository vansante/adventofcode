package main

import (
	"fmt"
)

type Direction int

const (
	right Direction = iota
	up
	left
	down
)

func main() {
	input := 265149

	dir := right
	x, y := 500, 500
	amplitude := 1
	currentSteps := 0
	turns := 0

	grid := make([][]int, 1000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1000)
	}

	grid[x][y] = 1

	sum := 0
	for i := 1; i <= input; i++ {
		if i == 1 {
			continue
		}
		currentSteps++

		switch dir {
		case right:
			x++
		case up:
			y--
		case left:
			x--
		case down:
			y++
		}

		if currentSteps == amplitude {
			switch dir {
			case right:
				dir = up
			case up:
				dir = left
			case left:
				dir = down
			case down:
				dir = right
			}

			turns++
			if turns%2 == 0 {
				amplitude++
			}
			currentSteps = 0
		}

		sum = grid[x][y-1] + grid[x][y+1] + grid[x+1][y] + grid[x+1][y-1] + grid[x+1][y+1] + grid[x-1][y] + grid[x-1][y-1] + grid[x-1][y+1]
		grid[x][y] = sum

		if sum > input {
			break
		}
	}

	fmt.Printf("Last Coords [%d, %d]\n", x, y)
	fmt.Printf("Last Sum: %d\n", sum)
}
