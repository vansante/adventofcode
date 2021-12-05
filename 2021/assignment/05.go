package assignment

import (
	"fmt"
	"log"
	"math"
)

type Day05 struct{}

const d05MaxCoord = 1000

type d05Coord struct {
	x, y int64
}

type d05Line struct {
	a, b d05Coord
}

func (l *d05Line) isStraight() bool {
	return l.a.x == l.b.x || l.a.y == l.b.y
}

func (l *d05Line) isDiagonal() bool {
	return int(math.Abs(float64(l.a.x-l.b.x))) == int(math.Abs(float64(l.a.y-l.b.y)))
}

func (l *d05Line) slopeX() int64 {
	if l.a.x > l.b.x {
		return -1
	} else if l.a.x < l.b.x {
		return 1
	}
	return 0
}

func (l *d05Line) slopeY() int64 {
	if l.a.y > l.b.y {
		return -1
	} else if l.a.y < l.b.y {
		return 1
	}
	return 0
}

type d05Grid struct {
	coords [][]int
}

func (d *Day05) newGrid() d05Grid {
	g := d05Grid{
		coords: make([][]int, d05MaxCoord),
	}
	for i := range g.coords {
		g.coords[i] = make([]int, d05MaxCoord)
	}
	return g
}

func (g *d05Grid) markLine(l d05Line) {
	for x, y := l.a.x, l.a.y; ; {
		g.coords[y][x]++

		x += l.slopeX()
		y += l.slopeY()

		if x == l.b.x && y == l.b.y {
			g.coords[y][x]++
			return
		}
	}
}

func (g *d05Grid) print() {
	fmt.Println()
	for y := range g.coords {
		fmt.Println(g.coords[y])
	}
	fmt.Println()
}

func (g *d05Grid) countCrossings() int64 {
	var count int64
	for y := range g.coords {
		for x := range g.coords[y] {
			if g.coords[y][x] > 1 {
				count++
			}
		}
	}
	return count
}

func (d *Day05) GetLines(input string) []d05Line {
	split := SplitLines(input)

	lines := make([]d05Line, 0, len(split))
	for _, line := range split {
		l := d05Line{}
		n, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &l.a.x, &l.a.y, &l.b.x, &l.b.y)
		if err != nil || n != 4 {
			log.Panicf("[%s] error parsing line: %v | %d", line, err, n)
		}

		lines = append(lines, l)
	}
	return lines
}

func (d *Day05) SolveI(input string) int64 {
	ls := d.GetLines(input)

	g := d.newGrid()
	for i := range ls {
		l := ls[i]
		if !l.isStraight() {
			continue
		}
		g.markLine(l)
	}
	//g.print()

	return g.countCrossings()
}

func (d *Day05) SolveII(input string) int64 {
	ls := d.GetLines(input)

	g := d.newGrid()
	for i := range ls {
		l := ls[i]
		if !l.isStraight() && !l.isDiagonal() {
			continue
		}
		g.markLine(l)
	}
	//g.print()

	return g.countCrossings()
}
