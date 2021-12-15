package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Day15 struct{}

type d15Coord struct {
	x, y int
}

func (c d15Coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c d15Coord) id() int64 {
	return int64(c.y)*10_000_000 + int64(c.x)
}

func (c d15Coord) add(c2 d15Coord) d15Coord {
	return d15Coord{c.x + c2.x, c.y + c2.y}
}

func (c d15Coord) equals(c2 d15Coord) bool {
	return c.x == c2.x && c.y == c2.y
}

var (
	d15Vectors = []d15Coord{
		{0, -1}, // top
		{1, 0},  // right
		{0, 1},  // bottom
		{-1, 0}, // left
	}
)

type d15Grid struct {
	y []d15Line
}

type d15Line struct {
	x []int
}

func (g *d15Grid) get(x, y, wrapCount int, defaultVal int) int {
	if wrapCount <= 0 {
		panic("invalid count")
	}
	lenY := len(g.y)
	if y < 0 || y >= lenY*wrapCount {
		return defaultVal
	}
	lenX := len(g.y[0].x)
	if x < 0 || x >= lenX*wrapCount {
		return defaultVal
	}

	penalty := 0
	nwY := y
	if y >= lenY {
		penalty += y / lenY
		nwY = y % lenY
	}
	nwX := x
	if x >= lenX {
		penalty += x / lenX
		nwX = x % lenX
	}

	val := (g.y[nwY].x[nwX] + penalty) % 9
	if val == 0 {
		return 9
	}
	return val
}

func (g *d15Grid) print() {
	for y := range g.y {
		for x := range g.y[y].x {
			print(g.y[y].x[x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day15) getGrid(input string) d15Grid {
	split := SplitLines(input)

	g := d15Grid{}
	for _, line := range split {
		l := d15Line{
			x: MakeInts(strings.Split(line, "")),
		}

		g.y = append(g.y, l)
	}
	return g
}

func (g *d15Grid) walkGrid(walker func(x, y, val int)) {
	for y := range g.y {
		for x := range g.y[y].x {
			walker(x, y, g.y[y].x[x])
		}
	}
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (g *d15Grid) dijkstra(start, end d15Coord, wrapCount int) int64 {
	visited := make(map[int64]struct{})
	dist := make(map[int64]int64)
	prev := make(map[int64]d15Coord)
	queue := make([]d15Coord, 1)

	const MaxDist = 100_000_000_000

	for y := 0; y < len(g.y)*wrapCount; y++ {
		for x := 0; x < len(g.y[0].x)*wrapCount; x++ {
			dist[d15Coord{x, y}.id()] = MaxDist
		}
	}

	dist[start.id()] = 0
	queue[0] = start

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return dist[queue[i].id()] > dist[queue[j].id()]
		})
		var cur d15Coord
		cur, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if cur.equals(end) {
			break
		}
		visited[cur.id()] = struct{}{}

		for _, v := range d15Vectors {
			ngb := cur.add(v)
			if _, ok := visited[ngb.id()]; ok {
				continue
			}

			cost := int64(g.get(ngb.x, ngb.y, wrapCount, math.MaxInt))
			if cost == math.MaxInt { // skip coords out of map
				continue
			}
			shortestDist := dist[cur.id()] + cost
			currentDist := dist[ngb.id()]
			if shortestDist < currentDist {
				queue = append(queue, ngb)
				dist[ngb.id()] = shortestDist
				prev[ngb.id()] = cur
			}
		}
	}

	return dist[end.id()]
}

func (d *Day15) SolveI(input string) int64 {
	g := d.getGrid(input)
	start := d15Coord{0, 0}
	end := d15Coord{len(g.y) - 1, len(g.y[len(g.y)-1].x) - 1}

	return g.dijkstra(start, end, 1)
}

func (d *Day15) SolveII(input string) int64 {
	g := d.getGrid(input)
	start := d15Coord{0, 0}
	end := d15Coord{len(g.y)*5 - 1, len(g.y[len(g.y)-1].x)*5 - 1}

	return g.dijkstra(start, end, 5)
}
