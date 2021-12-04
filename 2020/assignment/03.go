package assignment

import (
	"fmt"
	"strings"
)

type Day03 struct{}

type d03Row []bool
type d03Grid []d03Row

func (d *Day03) processLines(lines []string) d03Grid {
	var rows []d03Row
	for i := range lines {
		var r d03Row
		for _, char := range strings.Split(lines[i], "") {
			switch char {
			case ".":
				r = append(r, false)
			case "#":
				r = append(r, true)
			default:
				panic(fmt.Sprintf("unexpected char: %v", char))
			}
		}
		rows = append(rows, r)
	}
	return rows
}

type d03Slope struct {
	x int
	y int
}

func (d *Day03) findSlopeTrees(g d03Grid, sl d03Slope) int {
	var x, trees int

	for y := 0; y < len(g); y += sl.y {
		curX := x % len(g[y])
		if g[y][curX] {
			trees++
		}
		x += sl.x
	}
	return trees
}

func (d *Day03) SolveI(input string) int64 {
	g := d.processLines(SplitLines(input))
	return int64(d.findSlopeTrees(g, d03Slope{3, 1}))
}

func (d *Day03) SolveII(input string) int64 {
	g := d.processLines(SplitLines(input))

	slopes := []d03Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	result := 0
	for i := range slopes {
		trees := d.findSlopeTrees(g, slopes[i])
		if i == 0 {
			result = trees
		} else {
			result *= trees
		}
	}

	return int64(result)
}
