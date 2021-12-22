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

func (c d22Cuboid) equals(other d22Cuboid) bool {
	return c.x1 == other.x1 && c.x2 == other.x2 &&
		c.y1 == other.y1 && c.y2 == other.y2 &&
		c.z1 == other.z1 && c.z2 == other.z2
}

func (c d22Cuboid) contains(other d22Cuboid) bool {
	return c.x1 <= other.x1 && c.x2 >= other.x2 &&
		c.y1 <= other.y1 && c.y2 >= other.y2 &&
		c.z1 <= other.z1 && c.z2 >= other.z2
}

func (c d22Cuboid) intersects(other d22Cuboid) bool {
	return c.x1 <= other.x2 && c.x2 >= other.x1 &&
		c.y1 <= other.y2 && c.y2 >= other.y1 &&
		c.z1 <= other.z2 && c.z2 >= other.z1
}
func (c d22Cuboid) volume() int64 {
	return int64(c.x2-c.x1) * int64(c.y2-c.y1) * int64(c.z2-c.z1)
}

func (c d22Cuboid) intersection(other d22Cuboid) d22Cuboid {
	return d22Cuboid{
		x1: MaxInt(c.x1, other.x1),
		x2: MinInt(c.x2, other.x2),
		y1: MaxInt(c.y1, other.y1),
		y2: MinInt(c.y2, other.y2),
		z1: MaxInt(c.x1, other.x1),
		z2: MinInt(c.x2, other.x2),
	}
}

func (c d22Cuboid) subtract(other d22Cuboid) []d22Cuboid {
	if other.contains(c) {
		return nil
	}
	if !c.intersects(other) {
		return []d22Cuboid{c}
	}

	var xSplits, ySplits, zSplits []int
	if c.x1 < other.x1 && other.x1 < c.x2 {
		xSplits = append(xSplits, other.x1)
	}
	if c.x1 < other.x2 && other.x2 < c.x2 {
		xSplits = append(xSplits, other.x2)
	}
	if c.y1 < other.y1 && other.y1 < c.y2 {
		ySplits = append(ySplits, other.y1)
	}
	if c.y1 < other.y2 && other.y2 < c.y2 {
		ySplits = append(ySplits, other.y2)
	}
	if c.z1 < other.z1 && other.z1 < c.z2 {
		zSplits = append(zSplits, other.z1)
	}
	if c.z1 < other.z2 && other.z2 < c.z2 {
		zSplits = append(zSplits, other.z2)
	}

	xv := append([]int{c.x1}, append(xSplits, c.x2)...)
	yv := append([]int{c.y1}, append(ySplits, c.y2)...)
	zv := append([]int{c.z1}, append(zSplits, c.z2)...)

	cubes := make([]d22Cuboid, 0, len(xv)*len(yv)*len(zv))
	for i := 0; i < len(xv)-1; i++ {
		for j := 0; j < len(yv)-1; j++ {
			for k := 0; k < len(zv)-1; k++ {
				cube := d22Cuboid{
					x1: xv[i],
					x2: xv[i+1],
					y1: yv[j],
					y2: yv[j+1],
					z1: zv[k],
					z2: zv[k+1],
				}
				if !other.contains(cube) {
					cubes = append(cubes, cube)
				}
			}
		}
	}

	return cubes
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

func (d *Day22) reboot(cubs []d22Cuboid) int64 {
	grid := make(map[d22Coord]bool, 1024)
	for _, c := range cubs {
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
	return int64(len(grid))
}

func (d *Day22) SolveI(input string) int64 {
	cubs := d.getCuboids(input)

	return d.reboot(cubs)
}

func (d *Day22) rebootII(cubs []d22Cuboid) int64 {
	list := make([]d22Cuboid, 0, len(cubs))

	for _, c := range cubs {
		c.x2++
		c.y2++
		c.z2++

		nwList := make([]d22Cuboid, 0, len(cubs))
		for i := 0; i < len(list); i++ {
			cubes := list[i].subtract(c)

			nwList = append(nwList, cubes...)
		}

		if c.on {
			nwList = append(nwList, c)
		}
		list = nwList
	}

	sum := int64(0)
	for i := range list {
		sum += list[i].volume()
	}
	return sum
}

func (d *Day22) SolveII(input string) int64 {
	cubs := d.getCuboids(input)
	return d.rebootII(cubs)
}
