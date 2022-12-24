package assignment

import (
	"github.com/vansante/adventofcode/2022/util"
)

type Day24 struct{}

const maxMinutes = 5_000

var (
	d24Up    = d24Coord{0, -1}
	d24Right = d24Coord{1, 0}
	d24Down  = d24Coord{0, 1}
	d24Left  = d24Coord{-1, 0}
)

var (
	d24Vectors = []d24Coord{
		d24Up,
		d24Right,
		d24Down,
		d24Left,
		{},
	}
)

type d24Coord struct {
	x, y int
}

func (c d24Coord) add(other d24Coord) d24Coord {
	return d24Coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

func (c d24Coord) equals(other d24Coord) bool {
	return c.x == other.x && c.y == other.y
}

type d24Blizzard struct {
	start  d24Coord
	vector d24Coord
}

type d24Grid struct {
	width, height int
	blz           []d24Blizzard
}

func (d *Day24) getGrid(input string) *d24Grid {
	g := &d24Grid{
		blz: make([]d24Blizzard, 0, 4096),
	}

	lines := util.SplitLines(input)
	g.height = len(lines) - 2
	for y, line := range lines {
		g.width = len(line) - 2
		for x, char := range line {
			var vector d24Coord
			switch char {
			case '.', '#':
				continue
			case '^':
				vector = d24Up
			case '>':
				vector = d24Right
			case 'v':
				vector = d24Down
			case '<':
				vector = d24Left
			default:
				panic("invalid char")
			}
			g.blz = append(g.blz, d24Blizzard{
				start:  d24Coord{x - 1, y - 1},
				vector: vector,
			})
		}
	}
	return g
}

func (g *d24Grid) hasBlizzard(c d24Coord, minute int) bool {
	for _, b := range g.blz {
		// Skip blizzards that cannot cross path
		if b.start.x != c.x && b.start.y != c.y {
			continue
		}

		x := (b.start.x + g.width + (b.vector.x * minute % g.width)) % g.width
		if x != c.x {
			continue
		}
		y := (b.start.y + g.height + (b.vector.y * minute % g.height)) % g.height
		if y != c.y {
			continue
		}
		return true
	}
	return false
}

func (g *d24Grid) clearMap(minute int) [][]bool {
	mp := make([][]bool, g.height)

	for y := 0; y < g.height; y++ {
		mp[y] = make([]bool, g.width)
		for x := 0; x < g.width; x++ {
			c := d24Coord{x, y}
			if g.hasBlizzard(c, minute) {
				continue
			}
			mp[y][x] = true
		}
	}
	return mp
}

func (g *d24Grid) walk(minute int, start, end d24Coord) int {
	minute++

	queue := make(map[d24Coord]struct{}, 256)
	queue[start] = struct{}{}
	for minute < maxMinutes {
		minute++
		clearMap := g.clearMap(minute)
		nxtQueue := make(map[d24Coord]struct{}, 256)
		for c := range queue {
			if c.equals(end) {
				return minute - 1
			}

			for _, vec := range d24Vectors {
				ngb := c.add(vec)

				// Start and end are always accessible
				if ngb.equals(start) || ngb.equals(end) {
					nxtQueue[ngb] = struct{}{}
					continue
				}

				if !g.inBounds(ngb) {
					//if !g.inBounds(ngb) || g.hasBlizzard(ngb, minute) {
					continue
				}

				// Check for blizzards
				if !clearMap[ngb.y][ngb.x] {
					continue
				}

				nxtQueue[ngb] = struct{}{}
			}
		}
		queue = nxtQueue
	}
	panic("route not found")
}

func (g *d24Grid) inBounds(c d24Coord) bool {
	return c.x >= 0 && c.x < g.width && c.y >= 0 && c.y < g.height
}

func (g *d24Grid) start() d24Coord {
	return d24Coord{x: 0, y: -1}
}

func (g *d24Grid) end() d24Coord {
	return d24Coord{x: g.width - 1, y: g.height}
}

func (d *Day24) SolveI(input string) any {
	g := d.getGrid(input)
	return g.walk(0, g.start(), g.end())
}

func (d *Day24) SolveII(input string) any {
	g := d.getGrid(input)

	start := g.start()
	end := g.end()
	minutes := g.walk(0, start, end)
	minutes = g.walk(minutes, end, start)
	minutes = g.walk(minutes, start, end)
	return minutes
}
