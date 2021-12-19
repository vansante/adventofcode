package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Day19 struct{}

const (
	d19MinMatch = 12
)

type d19Coord struct {
	x, y, z int
}

func (c d19Coord) distance(other d19Coord) int {
	return AbsInt(other.x-c.x) + AbsInt(other.y-c.y) + AbsInt(other.z-c.z)
}

func (c d19Coord) equals(other d19Coord) bool {
	return c.x == other.x && c.y == other.y && c.z == other.z
}

func (c d19Coord) orientations() []d19Coord {
	coords := make([]d19Coord, 24)
	// Y up
	coords[0] = d19Coord{c.x, c.y, c.z}
	coords[1] = d19Coord{-c.z, c.y, c.x}
	coords[2] = d19Coord{-c.x, c.y, -c.z}
	coords[3] = d19Coord{c.z, c.y, -c.x}
	// X up
	coords[4] = d19Coord{c.z, c.x, c.y}
	coords[5] = d19Coord{-c.y, c.x, c.z}
	coords[6] = d19Coord{-c.z, c.x, -c.y}
	coords[7] = d19Coord{c.y, c.x, -c.z}
	// Z up
	coords[8] = d19Coord{c.y, c.z, c.x}
	coords[9] = d19Coord{-c.x, c.z, c.y}
	coords[10] = d19Coord{-c.y, c.z, -c.x}
	coords[11] = d19Coord{c.x, c.z, -c.y}

	for i := 0; i < 12; i++ {
		mirror := coords[i]
		coords[i+12] = d19Coord{-mirror.x, mirror.y, mirror.z}
	}
	return coords
}

type d19Scanner struct {
	id        int
	relCoords []d19Coord
	distances [][]int
}

func (s *d19Scanner) calcDistances() {
	s.distances = make([][]int, len(s.relCoords))
	for i, c1 := range s.relCoords {
		s.distances[i] = make([]int, 0, len(s.relCoords))

		for j, c2 := range s.relCoords {
			if i == j {
				continue
			}

			s.distances[i] = append(s.distances[i], c1.distance(c2))
		}
		sort.Ints(s.distances[i])
	}
}

func (d *Day19) getScanners(input string) []d19Scanner {
	blocks := strings.Split(input, "\n\n")

	scanners := make([]d19Scanner, 0, 100)
	for _, block := range blocks {
		lines := SplitLines(block)
		s := d19Scanner{
			relCoords: make([]d19Coord, len(lines)-1),
		}
		n, err := fmt.Sscanf(lines[0], "--- scanner %d ---", &s.id)
		CheckErr(err)
		if n != 1 {
			panic("error scanning scanner line")
		}

		for i, line := range lines[1:] {
			n, err := fmt.Sscanf(line, "%d,%d,%d",
				&s.relCoords[i].x,
				&s.relCoords[i].y,
				&s.relCoords[i].z,
			)
			CheckErr(err)
			if n != 3 {
				panic("error scanning scanner coord")
			}
		}
		scanners = append(scanners, s)
	}
	return scanners
}

func (d *Day19) compare(sc1, sc2 d19Scanner) (X, Y, Z, rot int) {
	maxHits := math.MinInt
	d.exploreCoordinates(func(x, y, z int) {
		var hits int
		for i := 0; i < 24; i++ {
			for _, coord1 := range sc1.relCoords {
				for _, coord2 := range sc2.relCoords {
					coord2 = coord2.orientations()[i]
					if coord2.equals(coord1) {
						hits++
						break
					}
				}
			}
			if hits > maxHits && hits > d19MinMatch {
				maxHits = hits
				X, Y, Z = x, y, z
				rot = i
			}
		}
	})
	return X, Y, Z, rot
}

func (d *Day19) exploreCoordinates(explore func(x, y, z int)) {
	const (
		minCoord = -1000
		maxCoord = 1000
	)
	for x := minCoord; x < maxCoord; x++ {
		for y := minCoord; y < maxCoord; y++ {
			for z := minCoord; z < maxCoord; z++ {
				explore(x, y, z)
			}
		}
	}
}

func (d *Day19) SolveI(input string) int64 {
	scanners := d.getScanners(input)

	for i := range scanners {
		scanners[i].calcDistances()
	}

	for _, scan1 := range scanners {
		for _, scan2 := range scanners {
			if scan1.id == scan2.id {
				continue
			}
			d.compare(scan1, scan2)
		}
	}

	o := scanners[0].relCoords[3].orientations()
	fmt.Println(o)

	return 0
}

func (d *Day19) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}
