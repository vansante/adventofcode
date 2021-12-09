package assignment

import (
	"fmt"
	"math"
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

func (D *Day09) findLargestMultiplied(sizes []int64, n int) int64 {
	mult := int64(0)
	for i := 0; i < n; i++ {
		max := int64(math.MinInt64)
		idx := -1
		for j := range sizes {
			if sizes[j] > max {
				max = sizes[j]
				idx = j
			}
		}
		sizes = append(sizes[:idx], sizes[idx+1:]...)
		if i == 0 {
			mult = max
		} else {
			mult *= max
		}
	}
	return mult
}

func (d *Day09) SolveI(input string) int64 {
	m := d.getMap(input)
	return m.findRisk()
}

func (d *Day09) SolveII(input string) int64 {
	m := d.getMap(input)
	basins := m.findBasins()

	return d.findLargestMultiplied(basins, 3)
}
