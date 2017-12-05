package main

import (
	"fmt"
	"math"
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
	x, y := 0, 0
	amplitude := 1
	currentSteps := 0
	turns := 0

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
			if turns % 2 == 0 {
				amplitude++
			}
			currentSteps = 0
		}

		if i < 30 {
			fmt.Printf("%d: [%d, %d]\n", i, x, y)
		}
	}

	fmt.Printf("Last Coords [%d, %d]\n", x, y)

	xAbs := math.Abs(float64(x))
	yAbs := math.Abs(float64(y))
	fmt.Printf("Distance: %d\n", int(xAbs) + int(yAbs))
}
