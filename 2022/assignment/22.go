package assignment

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/vansante/adventofcode/2022/util"
)

type Day22 struct{}

type d22Coord struct {
	x, y int
}

type d22Tile uint8

const (
	d22TypeNothing d22Tile = iota
	d22TypeOpen
	d22TypeWall
)

type d22Grid struct {
	minX, minY, maxX, maxY int
	coords                 map[d22Coord]d22Tile
}

func (g *d22Grid) print() {
	fmt.Println()
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			switch g.coords[d22Coord{x, y}] {
			case d22TypeNothing:
				print(" ")
			case d22TypeOpen:
				print(".")
			case d22TypeWall:
				print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *d22Grid) set(x, y int, val d22Tile) {
	g.coords[d22Coord{x, y}] = val
	g.minX = util.Min(x, g.minX)
	g.maxX = util.Max(x, g.maxX)
	g.minY = util.Min(y, g.minY)
	g.maxY = util.Max(y, g.maxY)
}

func (g *d22Grid) get(coord d22Coord) d22Tile {
	return g.coords[coord]
}

type d22Direction struct {
	num  int
	turn string
}

func (d *Day22) getNotes(input string) (d22Grid, []d22Direction) {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	g := d22Grid{coords: make(map[d22Coord]d22Tile, 100_000)}
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case ' ':
				continue
			case '.':
				g.set(x+1, y+1, d22TypeOpen)
			case '#':
				g.set(x+1, y+1, d22TypeWall)
			}
		}
	}

	dirs := make([]d22Direction, 0, len(parts[1])/2)
	instr := strings.TrimSpace(parts[1])
	for i := 0; i < len(instr); i++ {
		if unicode.IsDigit(rune(instr[i])) {
			str := strings.Builder{}
			for i < len(instr) && unicode.IsDigit(rune(instr[i])) {
				str.WriteRune(rune(instr[i]))
				i++
			}
			num, err := strconv.ParseInt(str.String(), 10, 32)
			util.CheckErr(err)
			dirs = append(dirs, d22Direction{num: int(num)})
		}

		if i >= len(instr) {
			break
		}

		switch instr[i] {
		case 'L', 'R':
			dirs = append(dirs, d22Direction{turn: string(instr[i])})
		default:
			panic("invalid direction")
		}
	}

	return g, dirs
}

type d22Facing uint8

const (
	d22FaceRight = iota
	d22FaceDown
	d22FaceLeft
	d22FaceUp
)

func (f d22Facing) turn(turn string) d22Facing {
	switch turn {
	case "L":
		return (f + 4 - 1) % 4
	case "R":
		return (f + 1) % 4
	}
	panic("invalid turn")
}

func (f d22Facing) step(coord d22Coord) d22Coord {
	switch f {
	case d22FaceRight:
		coord.x++
	case d22FaceDown:
		coord.y++
	case d22FaceLeft:
		coord.x--
	case d22FaceUp:
		coord.y--
	}
	return coord
}

func (g *d22Grid) wrap(c d22Coord, facing d22Facing) (possible bool, coord d22Coord) {
	switch facing {
	case d22FaceRight:
		c.x = g.minX
	case d22FaceDown:
		c.y = g.minY
	case d22FaceLeft:
		c.x = g.maxX
	case d22FaceUp:
		c.y = g.maxY
	}
	for g.get(c) == d22TypeNothing {
		c = facing.step(c)
	}
	return g.get(c) == d22TypeOpen, c
}

func (g *d22Grid) normalizeStep(coord d22Coord, facing d22Facing, walk int) d22Coord {
	for i := 0; i < walk; i++ {
		// check next step
		switch g.get(facing.step(coord)) {
		case d22TypeOpen:
			coord = facing.step(coord)
		case d22TypeWall:
			// Return last step
			return coord
		case d22TypeNothing:
			// Wrap
			possible, newCoord := g.wrap(coord, facing)
			if possible {
				coord = newCoord
			}
		}
	}
	return coord
}

func (g *d22Grid) walk(start d22Coord, facing d22Facing, dirs []d22Direction) (d22Coord, d22Facing) {
	c := start
	f := facing
	for _, dir := range dirs {
		if dir.num == 0 {
			f = f.turn(dir.turn)
			continue
		}

		c = g.normalizeStep(c, f, dir.num)
	}
	return c, f
}

func (d *Day22) SolveI(input string) any {
	grid, directions := d.getNotes(input)
	grid.print()

	coords, facing := grid.walk(d22Coord{1, 1}, d22FaceRight, directions)

	return (1000 * coords.y) + (4 * coords.x) + int(facing)
}

func (d *Day22) SolveII(input string) any {
	return "Not Implemented Yet"
}
