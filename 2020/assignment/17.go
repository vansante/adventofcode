package assignment

import (
	"fmt"
	"strings"
)

type Day17 struct {
	neighborVectors  []d17Vector
	neighborHVectors []d17HVector
}

type d17Coord bool
type d17Plane [][]d17Coord
type d17Grid []d17Plane
type d17HGrid []d17Grid

func (d *Day17) retrieveGridPlane(in string) d17Plane {
	split := strings.Split(in, "\n")

	var input d17Plane
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		var dim []d17Coord
		for i := range chars {
			switch chars[i] {
			case ".":
				dim = append(dim, false)
			case "#":
				dim = append(dim, true)
			default:
				panic(chars[i])
			}
		}

		input = append(input, dim)
	}
	return input
}

func (p d17Plane) expand(expandSize int) d17Plane {
	newP := make(d17Plane, expandSize)
	for y := range newP {
		newP[y] = make([]d17Coord, expandSize)
	}

	for y := range p {
		newY := (expandSize/2 - (len(p) / 2)) + y
		for x := range p[y] {
			newX := (expandSize/2 - (len(p[y]) / 2)) + x
			newP[newY][newX] = p[y][x]
		}
	}
	return newP
}

func (d *Day17) makeGrid(expandPlanes int, startPlane d17Plane) d17Grid {
	grid := make(d17Grid, 2*expandPlanes+1)
	for z := range grid {
		grid[z] = make(d17Plane, len(startPlane))
		for y := range grid[z] {
			grid[z][y] = make([]d17Coord, len(startPlane[y]))
		}
	}
	grid[expandPlanes] = startPlane
	return grid
}

func (d *Day17) makeHGrid(expandPlanes int, startPlane d17Plane) d17HGrid {
	nwG := make(d17HGrid, 2*expandPlanes+1)
	for w := range nwG {
		nwG[w] = make(d17Grid, len(startPlane))
		for z := range nwG[w] {
			nwG[w][z] = make(d17Plane, len(startPlane))
			for y := range nwG[w][z] {
				nwG[w][z][y] = make([]d17Coord, len(startPlane[y]))
			}
		}
	}
	nwG[expandPlanes][expandPlanes] = startPlane
	return nwG
}

type d17Vector struct {
	x, y, z int
}

type d17HVector struct {
	x, y, z, w int
}

func (d *Day17) init() {
	d.neighborVectors = make([]d17Vector, 0, 3*3*3)
	d.neighborHVectors = make([]d17HVector, 0, 3*3*3*3)
	// Generate some vectors (tedious)
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x != 0 || y != 0 || z != 0 {
					d.neighborVectors = append(d.neighborVectors, d17Vector{
						x: x,
						y: y,
						z: z,
					})
				}
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					d.neighborHVectors = append(d.neighborHVectors, d17HVector{
						x: x,
						y: y,
						z: z,
						w: w,
					})
				}
			}
		}
	}
}

func (g d17Grid) get(x, y, z int, defaultVal d17Coord) d17Coord {
	if z < 0 || z >= len(g) {
		return defaultVal
	}
	if y < 0 || y >= len(g[z]) {
		return defaultVal
	}
	if x < 0 || x >= len(g[z][y]) {
		return defaultVal
	}
	return g[z][y][x]
}

func (g d17Grid) copy() d17Grid {
	newGrid := make(d17Grid, len(g))
	for z := range g {
		newGrid[z] = make(d17Plane, len(g[z]))
		for y := range g[z] {
			newGrid[z][y] = make([]d17Coord, len(g[z][y]))
			for x := range g[z][y] {
				newGrid[z][y][x] = g[z][y][x]
			}
		}
	}
	return newGrid
}

func (d *Day17) simulateGrid(g d17Grid) d17Grid {
	newG := g.copy()
	for z := range g {
		for y := range g[z] {
			for x := range g[z][y] {
				count := 0
				for i := range d.neighborVectors {
					v := d.neighborVectors[i]
					if g.get(x+v.x, y+v.y, z+v.z, false) {
						count++
					}
				}
				newG[z][y][x] = (g[z][y][x] && count >= 2 && count <= 3) ||
					(!g[z][y][x] && count == 3)
			}
		}
	}
	return newG
}

func (g d17HGrid) countActive() int {
	count := 0
	for w := range g {
		for z := range g[w] {
			for y := range g[w][z] {
				for x := range g[w][z][y] {
					if g[w][z][y][x] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (g d17HGrid) get(x, y, z, w int, defaultVal d17Coord) d17Coord {
	if w < 0 || w >= len(g) {
		return defaultVal
	}
	if z < 0 || z >= len(g[w]) {
		return defaultVal
	}
	if y < 0 || y >= len(g[w][z]) {
		return defaultVal
	}
	if x < 0 || x >= len(g[w][z][y]) {
		return defaultVal
	}
	return g[w][z][y][x]
}

func (g d17HGrid) copy() d17HGrid {
	newGrid := make(d17HGrid, len(g))
	for w := range g {
		newGrid[w] = g[w].copy()
	}
	return newGrid
}

func (d *Day17) simulateHGrid(g d17HGrid) d17HGrid {
	newG := g.copy()
	for w := range g {
		for z := range g[w] {
			for y := range g[w][z] {
				for x := range g[w][z][y] {
					count := 0
					for i := range d.neighborHVectors {
						v := d.neighborHVectors[i]
						if g.get(x+v.x, y+v.y, z+v.z, w+v.w, false) {
							count++
						}
					}
					newG[w][z][y][x] = (g[w][z][y][x] && count >= 2 && count <= 3) ||
						(!g[w][z][y][x] && count == 3)
				}
			}
		}
	}
	return newG
}

func (g d17Grid) countActive() int {
	count := 0
	for z := range g {
		for y := range g[z] {
			for x := range g[z][y] {
				if g[z][y][x] {
					count++
				}
			}
		}
	}
	return count
}

func (g d17Grid) print() {
	fmt.Println()
	for z := range g {
		fmt.Printf("\nZ: %d\n", z)
		for y := range g[z] {
			for x := range g[z][y] {
				if g[z][y][x] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
	}
	fmt.Println()
}

const d17ExpandPlane = 20
const d17ExpandPlanes = 10

func (d *Day17) SolveI(input string) int64 {
	d.init()

	originalP := d.retrieveGridPlane(input)
	p := originalP.expand(d17ExpandPlane)

	grid := d.makeGrid(d17ExpandPlanes, p)
	for i := 0; i < 6; i++ {
		grid = d.simulateGrid(grid)
	}

	return int64(grid.countActive())
}

func (d *Day17) SolveII(input string) int64 {
	d.init()

	originalP := d.retrieveGridPlane(input)
	p := originalP.expand(d17ExpandPlane)

	hgrid := d.makeHGrid(d17ExpandPlanes, p)
	for i := 0; i < 6; i++ {
		hgrid = d.simulateHGrid(hgrid)
	}
	return int64(hgrid.countActive())
}
