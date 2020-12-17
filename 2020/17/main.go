package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type coord bool
type plane [][]coord
type grid []plane
type hgrid []grid

func retrieveGridPlane(file string) plane {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input plane
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		var dim []coord
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

func (p plane) expand(expandSize int) plane {
	newP := make(plane, expandSize)
	for y := range newP {
		newP[y] = make([]coord, expandSize)
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

func makeGrid(expandPlanes int, startPlane plane) grid {
	grid := make(grid, 2*expandPlanes+1)
	for z := range grid {
		grid[z] = make(plane, len(startPlane))
		for y := range grid[z] {
			grid[z][y] = make([]coord, len(startPlane[y]))
		}
	}
	grid[expandPlanes] = startPlane
	return grid
}

func makeHGrid(expandPlanes int, startPlane plane) hgrid {
	nwG := make(hgrid, 2*expandPlanes+1)
	for w := range nwG {
		nwG[w] = make(grid, len(startPlane))
		for z := range nwG[w] {
			nwG[w][z] = make(plane, len(startPlane))
			for y := range nwG[w][z] {
				nwG[w][z][y] = make([]coord, len(startPlane[y]))
			}
		}
	}
	nwG[expandPlanes][expandPlanes] = startPlane
	return nwG
}

type vector struct {
	x, y, z int
}

type hvector struct {
	x, y, z, w int
}

var neighborVectors []vector
var neighborHVectors []hvector

func init() {
	// Generate some vectors (tedious)
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x != 0 || y != 0 || z != 0 {
					neighborVectors = append(neighborVectors, vector{
						x: x,
						y: y,
						z: z,
					})
				}
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					neighborHVectors = append(neighborHVectors, hvector{
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

func (g grid) get(x, y, z int, defaultVal coord) coord {
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

func (g grid) copy() grid {
	newGrid := make(grid, len(g))
	for z := range g {
		newGrid[z] = make(plane, len(g[z]))
		for y := range g[z] {
			newGrid[z][y] = make([]coord, len(g[z][y]))
			for x := range g[z][y] {
				newGrid[z][y][x] = g[z][y][x]
			}
		}
	}
	return newGrid
}

func (g grid) simulate() grid {
	newG := g.copy()
	for z := range g {
		for y := range g[z] {
			for x := range g[z][y] {
				count := 0
				for i := range neighborVectors {
					v := neighborVectors[i]
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

func (g hgrid) countActive() int {
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

func (g hgrid) get(x, y, z, w int, defaultVal coord) coord {
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

func (g hgrid) copy() hgrid {
	newGrid := make(hgrid, len(g))
	for w := range g {
		newGrid[w] = g[w].copy()
	}
	return newGrid
}

func (g hgrid) simulate() hgrid {
	newG := g.copy()
	for w := range g {
		for z := range g[w] {
			for y := range g[w][z] {
				for x := range g[w][z][y] {
					count := 0
					for i := range neighborHVectors {
						v := neighborHVectors[i]
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

func (g grid) countActive() int {
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

func (g grid) print() {
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

const expandPlane = 64
const expandPlanes = 24

func main() {
	wd, _ := os.Getwd()
	originalP := retrieveGridPlane(filepath.Join(wd, "17/input.txt"))

	p := originalP.expand(expandPlane)

	grid := makeGrid(expandPlanes, p)
	for i := 0; i < 6; i++ {
		grid = grid.simulate()
	}

	fmt.Printf("Active cubes: %d\n\n", grid.countActive())

	hgrid := makeHGrid(expandPlanes, p)
	for i := 0; i < 6; i++ {
		hgrid = hgrid.simulate()
	}
	fmt.Printf("Active hypercubes: %d\n\n", hgrid.countActive())
}
