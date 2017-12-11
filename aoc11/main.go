package main

import (
	"fmt"
	"io/ioutil"
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

	x, y := 0, 0
	for i := range directions {
		x, y = directions[i].applyDirection(x, y)
		fmt.Printf("%s: %d, %d\n", directions[i], x, y)
	}

	fmt.Printf("Last Coords [%d, %d]\n", x, y)

	steps := 0
	for x != 0 || y != 0 {
		var dir Direction
		if x < 0 {
			if y <= 0 {
				dir = se
			} else if y > 0 {
				dir = ne
			}
		} else if x > 0 {
			if y <= 0 {
				dir = sw
			} else if y > 0 {
				dir = nw
			}
		} else {
			if y < 0 {
				dir = s
			} else if y > 0 {
				dir = n
			}
		}
		x, y = dir.applyDirection(x, y)
		//fmt.Printf("%s: %d, %d\n", dir, x, y)
		steps++
	}

	fmt.Printf("Distance: %d\n", steps)
}

func (dir Direction) applyDirection(x, y int) (newX, newY int) {
	newX, newY = x, y
	evenCol := x%2 == 0

	switch dir {
	case n:
		newY--
	case ne:
		if evenCol {
			newY--
		}
		newX++
	case se:
		if !evenCol {
			newY++
		}
		newX++
	case s:
		newY++
	case sw:
		if !evenCol {
			newY++
		}
		newX--
	case nw:
		if evenCol {
			newY--
		}
		newX--
	default:
		panic(dir)
	}

	return
}
