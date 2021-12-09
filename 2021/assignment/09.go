package assignment

import (
	"fmt"
	"sort"
	"strings"
)

type Day09 struct{}

type d09Vector struct {
	x, y int
}

var (
	d09Vectors = []d09Vector{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

type d09map struct {
	points [][]int
}

func (m *d09map) get(x, y int, defaultVal int) int {
	if y < 0 || y >= len(m.points) {
		return defaultVal
	}
	if x < 0 || x >= len(m.points[y]) {
		return defaultVal
	}
	return m.points[y][x]
}

func (m *d09map) findRisk() int64 {
	sum := int64(0)
	for y := range m.points {
		for x := range m.points[y] {
			lower := true
			for _, v := range d09Vectors {
				risk := m.get(x+v.x, y+v.y, 11)
				lower = lower && risk > m.points[y][x]
			}
			if lower {
				sum += int64(m.points[y][x] + 1)
			}
		}
	}
	return sum
}

func (m *d09map) findBasins() []int64 {
	basins := make(map[string]int64)

	for y := range m.points {
		for x := range m.points[y] {
			if m.points[y][x] >= 9 {
				continue
			}

			curX := x
			curY := y
			var lowerX, lowerY int
			for {
				lowest := 11

				for _, v := range d09Vectors {
					neighbour := m.get(curX+v.x, curY+v.y, 11)

					if neighbour < lowest {
						lowest = neighbour
						lowerX = curX + v.x
						lowerY = curY + v.y
					}
				}

				if lowest >= m.points[curY][curX] {
					basins[fmt.Sprintf("%d_%d", curY, curX)]++
					break
				}

				curX = lowerX
				curY = lowerY
			}
		}
	}

	res := make([]int64, len(basins))
	i := 0
	for k := range basins {
		res[i] = basins[k]
		i++
	}
	return res
}

func (d *Day09) getMap(input string) d09map {
	lines := SplitLines(input)
	m := d09map{
		points: make([][]int, len(lines)),
	}
	for i := range lines {
		points := strings.Split(lines[i], "")
		m.points[i] = MakeInts(points)
	}
	return m
}

func (d *Day09) SolveI(input string) int64 {
	m := d.getMap(input)
	return m.findRisk()
}

func (d *Day09) SolveII(input string) int64 {
	m := d.getMap(input)
	basins := m.findBasins()

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}
