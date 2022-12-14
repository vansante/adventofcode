package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day14 struct{}

type d14Coord struct {
	x, y int
}

type d14Grid struct {
	y []d14Line
}

const (
	d14Rock = 10
	d14Sand = 20
)

func (g *d14Grid) print() {
	fmt.Println()
	for y := range g.y {
		for x := range g.y[y].x {
			switch g.get(x, y, 0) {
			case d14Rock:
				print("#")
			case d14Sand:
				print("0")
			default:
				print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *d14Grid) drawRocks(input string) {
	lines := util.SplitLines(input)

	for _, line := range lines {
		coordStrs := strings.Split(line, " -> ")
		coords := make([]d14Coord, len(coordStrs))
		for i := range coordStrs {
			n, err := fmt.Sscanf(coordStrs[i], "%d,%d", &coords[i].x, &coords[i].y)
			util.CheckErr(err)
			if n != 2 {
				panic("invalid matches")
			}
		}
		g.drawRockLine(coords)
	}
}

func (g *d14Grid) drawRockLine(line []d14Coord) {
	for i, c := range line {
		g.set(c.x, c.y, d14Rock)
		if i == len(line)-1 {
			continue
		}

		nxt := line[i+1]
		switch {
		case c.x < nxt.x && c.y == nxt.y:
			for x := c.x; x <= nxt.x; x++ {
				g.set(x, c.y, d14Rock)
			}
		case c.x > nxt.x && c.y == nxt.y:
			for x := nxt.x; x <= c.x; x++ {
				g.set(x, c.y, d14Rock)
			}
		case c.y < nxt.y && c.x == nxt.x:
			for y := c.y; y <= nxt.y; y++ {
				g.set(c.x, y, d14Rock)
			}
		case c.y > nxt.y && c.x == nxt.x:
			for y := nxt.x; y <= c.x; y++ {
				g.set(c.x, y, d14Rock)
			}
		default:
			panic("invalid coords")
		}
	}
}

type d14Line struct {
	x []int
}

const d14XTranslate = -450

func (g *d14Grid) set(x, y, val int) {
	// Translate
	//x += d14XTranslate
	if y < 0 || y >= len(g.y) {
		panic("y out of bounds")
	}
	if x < 0 || x >= len(g.y[y].x) {
		panic("x out of bounds")
	}
	g.y[y].x[x] = val
}

func (g *d14Grid) get(x, y, defaultVal int) int {
	//x += d14XTranslate
	if y < 0 || y >= len(g.y) {
		return defaultVal
	}
	if x < 0 || x >= len(g.y[y].x) {
		return defaultVal
	}
	return g.y[y].x[x]
}

const (
	d14DropYLimit = 1000
)

func (g *d14Grid) dropSand(x, y int) (lastX, lastY int) {
	if g.get(x, y, 0) != 0 {
		panic("sand blocked")
	}

	for {
		if y > d14DropYLimit {
			return x, y
		}
		if g.get(x, y+1, 0) == 0 {
			y++
			continue
		}
		if g.get(x-1, y+1, 0) == 0 {
			x--
			y++
			continue
		}
		if g.get(x+1, y+1, 0) == 0 {
			x++
			y++
			continue
		}
		// Comes to rest.
		g.set(x, y, d14Sand)
		return x, y
	}
}

func (d *Day14) makeGrid(width, height int) *d14Grid {
	g := &d14Grid{
		y: make([]d14Line, height),
	}
	for y := range g.y {
		g.y[y].x = make([]int, width)
	}
	return g
}

func (d *Day14) SolveI(input string) any {
	g := d.makeGrid(600, 600)

	g.drawRocks(input)

	i := 0
	for {
		_, y := g.dropSand(500, 0)
		if y >= d14DropYLimit {
			break
		}
		i++
	}

	//g.print()

	// > 462
	return i
}

func (d *Day14) SolveII(input string) any {
	return "Not Implemented Yet"
}
