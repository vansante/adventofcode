package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Direction string

const (
	n  Direction = "n"
	ne Direction = "ne"
	se Direction = "se"
	s  Direction = "s"
	sw Direction = "sw"
	nw Direction = "nw"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc11/input.txt")
	if err != nil {
		panic(err)
	}

	var directions []Direction
	dirArr := strings.Split(string(input), ",")
	for i := range dirArr {
		directions = append(directions, Direction(dirArr[i]))
	}

	x, y, z := 0, 0, 0
	distance := 0
	maxDistance := 0
	for i := range directions {
		x, y, z = directions[i].applyDirection(x, y, z)

		distance = int(math.Abs(float64(x))+math.Abs(float64(y))+math.Abs(float64(z))) / 2
		if distance > maxDistance {
			maxDistance = distance
		}
		fmt.Printf("%s: %d, %d, %d\n", directions[i], x, y, z)
	}

	fmt.Printf("Last Coords [%d, %d, %d]\n", x, y, z)

	distance = int(math.Abs(float64(x))+math.Abs(float64(y))+math.Abs(float64(z))) / 2

	fmt.Printf("Distance: %d, Max distance: %d\n", distance, maxDistance)
}

func (dir Direction) applyDirection(x, y, z int) (newX, newY, newZ int) {
	newX, newY, newZ = x, y, z

	switch dir {
	case n:
		newY++
		newZ--
	case ne:
		newX++
		newZ--
	case se:
		newX++
		newY--
	case s:
		newY--
		newZ++
	case sw:
		newX--
		newZ++
	case nw:
		newX--
		newY++
	default:
		panic(dir)
	}

	return
}
