package assignment

import (
	"fmt"
	"math"

	"github.com/vansante/adventofcode/2022/util"
)

type Day23 struct{}

type d23Coord struct {
	x, y int
}

func (c d23Coord) add(other d23Coord) d23Coord {
	return d23Coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

var (
	d23Vectors = []d23Coord{
		{0, -1},  // top
		{1, -1},  // top right
		{1, 0},   // right
		{1, 1},   // bottom right
		{0, 1},   // bottom
		{-1, 1},  // bottom left
		{-1, 0},  // left
		{-1, -1}, // top left
	}

	d23Cardinals = [][]d23Coord{
		{d23Vectors[7], d23Vectors[0], d23Vectors[1]}, // North
		{d23Vectors[3], d23Vectors[4], d23Vectors[5]}, // South
		{d23Vectors[5], d23Vectors[6], d23Vectors[7]}, // West
		{d23Vectors[1], d23Vectors[2], d23Vectors[3]}, // East
	}
)

type d23Elf struct {
	position d23Coord
	proposal *d23Coord
	round    int
}

type d23Grid struct {
	coords map[d23Coord]*d23Elf
}

func (g *d23Grid) print() {
	min, max := g.minMax()

	fmt.Println()
	for y := min.y - 1; y <= max.y+1; y++ {
		for x := min.x - 1; x <= max.x+1; x++ {
			if g.coords[d23Coord{x, y}] == nil {
				print(".")
				continue
			}

			print("#")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day23) getElves(input string) d23Grid {
	lines := util.SplitLines(input)

	elves := d23Grid{
		coords: make(map[d23Coord]*d23Elf, 1024),
	}

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '.':
				continue
			case '#':
				elves.coords[d23Coord{x, y}] = &d23Elf{
					position: d23Coord{x, y},
				}
			default:
				panic("invalid grid")
			}
		}
	}
	return elves
}

func (g *d23Grid) hasElf(start d23Coord, vectors []d23Coord) bool {
	for _, vec := range vectors {
		if g.coords[start.add(vec)] != nil {
			return true
		}
	}
	return false
}

func (g *d23Grid) propose(round int) map[d23Coord][]*d23Elf {
	proposals := make(map[d23Coord][]*d23Elf, 1024)

	for c, e := range g.coords {
		if !g.hasElf(c, d23Vectors) {
			continue // Not going to do anything!
		}

		for i := round; i < round+len(d23Cardinals); i++ {
			cardinal := i % len(d23Cardinals)
			if g.hasElf(c, d23Cardinals[cardinal]) {
				continue // Not going to move here!
			}

			coord := c.add(d23Cardinals[cardinal][1])
			e.proposal = &coord
			proposals[coord] = append(proposals[coord], e)
			break
		}
	}

	return proposals
}

func (g *d23Grid) moveElf(elf *d23Elf, nw d23Coord) {
	if g.coords[nw] != nil {
		panic("move onto elf")
	}

	g.coords[nw] = elf             // Set elf at new coord
	delete(g.coords, elf.position) // remove from current coord
	elf.position = nw
}

func (g *d23Grid) round(round int) int {
	proposals := g.propose(round)
	if len(proposals) == 0 {
		return 0
	}

	moved := 0
	for _, e := range g.coords {
		if e.proposal == nil {
			e.round++
			continue
		}

		coord := *e.proposal
		// Reset proposal for next round
		e.proposal = nil
		switch len(proposals[coord]) {
		case 0:
			panic("proposal expected")
		case 1:
			// We can move
			g.moveElf(e, coord)
			moved++
		default:
			// More than one, cancel move
			continue
		}
	}
	return moved
}

func (g *d23Grid) move(maxRounds int) int {
	for i := 0; i < maxRounds; i++ {
		var moves int
		moves = g.round(i)
		if moves == 0 {
			return i + 1
		}

		//fmt.Printf("==== ROUND %d ====", i+1)
		//g.print()
	}
	return maxRounds
}

func (g *d23Grid) minMax() (min d23Coord, max d23Coord) {
	min = d23Coord{x: math.MaxInt, y: math.MaxInt}
	max = d23Coord{x: math.MinInt, y: math.MinInt}
	for c := range g.coords {
		min.x = util.Min(min.x, c.x)
		min.y = util.Min(min.y, c.y)

		max.x = util.Max(max.x, c.x)
		max.y = util.Max(max.y, c.y)
	}
	return min, max
}

func (g *d23Grid) emptyTiles() int {
	min, max := g.minMax()

	x := util.Abs(max.x-min.x) + 1
	y := util.Abs(max.y-min.y) + 1

	return x*y - len(g.coords)
}

func (d *Day23) SolveI(input string) any {
	elves := d.getElves(input)

	elves.move(10)
	return elves.emptyTiles()
}

func (d *Day23) SolveII(input string) any {
	elves := d.getElves(input)
	return elves.move(math.MaxInt)
}
