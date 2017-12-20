package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Vector struct {
	X int64
	Y int64
	Z int64
}

func (v *Vector) ManhattanDistance(v2 *Vector) int64 {
	x := math.Abs(float64(v.X - v2.X))
	y := math.Abs(float64(v.Y - v2.Y))
	z := math.Abs(float64(v.Z - v2.Z))

	return int64(x + y + z)
}

func (v *Vector) Equal(v2 *Vector) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

type Particle struct {
	Pos Vector
	Vel Vector
	Acc Vector
}

func (p *Particle) Step() {
	p.Vel.X += p.Acc.X
	p.Vel.Y += p.Acc.Y
	p.Vel.Z += p.Acc.Z

	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y
	p.Pos.Z += p.Vel.Z
}

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc20/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	zeroVector := Vector{}

	var particles []*Particle
	for i := range lines {
		p := ParticleFromLine(lines[i])
		particles = append(particles, &p)
	}

	idx := -1
	minAccDistance := int64(math.MaxInt64)
	minVelDistance := int64(math.MaxInt64)
	maxPosDistance := int64(-1)
	for i := range particles {
		accDistance := particles[i].Acc.ManhattanDistance(&zeroVector)
		velDistance := particles[i].Vel.ManhattanDistance(&zeroVector)
		if accDistance < minAccDistance {
			minAccDistance = accDistance
			minVelDistance = velDistance
			idx = i
		} else if accDistance == minAccDistance {
			if velDistance < minVelDistance {
				minAccDistance = accDistance
				minVelDistance = velDistance
				idx = i
			}
		}

		if maxPosDistance < particles[i].Pos.ManhattanDistance(&zeroVector) {
			maxPosDistance = particles[i].Pos.ManhattanDistance(&zeroVector)
		}
	}

	fmt.Printf("Lowest: %d at %d\n", minAccDistance, idx)
	fmt.Printf("Max distance: %d\n", maxPosDistance)

	particleMap := make(map[int]*Particle)
	for p := range particles {
		particleMap[p] = particles[p]
	}

	for i := int64(0); i < maxPosDistance*2; i++ {
		for p := range particleMap {
			particles[p].Step()
		}

		var removals []int
		for p1 := range particleMap {
			for p2 := range particleMap {
				if p1 == p2 {
					continue
				}
				if particleMap[p1] == nil || particleMap[p2] == nil {
					continue
				}
				if particleMap[p1].Pos.Equal(&particleMap[p2].Pos) {
					fmt.Printf("p%d [%v] <=> p%d [%v]\n", p1, particleMap[p1].Pos, p2, particleMap[p2].Pos)
					removals = append(removals, p1, p2)
				}
			}
		}
		for _, i := range removals {
			delete(particleMap, i)
		}
		fmt.Printf("Step %d | %d\n", i, len(particleMap))
	}

	fmt.Printf("Count: %d\n", len(particleMap))
}

func ParticleFromLine(line string) (p Particle) {
	_, err := fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.Pos.X, &p.Pos.Y, &p.Pos.Z, &p.Vel.X, &p.Vel.Y, &p.Vel.Z, &p.Acc.X, &p.Acc.Y, &p.Acc.Z)
	if err != nil {
		panic(err)
	}
	return
}

