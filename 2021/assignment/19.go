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

func (c d19Coord) String() string {
	return fmt.Sprintf("[%d,%d,%d]", c.x, c.y, c.z)
}

func (c d19Coord) distance(other d19Coord) int {
	return AbsInt(other.x-c.x) + AbsInt(other.y-c.y) + AbsInt(other.z-c.z)
}

func (c d19Coord) add(other d19Coord) d19Coord {
	return d19Coord{c.x + other.x, c.y + other.y, c.z + other.z}
}

func (c d19Coord) subtract(other d19Coord) d19Coord {
	return d19Coord{c.x - other.x, c.y - other.y, c.z - other.z}
}

func (c d19Coord) equals(other d19Coord) bool {
	return c.x == other.x && c.y == other.y && c.z == other.z
}

func (c d19Coord) orientations() []d19Coord {
	x, y, z := c.x, c.y, c.z
	return []d19Coord{
		{x, y, z},
		{-y, x, z},
		{-x, -y, z},
		{y, -x, z},

		{-x, y, -z},
		{y, x, -z},
		{x, -y, -z},
		{-y, -x, -z},

		{-z, y, x},
		{-z, x, -y},
		{-z, -y, -x},
		{-z, -x, y},

		{z, y, -x},
		{z, x, y},
		{z, -y, x},
		{z, -x, -y},

		{x, -z, y},
		{-y, -z, x},
		{-x, -z, -y},
		{y, -z, -x},

		{x, z, -y},
		{-y, z, -x},
		{-x, z, y},
		{y, z, x},
	}
}

type d19Scanner struct {
	id          int
	relCoords   []d19Coord
	distances   [][]int
	translation d19Coord
	oriented    bool
}

func (s d19Scanner) String() string {
	return fmt.Sprintf("Scanner{%02d}", s.id)
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

func (s *d19Scanner) setOrientation(translation d19Coord, rotation int) bool {
	if translation.equals(d19Coord{}) && rotation == 0 {
		return false
	}

	fmt.Printf("%s | Set translation %s + rotation: %d\n", s, translation, rotation)
	if s.oriented {
		panic("already oriented")
	}
	s.translation = translation
	s.oriented = true
	for i := range s.relCoords {
		s.relCoords[i] = s.relCoords[i].orientations()[rotation].subtract(translation)
	}
	return true
}

func (s *d19Scanner) markBeacons(mp map[string]int) {
	for _, c := range s.relCoords {
		mp[c.String()]++
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

func (d *Day19) distIntersect(s1, s2 []int) int {
	matches := 0
	for _, search := range s1 {
		idx := sort.Search(len(s2), func(i int) bool {
			return s2[i] >= search
		})
		if idx < len(s2) && s2[idx] == search {
			matches++
		}
	}
	return matches
}

func (d *Day19) matchOrientations(sc1, sc2 *d19Scanner, minMatches int) bool {
	type pair struct {
		b1, b2 d19Coord
	}
	var beaconPair *pair
	for b1, list1 := range sc1.distances {
		for b2, list2 := range sc2.distances {
			hits := d.distIntersect(list1, list2)
			if hits >= minMatches {
				beaconPair = &pair{sc1.relCoords[b1], sc2.relCoords[b2]}
				break
			}
		}
	}

	if beaconPair == nil {
		return false
	}

	var rotation int
	var b2Coord d19Coord
	for rotation, b2Coord = range beaconPair.b2.orientations() {
		translation := b2Coord.subtract(beaconPair.b1)

		var matches int
		for _, coord2 := range sc2.relCoords {
			transCoord := coord2.orientations()[rotation].subtract(translation)

			for _, coord1 := range sc1.relCoords {
				if coord1.equals(transCoord) {
					matches++
				}
			}

			if matches >= minMatches {
				return sc2.setOrientation(translation, rotation)
			}
		}
	}
	return false
}

func (d *Day19) matchScannerOrientations(scanners []d19Scanner) {
	for i := range scanners {
		scanners[i].calcDistances()
	}

	scanners[0].oriented = true
	for i := 1; i < len(scanners); i++ {
		if !d.matchOrientations(&scanners[0], &scanners[i], d19MinMatch-1) {
			continue
		}
	}

	oriented := true
	for oriented {
		oriented = false
		for i := range scanners {
			if !scanners[i].oriented {
				continue
			}
			for j := range scanners {
				if i == j || scanners[j].oriented {
					continue
				}

				oriented = oriented || d.matchOrientations(&scanners[i], &scanners[j], d19MinMatch-1)
			}
		}
	}
}

func (d *Day19) SolveI(input string) int64 {
	scanners := d.getScanners(input)
	d.matchScannerOrientations(scanners)

	mp := make(map[string]int)
	for _, s := range scanners {
		if !s.oriented {
			panic(fmt.Sprintf("%s is not oriented", s))
		}
		s.markBeacons(mp)
	}

	return int64(len(mp))
}

func (d *Day19) SolveII(input string) int64 {
	scanners := d.getScanners(input)
	d.matchScannerOrientations(scanners)

	max := math.MinInt
	for _, scan1 := range scanners {
		for _, scan2 := range scanners {
			dist := scan1.translation.distance(scan2.translation)
			if dist > max {
				max = dist
			}
		}
	}

	return int64(max)
}
