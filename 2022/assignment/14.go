package assignment

import (
	"fmt"
	"math"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day14 struct{}

type d14Coord struct {
	x, y int
}

type d14Grid struct {
	minX, minY, maxX, maxY int
	settled                map[d14Coord]uint8
}

const (
	d14Nothing = 0
	d14Rock    = 1
	d14Sand    = 2
)

func (g *d14Grid) print() {
	fmt.Println()
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			switch g.settled[d14Coord{x, y}] {
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
			for y := nxt.y; y <= c.y; y++ {
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

func (g *d14Grid) set(x, y int, val uint8) {
	g.settled[d14Coord{x, y}] = val
	g.minX = util.Min(x, g.minX)
	g.maxX = util.Max(x, g.maxX)
	g.minY = util.Min(y, g.minY)
	g.maxY = util.Max(y, g.maxY)
}

func (g *d14Grid) get(x, y int) uint8 {
	return g.settled[d14Coord{x, y}]
}

const (
	d14DropYLimit = 1300
)

func (g *d14Grid) dropSand(x, y int, print bool) (lastX, lastY int) {
	if g.get(x, y) != 0 {
		panic("sand blocked")
	}

	for {
		if print {
			fmt.Println("x", x, "y", y)
		}
		if y > d14DropYLimit {
			return x, y
		}
		if g.get(x, y+1) == d14Nothing {
			y++
			continue
		}
		if g.get(x-1, y+1) == d14Nothing {
			x--
			y++
			continue
		}
		if g.get(x+1, y+1) == d14Nothing {
			x++
			y++
			continue
		}
		// Comes to rest.
		g.set(x, y, d14Sand)
		return x, y
	}
}

func (d *Day14) makeGrid() *d14Grid {
	g := &d14Grid{
		settled: make(map[d14Coord]uint8, 1024),
		minX:    math.MaxInt,
		minY:    math.MaxInt,
	}
	return g
}

func (d *Day14) SolveI(input string) any {
	g := d.makeGrid()
	g.drawRocks(input)
	fmt.Println(len(g.settled))
	i := 0
	for {
		_, y := g.dropSand(500, 0, false)
		if y >= d14DropYLimit {
			g.dropSand(500, 0, false)
			break
		}
		i++
	}

	g.print()

	// > 462, 463
	// = 897'
	fmt.Println(len(g.settled))
	return i
}

func (d *Day14) SolveII(input string) any {
	return "Not Implemented Yet"
}
