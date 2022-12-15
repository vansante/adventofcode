package assignment

import (
	"fmt"
	"math"

	"github.com/vansante/adventofcode/2022/util"
)

type Day15 struct{}

type d15Coord struct {
	x, y int
}

type d15Sensor struct {
	location d15Coord
	closest  d15Coord
}

const (
	d15_Empty uint8 = iota
	d15_Sensor
	d15_Beacon
	d15_NoBeacon
)

type d15Grid struct {
	minX, minY, maxX, maxY int
	coords                 map[d15Coord]uint8
}

func (d *Day15) makeGrid() *d15Grid {
	g := &d15Grid{
		coords: make(map[d15Coord]uint8, 100_000),
		minX:   math.MaxInt,
		minY:   math.MaxInt,
	}
	return g
}

func (g *d15Grid) set(c d15Coord, val uint8) {
	g.coords[c] = val
	g.minX = util.Min(c.x, g.minX)
	g.maxX = util.Max(c.x, g.maxX)
	g.minY = util.Min(c.y, g.minY)
	g.maxY = util.Max(c.y, g.maxY)
}

func (g *d15Grid) get(c d15Coord) uint8 {
	return g.coords[c]
}

func (g *d15Grid) area(middle d15Coord, distance int, walker func(coord d15Coord)) {
	offset := 0
	for y := middle.y - distance; y <= middle.y+distance; y++ {
		for x := middle.x - offset; x <= middle.x+offset; x++ {
			walker(d15Coord{x, y})
		}
		if y < middle.y {
			offset++
		} else {
			offset--
		}
	}
}

func (g *d15Grid) print() {
	fmt.Println()
	for y := g.minY; y <= g.maxY; y++ {
		fmt.Printf("%03d ", y)
		for x := g.minX; x <= g.maxX; x++ {
			switch g.coords[d15Coord{x, y}] {
			case d15_Beacon:
				print("B")
			case d15_Sensor:
				print("S")
			case d15_NoBeacon:
				print("-")
			default:
				print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day15) getSensors(input string) []d15Sensor {
	lines := util.SplitLines(input)

	sensors := make([]d15Sensor, len(lines))
	for i, line := range lines {
		s := &sensors[i]
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.location.x, &s.location.y,
			&s.closest.x, &s.closest.y,
		)
		util.CheckErr(err)
		if n != 4 {
			panic("invalid matches")
		}
	}
	return sensors
}

func (d *Day15) SolveI(input string) any {
	sensors := d.getSensors(input)
	grid := d.makeGrid()

	//const yLine = 10
	const yLine = 2_000_000

	for i := range sensors {
		s := &sensors[i]
		grid.set(s.location, d15_Sensor)
		grid.set(s.closest, d15_Beacon)

		grid.area(s.location,
			util.ManhattanDistance(s.location.x, s.location.y, s.closest.x, s.closest.y),
			func(coord d15Coord) {
				if coord.y != yLine {
					return
				}
				if grid.get(coord) != d15_Empty {
					return
				}
				grid.set(coord, d15_NoBeacon)
			},
		)

	}
	//grid.print()
	fmt.Println(grid.minX, grid.maxX)
	fmt.Println(grid.minY, grid.maxY)

	sum := 0
	for x := grid.minX; x <= grid.maxX; x++ {
		if grid.get(d15Coord{x, yLine}) == d15_NoBeacon {
			sum++
		}
	}

	return sum
}

func (d *Day15) SolveII(input string) any {
	return "Not Implemented Yet"
}
