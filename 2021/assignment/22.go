package assignment

import (
	"fmt"
	"strings"
)

type Day22 struct{}

type d22Coord struct {
	x, y, z int
}

type d22Cuboid struct {
	on                     bool
	x1, x2, y1, y2, z1, z2 int
}

func (d *Day22) getCuboids(input string) []d22Cuboid {
	lines := SplitLines(input)

	cuboids := make([]d22Cuboid, 0, len(lines))
	for _, line := range lines {
		c := d22Cuboid{}
		if strings.HasPrefix(line, "on") {
			c.on = true
			line = line[3:]
		} else {
			line = line[4:]
		}
		n, err := fmt.Sscanf(line, "x=%d..%d,y=%d..%d,z=%d..%d", &c.x1, &c.x2, &c.y1, &c.y2, &c.z1, &c.z2)
		CheckErr(err)
		if n != 6 {
			panic("invalid input")
		}
		cuboids = append(cuboids, c)
	}
	return cuboids
}

func (d *Day22) SolveI(input string) int64 {
	cub := d.getCuboids(input)

	grid := make(map[d22Coord]bool, 1024)
	for _, c := range cub {
		for z := c.z1; z <= c.z2; z++ {
			if z < -50 || z > 50 {
				continue
			}
			for y := c.y1; y <= c.y2; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for x := c.x1; x <= c.x2; x++ {
					if x < -50 || x > 50 {
						continue
					}
					if c.on {
						grid[d22Coord{x, y, z}] = true
					} else {
						delete(grid, d22Coord{x, y, z})
					}
				}
			}
		}
	}

	fmt.Println(len(grid))
	return int64(len(grid))
}

func (d *Day22) SolveII(input string) int64 {
	return 0
}
