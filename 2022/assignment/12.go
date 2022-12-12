package assignment

import (
	"fmt"
	"math"
	"sort"

	"github.com/vansante/adventofcode/2022/util"
)

type Day12 struct{}

type d12Coord struct {
	x, y int
}

func (c d12Coord) equals(other d12Coord) bool {
	return c.x == other.x && c.y == other.y
}

func (c d12Coord) add(other d12Coord) d12Coord {
	return d12Coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

var (
	d12Vectors = []d12Coord{
		{0, -1}, // top
		{1, 0},  // right
		{0, 1},  // bottom
		{-1, 0}, // left
	}
)

type d12Grid struct {
	y []d12Line
}

type d12Line struct {
	x []int
}

func (g *d12Grid) get(x, y, defaultVal int) int {
	lenY := len(g.y)
	if y < 0 || y >= lenY {
		return defaultVal
	}
	lenX := len(g.y[0].x)
	if x < 0 || x >= lenX {
		return defaultVal
	}
	return g.y[y].x[x]
}

func (g *d12Grid) print() {
	for y := range g.y {
		for x := range g.y[y].x {
			print(g.y[y].x[x])
		}
		fmt.Println()
	}
	fmt.Println()
}

const (
	startVal         = 'S'
	endVal           = 'E'
	possibleStartVal = 'a'
)

func (d *Day12) getGrid(input string) (g d12Grid, start, end d12Coord) {
	split := util.SplitLines(input)

	for y, line := range split {
		l := d12Line{
			x: make([]int, len(line)),
		}
		for x, char := range line {
			switch char {
			case startVal:
				start = d12Coord{x, y}
				char = 'a'
			case endVal:
				end = d12Coord{x, y}
				char = 'z'
			}

			l.x[x] = int(char)
		}

		g.y = append(g.y, l)
	}
	return g, start, end
}

func (g *d12Grid) walkGrid(walker func(x, y, val int)) {
	for y := range g.y {
		for x := range g.y[y].x {
			walker(x, y, g.y[y].x[x])
		}
	}
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (g *d12Grid) distance(start, end d12Coord) int64 {
	visited := make(map[d12Coord]struct{}, 256)
	dist := make(map[d12Coord]int64, 256)
	prev := make(map[d12Coord]d12Coord, 256)
	queue := make([]d12Coord, 1, 1_024)

	const MaxDist = math.MaxInt

	for y := 0; y < len(g.y); y++ {
		for x := 0; x < len(g.y[0].x); x++ {
			dist[d12Coord{x, y}] = MaxDist
		}
	}

	dist[start] = 0
	queue[0] = start

	for len(queue) > 0 {
		var cur d12Coord
		cur, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if cur.equals(end) {
			break
		}
		visited[cur] = struct{}{}

		for _, v := range d12Vectors {
			neighbours := cur.add(v)
			if _, ok := visited[neighbours]; ok {
				continue
			}

			thisHeight := g.get(cur.x, cur.y, math.MaxInt)
			ngbHeight := g.get(neighbours.x, neighbours.y, math.MaxInt)
			if ngbHeight > thisHeight+1 { // skip too high coords and coords out of map
				continue
			}
			shortestDist := dist[cur] + 1 // Cost is always 1
			currentDist := dist[neighbours]
			if shortestDist >= currentDist {
				continue
			}

			// Insert neighbour into priority queue, lowest distance last
			idx := sort.Search(len(queue), func(i int) bool {
				return dist[queue[i]] <= shortestDist
			})
			queue = append(queue[:idx], append([]d12Coord{neighbours}, queue[idx:]...)...)

			dist[neighbours] = shortestDist
			prev[neighbours] = cur
		}
	}

	return dist[end]
}

func (g *d12Grid) findStarts() []d12Coord {
	c := make([]d12Coord, 0, 128)
	g.walkGrid(func(x, y, val int) {
		if val == possibleStartVal {
			c = append(c, d12Coord{x, y})
		}
	})
	return c
}

func (d *Day12) SolveI(input string) any {
	grid, start, end := d.getGrid(input)

	return grid.distance(start, end)
}

func (d *Day12) SolveII(input string) any {
	grid, _, end := d.getGrid(input)
	starts := grid.findStarts()

	min := int64(math.MaxInt64)
	for _, start := range starts {
		dist := grid.distance(start, end)
		if dist < min {
			min = dist
		}
	}
	return min
}
