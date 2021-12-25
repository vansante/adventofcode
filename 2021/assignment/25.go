package assignment

import (
	"fmt"
	"strings"
)

type Day25 struct{}

type d25Direction int8

const (
	d25East d25Direction = iota + 1
	d25South
)

type d25Cucumber struct {
	dir d25Direction
}

func (c *d25Cucumber) equals(other *d25Cucumber) bool {
	if c == nil && other == nil {
		return true
	}
	if c == nil || other == nil {
		return false
	}
	return c.dir == other.dir
}

type d25Grid struct {
	y []d25Line
}

type d25Line struct {
	x []*d25Cucumber
}

func (g *d25Grid) get(x, y int) *d25Cucumber {
	if y >= len(g.y) {
		y = 0
	}
	if x >= len(g.y[y].x) {
		x = 0
	}
	return g.y[y].x[x]
}

func (g *d25Grid) set(x, y int, c *d25Cucumber) {
	if y >= len(g.y) {
		y = 0
	}
	if x >= len(g.y[y].x) {
		x = 0
	}
	g.y[y].x[x] = c
}

func (g *d25Grid) walkGrid(walker func(x, y int, c *d25Cucumber)) {
	for y := range g.y {
		for x := range g.y[y].x {
			walker(x, y, g.y[y].x[x])
		}
	}
}

func (g *d25Grid) moveStep() *d25Grid {
	nw := &d25Grid{y: make([]d25Line, len(g.y))}
	for y := range nw.y {
		nw.y[y] = d25Line{x: make([]*d25Cucumber, len(g.y[y].x))}
	}

	g.walkGrid(func(x, y int, c *d25Cucumber) {
		switch {
		case c == nil:
			return
		case c.dir == d25East:
			if g.get(x+1, y) == nil {
				nw.set(x+1, y, c)
			} else {
				nw.set(x, y, c)
			}
		}
	})
	g.walkGrid(func(x, y int, c *d25Cucumber) {
		switch {
		case c == nil:
			return
		case c.dir == d25South:
			if nw.get(x, y+1) == nil && (g.get(x, y+1) == nil || g.get(x, y+1).dir == d25East) {
				nw.set(x, y+1, c)
			} else {
				nw.set(x, y, c)
			}
		}
	})
	return nw
}

func (g *d25Grid) equals(other *d25Grid) bool {
	equal := true
	g.walkGrid(func(x, y int, c *d25Cucumber) {
		equal = equal && c.equals(other.get(x, y))
	})
	return equal
}

func (g *d25Grid) print() {
	for y := range g.y {
		for _, c := range g.y[y].x {
			if c == nil {
				print(".")
			} else if c.dir == d25East {
				print(">")
			} else if c.dir == d25South {
				print("v")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day25) getGrid(input string) *d25Grid {
	split := SplitLines(input)

	g := &d25Grid{}
	for _, line := range split {
		chars := strings.Split(line, "")
		l := d25Line{
			x: make([]*d25Cucumber, len(chars)),
		}
		for i, char := range chars {
			switch char {
			case ".":
			case ">":
				l.x[i] = &d25Cucumber{dir: d25East}
			case "v":
				l.x[i] = &d25Cucumber{dir: d25South}
			default:
				panic("unknown character")
			}
		}

		g.y = append(g.y, l)
	}
	return g
}

func (d *Day25) SolveI(input string) int64 {
	g := d.getGrid(input)
	g.print()

	var i int64
	for i = 1; i < 10000; i++ {
		nw := g.moveStep()

		if g.equals(nw) {
			break
		}
		g = nw
	}

	g.print()

	return i
}

func (d *Day25) SolveII(input string) int64 {
	return 0
}
