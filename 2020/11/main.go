package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type grid [][]string

func retrieveGrid(file string) grid {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

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

func (g grid) doTheShufflePtI() (grid, bool) {
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

func (g grid) doTheShufflePtII() (grid, bool) {
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

func (g grid) getSurroundings(y, x int) []string {
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

func (g grid) countDirectionOccupied(y, x int) int {
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

func (g grid) isDirectionOccupied(y, x, yOff, xOff int) bool {
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

func (g grid) copy() grid {
	newGrid := make(grid, len(g))
	for y := range g {
		newGrid[y] = make([]string, len(g[y]))
		for x := range g[y] {
			newGrid[y][x] = g[y][x]
		}
	}
	return newGrid
}

func (g grid) print() {
	fmt.Println()
	for y := range g {
		fmt.Println(g[y])
	}
	fmt.Println()
}

func (g grid) countOccupied() int {
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

func main() {
	wd, _ := os.Getwd()
	g := retrieveGrid(filepath.Join(wd, "11/input.txt"))

	ptI := g.copy()
	for {
		var changed bool
		ptI, changed = ptI.doTheShufflePtI()
		if !changed {
			break
		}
	}
	ptI.print()

	fmt.Printf("Part I: Amount occupied: %d\n\n", ptI.countOccupied())

	ptII := g.copy()
	for {
		var changed bool
		ptII, changed = ptII.doTheShufflePtII()
		if !changed {
			break
		}
	}
	ptII.print()

	fmt.Printf("Part II: Amount occupied: %d\n\n", ptII.countOccupied())
}
