package assignment

import (
	"fmt"
	"math"

	"github.com/vansante/adventofcode/2022/util"
)

type Day18 struct{}

const (
	d18Top uint8 = iota
	d18Bottom
	d18Front
	d18Back
	d18Left
	d18Right
)

var d18NeighborVectors = map[uint8]d18Coord{
	d18Top:    {0, 0, -1},
	d18Bottom: {0, 0, 1},
	d18Front:  {0, -1, 0},
	d18Back:   {0, 1, 0},
	d18Left:   {-1, 0, 0},
	d18Right:  {1, 0, 0},
}

type d18Coord struct {
	x, y, z int
}

func (c d18Coord) add(other d18Coord) d18Coord {
	return d18Coord{
		x: c.x + other.x,
		y: c.y + other.y,
		z: c.z + other.z,
	}
}

type d18Cube struct {
	d18Coord
	sides [6]uint
}

func (d *Day18) getCubes(input string) []d18Cube {
	lines := util.SplitLines(input)

	c := make([]d18Cube, len(lines))
	for i, line := range lines {
		n, err := fmt.Sscanf(line, "%d,%d,%d", &c[i].x, &c[i].y, &c[i].z)
		util.CheckErr(err)
		if n != 3 {
			panic("invalid matches")
		}
	}
	return c
}

func (d *Day18) placeCubes(cubes []d18Cube) map[d18Coord]*d18Cube {
	mp := make(map[d18Coord]*d18Cube, 10_000)
	for i := range cubes {
		mp[cubes[i].d18Coord] = &cubes[i]
	}
	return mp
}

func (d *Day18) markSides(cubes map[d18Coord]*d18Cube) {
	for coord, cube := range cubes {
		for i, vec := range d18NeighborVectors {
			ngbCoord := coord.add(vec)
			ngb := cubes[ngbCoord]
			if ngb == nil {
				continue
			}

			cube.sides[i]++
		}
	}
}

func (d *Day18) markVoids(cubes map[d18Coord]*d18Cube, reachable map[d18Coord]struct{}) {
	for coord, cube := range cubes {
		for i, vec := range d18NeighborVectors {
			ngbCoord := coord.add(vec)
			if _, ok := reachable[ngbCoord]; ok {
				continue
			}

			cube.sides[i]++
		}
	}
}

func (d *Day18) countSides(mp map[d18Coord]*d18Cube) int64 {
	var sum int64

	for _, c := range mp {
		for i := range c.sides {
			if c.sides[i] == 0 {
				sum++
			}
		}
	}
	return sum
}

func (d *Day18) SolveI(input string) any {
	c := d.placeCubes(d.getCubes(input))
	d.markSides(c)

	return d.countSides(c)
}

func (d *Day18) findBorderCoords(cubes []d18Cube) (min, max d18Coord) {
	min.x = math.MaxInt
	min.y = math.MaxInt
	min.z = math.MaxInt
	for _, cube := range cubes {
		max.x = util.Max(max.x, cube.x)
		max.y = util.Max(max.y, cube.y)
		max.z = util.Max(max.z, cube.z)

		min.x = util.Min(min.x, cube.x)
		min.y = util.Min(min.y, cube.y)
		min.z = util.Min(min.z, cube.z)
	}
	// Add margins
	max.x++
	max.y++
	max.z++

	min.x--
	min.y--
	min.z--
	return min, max
}

func (d *Day18) reachableWalker(current, min, max d18Coord, coords map[d18Coord]*d18Cube, reachable *map[d18Coord]struct{}) {
	r := *(reachable)
	for _, vec := range d18NeighborVectors {
		c := current.add(vec)

		switch {
		case c.x < min.x || c.y < min.y || c.z < min.z:
			continue // Skip out of borders
		case c.x > max.x || c.y > max.y || c.z > max.z:
			continue // Skip out of borders
		case coords[c] != nil:
			continue // Skip places with cube
		}

		if _, ok := r[c]; ok {
			continue // Skip already reached places
		}

		// Mark as reachable
		r[c] = struct{}{}
		// Do neighbours
		d.reachableWalker(c, min, max, coords, reachable)
	}
}

func (d *Day18) findReachableCoords(coords map[d18Coord]*d18Cube, min, max d18Coord) map[d18Coord]struct{} {
	reachable := make(map[d18Coord]struct{}, 100_000)
	d.reachableWalker(min, min, max, coords, &reachable)
	return reachable
}

func (d *Day18) SolveII(input string) any {
	cubes := d.getCubes(input)
	coords := d.placeCubes(cubes)
	min, max := d.findBorderCoords(cubes)
	reachable := d.findReachableCoords(coords, min, max)

	d.markSides(coords)
	d.markVoids(coords, reachable)

	return d.countSides(coords)
}
