package assignment

import (
	"fmt"
	"strings"
)

type Day19 struct{}

const (
	d19MinMatch = 12
)

type d19Coord struct {
	x, y, z int
}

type d19Scanner struct {
	id        int
	relCoords []d19Coord
}

func (d *d19Scanner) orientations() []d19Coord {
	coords := make([]d19Coord, len(d.relCoords)*24)

	i := 0
	for _, c := range d.relCoords {
		// Y up
		coords[i] = d19Coord{c.x, c.y, c.z}
		coords[i+1] = d19Coord{-c.z, c.y, c.x}
		coords[i+2] = d19Coord{-c.x, c.y, -c.z}
		coords[i+3] = d19Coord{c.z, c.y, -c.x}
		// X up
		coords[i+4] = d19Coord{c.z, c.x, c.y}
		coords[i+5] = d19Coord{-c.y, c.x, c.z}
		coords[i+6] = d19Coord{-c.z, c.x, -c.y}
		coords[i+7] = d19Coord{c.y, c.x, -c.z}
		// Z up7
		coords[i+8] = d19Coord{c.y, c.z, c.x}
		coords[i+9] = d19Coord{-c.x, c.z, c.y}
		coords[i+10] = d19Coord{-c.y, c.z, -c.x}
		coords[i+11] = d19Coord{c.x, c.z, -c.y}

		for j := i; j < i+12; j++ {
			mirror := coords[i]
			coords[j+12] = d19Coord{-mirror.x, mirror.y, mirror.z}
		}

		i += 24
	}

	return coords
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

func (d *Day19) SolveI(input string) int64 {
	scanners := d.getScanners(input)

	o := scanners[0].orientations()
	fmt.Println(o)

	return 0
}

func (d *Day19) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}
