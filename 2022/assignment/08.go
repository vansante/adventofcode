package assignment

import (
	"log"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day08 struct{}

type d08Coord struct {
	x, y int
}

type d08Grid struct {
	y []d08Line
}

type d08Line struct {
	x []int
}

func (g *d08Grid) get(x, y, defaultVal int) int {
	if y < 0 || y >= len(g.y) {
		return defaultVal
	}
	if x < 0 || x >= len(g.y[y].x) {
		return defaultVal
	}
	return g.y[y].x[x]
}

func (g *d08Grid) increase(x, y, delta int) {
	if y < 0 || y >= len(g.y) {
		return
	}
	if x < 0 || x >= len(g.y[y].x) {
		return
	}
	g.y[y].x[x] += delta
}

func (d *Day08) getGrid(input string) d08Grid {
	split := util.SplitLines(input)

	g := d08Grid{}
	for _, line := range split {
		l := d08Line{
			x: util.ParseInts(strings.Split(line, "")),
		}

		g.y = append(g.y, l)
	}
	return g
}

func (g *d08Grid) walkGrid(walker func(x, y, val int)) {
	for y := range g.y {
		for x := range g.y[y].x {
			walker(x, y, g.y[y].x[x])
		}
	}
}

func (d *Day08) SolveI(input string) any {
	mark := make(map[d08Coord]struct{}, 2_048)
	grid := d.getGrid(input)
	grid.walkGrid(func(x, y, val int) {
		// x to the left
		sightline := true
		for curX := x - 1; curX >= 0; curX-- {
			if grid.get(curX, y, 0) >= val {
				sightline = false
				break
			}
		}
		if sightline {
			mark[d08Coord{x, y}] = struct{}{}
			return
		}

		// x to the right
		sightline = true
		for curX := x + 1; curX < len(grid.y[y].x); curX++ {
			if grid.get(curX, y, 0) >= val {
				sightline = false
				break
			}
		}
		if sightline {
			mark[d08Coord{x, y}] = struct{}{}
			return
		}

		// y to the top
		sightline = true
		for curY := y - 1; curY >= 0; curY-- {
			if grid.get(x, curY, 0) >= val {
				sightline = false
				break
			}
		}
		if sightline {
			mark[d08Coord{x, y}] = struct{}{}
			return
		}

		// y to the bottom
		sightline = true
		for curY := y + 1; curY < len(grid.y); curY++ {
			if grid.get(x, curY, 0) >= val {
				sightline = false
				break
			}
		}
		if sightline {
			mark[d08Coord{x, y}] = struct{}{}
			return
		}
	})
	return len(mark)
}

func (d *Day08) SolveII(input string) any {
	scores := make([]int64, 0, 2_048)
	grid := d.getGrid(input)
	grid.walkGrid(func(x, y, val int) {
		score := int64(1)
		// Skip edges
		if x == 0 || x == len(grid.y[0].x)-1 {
			return
		}
		if y == 0 || y == len(grid.y)-1 {
			return
		}

		// x to the left
		for curX := x - 1; curX >= 0; curX-- {
			if grid.get(curX, y, 10) >= val || curX == 0 {
				score *= int64(x - curX)
				break
			}
		}

		// x to the right
		for curX := x + 1; curX < len(grid.y[y].x); curX++ {
			if grid.get(curX, y, 10) >= val || curX == len(grid.y[y].x)-1 {
				score *= int64(curX - x)
				break
			}
		}

		// y to the top
		for curY := y - 1; curY >= 0; curY-- {
			if grid.get(x, curY, 10) >= val || curY == 0 {
				score *= int64(y - curY)
				break
			}
		}

		// y to the bottom
		for curY := y + 1; curY < len(grid.y); curY++ {
			if grid.get(x, curY, 10) >= val || curY == len(grid.y)-1 {
				score *= int64(curY - y)
				break
			}
		}
		if score < 0 {
			log.Panicf("invalid score: %d.%d [%d]", x, y, score)
		}

		scores = append(scores, score)
	})

	return util.MaxSlice(scores)
}
