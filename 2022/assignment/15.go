package assignment

import (
	"fmt"
	"github.com/vansante/adventofcode/2022/util"
	"math"
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

func (g *d15Grid) area(middle d15Coord, distance int, yLine int, walker func(coord d15Coord)) {
	if middle.y-distance > yLine || middle.y+distance < yLine {
		return
	}

	offset := 0
	for y := middle.y - distance; y <= middle.y+distance; y++ {
		if y == yLine {
			for x := middle.x - offset; x <= middle.x+offset; x++ {
				walker(d15Coord{x, y})
			}
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

func (d *Day15) findImpossible(input string, yLine int) *d15Grid {
	sensors := d.getSensors(input)
	grid := d.makeGrid()

	for i := range sensors {
		s := &sensors[i]
		grid.set(s.location, d15_Sensor)
		grid.set(s.closest, d15_Beacon)

		grid.area(s.location,
			util.ManhattanDistance(s.location.x, s.location.y, s.closest.x, s.closest.y),
			yLine, func(coord d15Coord) {
				if grid.get(coord) != d15_Empty {
					return
				}
				grid.set(coord, d15_NoBeacon)
			},
		)
	}
	return grid
}

func (d *Day15) SolveI(input string) any {
	const yLine = 200_000
	grid := d.findImpossible(input, yLine)

	sum := 0
	for x := grid.minX; x <= grid.maxX; x++ {
		if grid.get(d15Coord{x, yLine}) == d15_NoBeacon {
			sum++
		}
	}

	return sum
}

type d15Range struct {
	start, end int
}

func (d *Day15) findPossible(input string, min, max int) *d15Coord {
	sensors := d.getSensors(input)

	var result *d15Coord
	for y := min; y <= max; y++ {
		ranges := make([]d15Range, 0)
		for i := range sensors {
			s := sensors[i]

			distance := util.ManhattanDistance(s.location.x, s.location.y, s.closest.x, s.closest.y)
			yDistance := util.Abs(s.location.y - y)
			if yDistance > distance {
				continue // Sensor not in range
			}

			rng := d15Range{
				start: util.Max(min, s.location.x-(distance-yDistance)),
				end:   util.Min(max, s.location.x+(distance-yDistance)),
			}

			if len(ranges) == 0 {
				ranges = append(ranges, rng)
				continue
			}
			nwRanges := util.CopySlice(ranges)
			removed := 0
			for j := range ranges {
				r := ranges[j]
				if rng.start > r.end+1 || r.start > rng.end+1 {
					continue
				}

				// Merge the range:
				rng.start = util.Min(rng.start, r.start)
				rng.end = util.Max(rng.end, r.end)

				// Remove the old range:
				nwRanges = append(nwRanges[:j-removed], nwRanges[j-removed+1:]...)
				removed++
			}
			ranges = append(nwRanges, rng)
		}

		util.SliceReverse(ranges)

		if ranges[0].start == min && ranges[0].end == max {
			if len(ranges) != 1 {
				panic("invalid nonmerged ranges")
			}
			continue
		}

		if result != nil {
			panic("multiple result coordinates")
		}
		
		result = &d15Coord{ranges[0].end + 1, y}
	}

	return result
}

func (d *Day15) SolveII(input string) any {
	coord := d.findPossible(input, 0, 4_000_000)

	return coord.x*4_000_000 + coord.y
}
