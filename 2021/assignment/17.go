package assignment

import (
	"fmt"
	"math"
)

type Day17 struct{}

type d17Vector struct {
	x, y int64
}

func (v *d17Vector) add(v2 *d17Vector) {
	v.x += v2.x
	v.y += v2.y
}

func (v *d17Vector) decaySpeed() {
	if v.x > 0 {
		v.x--
	} else if v.x < 0 {
		v.x++
	}
	v.y--
}

type d17TargetArea struct {
	xMin, xMax int64
	yMin, yMax int64
}

func (a *d17TargetArea) hit(c *d17Vector) bool {
	if c.x < a.xMin || c.x > a.xMax {
		return false
	}
	if c.y < a.yMin || c.y > a.yMax {
		return false
	}
	return true
}

func (a *d17TargetArea) simulate(vec d17Vector) (maxY int64, reached bool) {
	const steps = 500

	c := &d17Vector{0, 0}
	v := &d17Vector{vec.x, vec.y}
	maxY = int64(math.MinInt)
	for i := 0; i < steps; i++ {
		c.add(v)

		if c.y > maxY {
			maxY = c.y
		}

		v.decaySpeed()

		if a.hit(c) {
			return maxY, true
		}
		if c.x > a.xMax || c.y < a.yMin {
			return maxY, false
		}
	}
	return maxY, false
}

func (d *Day17) getTargetArea(input string) d17TargetArea {
	ta := d17TargetArea{}
	lines := SplitLines(input)
	n, err := fmt.Sscanf(lines[0], "target area: x=%d..%d, y=%d..%d", &ta.xMin, &ta.xMax, &ta.yMin, &ta.yMax)
	CheckErr(err)
	if n != 4 {
		panic("error parsing")
	}
	return ta
}

func (d *Day17) SolveI(input string) int64 {
	ta := d.getTargetArea(input)

	const simulateSize = 500
	maxY := int64(math.MinInt)
	for x := -simulateSize; x < simulateSize; x++ {
		for y := -simulateSize; y < simulateSize; y++ {
			v := d17Vector{int64(x), int64(y)}

			curMaxY, hit := ta.simulate(v)
			if hit && curMaxY > maxY {
				maxY = curMaxY
			}
		}
	}

	return maxY
}

func (d *Day17) SolveII(input string) int64 {
	ta := d.getTargetArea(input)

	const simulateSize = 500
	hits := int64(0)
	for x := -simulateSize; x < simulateSize; x++ {
		for y := -simulateSize; y < simulateSize; y++ {
			v := d17Vector{int64(x), int64(y)}

			_, hit := ta.simulate(v)
			if hit {
				hits++
			}
		}
	}
	return hits
}
