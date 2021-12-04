package assignment

import (
	"fmt"
	"strings"
)

type Day11 struct{}

type d11Grid [][]string

func (d *Day11) retrieveGrid(in string) d11Grid {
	split := strings.Split(in, "\n")

	var input [][]string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		input = append(input, strings.Split(line, ""))
	}
	return input
}

func (g d11Grid) doTheShufflePtI() (d11Grid, bool) {
	newGrid := g.copy()
	var changed bool
	for y := range g {
		for x := range g[y] {
			s := g[y][x]
			if s == "." {
				continue
			}

			occupied := 0
			seats := g.getSurroundings(y, x)
			for _, seat := range seats {
				if seat == "#" {
					occupied++
				}
			}
			switch s {
			case "L":
				if occupied == 0 {
					newGrid[y][x] = "#"
					changed = true
				}
			case "#":
				if occupied >= 4 {
					newGrid[y][x] = "L"
					changed = true
				}
			}
		}
	}
	return newGrid, changed
}

func (g d11Grid) doTheShufflePtII() (d11Grid, bool) {
	newGrid := g.copy()
	var changed bool
	for y := range g {
		for x := range g[y] {
			s := g[y][x]
			if s == "." {
				continue
			}

			occupied := g.countDirectionOccupied(y, x)
			switch s {
			case "L":
				if occupied == 0 {
					newGrid[y][x] = "#"
					changed = true
				}
			case "#":
				if occupied >= 5 {
					newGrid[y][x] = "L"
					changed = true
				}
			}
		}
	}
	return newGrid, changed
}

func (g d11Grid) getSurroundings(y, x int) []string {
	var seats []string
	if y-1 >= 0 {
		if x-1 >= 0 {
			seats = append(seats, g[y-1][x-1])
		}
		seats = append(seats, g[y-1][x])
		if x+1 < len(g[y]) {
			seats = append(seats, g[y-1][x+1])
		}
	}
	if x-1 >= 0 {
		seats = append(seats, g[y][x-1])
	}
	if x+1 < len(g[y]) {
		seats = append(seats, g[y][x+1])
	}
	if y+1 < len(g) {
		if x-1 >= 0 {
			seats = append(seats, g[y+1][x-1])
		}
		seats = append(seats, g[y+1][x])
		if x+1 < len(g[y]) {
			seats = append(seats, g[y+1][x+1])
		}
	}
	return seats
}

func (g d11Grid) countDirectionOccupied(y, x int) int {
	total := 0
	if g.isDirectionOccupied(y, x, -1, -1) {
		total++
	}
	if g.isDirectionOccupied(y, x, -1, 0) {
		total++
	}
	if g.isDirectionOccupied(y, x, -1, 1) {
		total++
	}
	if g.isDirectionOccupied(y, x, 0, -1) {
		total++
	}
	if g.isDirectionOccupied(y, x, 0, 1) {
		total++
	}
	if g.isDirectionOccupied(y, x, 1, -1) {
		total++
	}
	if g.isDirectionOccupied(y, x, 1, 0) {
		total++
	}
	if g.isDirectionOccupied(y, x, 1, 1) {
		total++
	}
	return total
}

func (g d11Grid) isDirectionOccupied(y, x, yOff, xOff int) bool {
	for {
		y += yOff
		x += xOff
		if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
			return false
		}
		switch g[y][x] {
		case "#":
			return true
		case "L":
			return false
		}
	}
}

func (g d11Grid) copy() d11Grid {
	newGrid := make(d11Grid, len(g))
	for y := range g {
		newGrid[y] = make([]string, len(g[y]))
		for x := range g[y] {
			newGrid[y][x] = g[y][x]
		}
	}
	return newGrid
}

func (g d11Grid) print() {
	fmt.Println()
	for y := range g {
		fmt.Println(g[y])
	}
	fmt.Println()
}

func (g d11Grid) countOccupied() int {
	total := 0
	for y := range g {
		for x := range g[y] {
			if g[y][x] == "#" {
				total++
			}
		}
	}
	return total
}

func (d *Day11) SolveI(input string) int64 {
	g := d.retrieveGrid(input)

	for {
		var changed bool
		g, changed = g.doTheShufflePtI()
		if !changed {
			break
		}
	}
	g.print()
	return int64(g.countOccupied())
}

func (d *Day11) SolveII(input string) int64 {
	g := d.retrieveGrid(input)

	for {
		var changed bool
		g, changed = g.doTheShufflePtII()
		if !changed {
			break
		}
	}
	g.print()
	return int64(g.countOccupied())
}
