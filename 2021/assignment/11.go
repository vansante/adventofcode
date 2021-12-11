package assignment

import (
	"fmt"
	"strings"
)

type Day11 struct{}

const (
	d11FlashLevel = 9
)

type d11Vector struct {
	x, y int
}

var (
	d11Vectors = []d09Vector{
		{0, -1},  // top
		{1, -1},  // top right
		{1, 0},   // right
		{1, 1},   // bottom right
		{0, 1},   // bottom
		{-1, 1},  // bottom left
		{-1, 0},  // left
		{-1, -1}, // top left
	}
)

type d11Grid struct {
	y []d11Line
}

type d11Line struct {
	x []int
}

func (g *d11Grid) get(x, y, defaultVal int) int {
	if y < 0 || y >= len(g.y) {
		return defaultVal
	}
	if x < 0 || x >= len(g.y[y].x) {
		return defaultVal
	}
	return g.y[y].x[x]
}

func (g *d11Grid) increase(x, y, delta int) {
	if y < 0 || y >= len(g.y) {
		return
	}
	if x < 0 || x >= len(g.y[y].x) {
		return
	}
	g.y[y].x[x] += delta
}

func (d *Day11) getGrid(input string) d11Grid {
	split := SplitLines(input)

	g := d11Grid{}
	for _, line := range split {
		l := d11Line{
			x: MakeInts(strings.Split(line, "")),
		}

		g.y = append(g.y, l)
	}
	return g
}

func (g *d11Grid) walkGrid(walker func(x, y, val int)) {
	for y := range g.y {
		for x := range g.y[y].x {
			walker(x, y, g.y[y].x[x])
		}
	}
}

func (g *d11Grid) copy() *d11Grid {
	nw := &d11Grid{y: make([]d11Line, len(g.y))}
	for y := range g.y {
		nw.y = make([]d11Line, len(g.y[y].x))
		for x := range g.y[y].x {
			nw.y[y].x[x] = g.y[y].x[x]
		}
	}
	return nw
}

func (g *d11Grid) nextStep() int64 {
	g.walkGrid(func(x, y, val int) {
		g.increase(x, y, 1)
	})

	flashes := int64(0)
	nwFlashes := int64(1)
	hasFlashed := make(map[string]struct{})

	for nwFlashes > 0 {
		nwFlashes = 0
		g.walkGrid(func(x, y, val int) {
			if val <= d11FlashLevel {
				return
			}
			coord := fmt.Sprintf("%d,%d", x, y)
			if _, ok := hasFlashed[coord]; ok {
				return
			}
			nwFlashes++
			hasFlashed[coord] = struct{}{}

			for _, v := range d11Vectors {
				g.increase(x+v.x, y+v.y, 1)
			}
		})
		flashes += nwFlashes
	}
	g.walkGrid(func(x, y, val int) {
		if val > d11FlashLevel {
			g.y[y].x[x] = 0
		}
	})
	return flashes
}

func (g *d11Grid) steps(n int) int64 {
	total := int64(0)
	for i := 0; i < n; i++ {
		total += g.nextStep()
	}
	return total
}

func (g *d11Grid) findAllFlash() int64 {
	step := int64(0)
	for {
		total := g.nextStep()
		step++
		if total == int64(len(g.y)*len(g.y[0].x)) {
			return step
		}
	}
}

func (g *d11Grid) print() {
	for y := range g.y {
		for x := range g.y[y].x {
			print(g.y[y].x[x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day11) SolveI(input string) int64 {
	g := d.getGrid(input)

	return g.steps(100)
}

func (d *Day11) SolveII(input string) int64 {
	g := d.getGrid(input)

	return g.findAllFlash()
}
