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
	d22FaceRight d22Facing = iota
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

type d22Wrapper func(c d22Coord, facing d22Facing) (bool, d22Coord, d22Facing)

func (g *d22Grid) wrapLinear(c d22Coord, facing d22Facing) (bool, d22Coord, d22Facing) {
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
	return g.get(c) == d22TypeOpen, c, facing
}

const (
	d22Side = 50
)

// Warning, hardcoded wrapping coming up
func (g *d22Grid) wrapCube(c d22Coord, f d22Facing) (bool, d22Coord, d22Facing) {
	switch {
	case c.x <= 2*d22Side && c.y <= d22Side: // side A
		switch f {
		case d22FaceUp: // Wrap to F
			f = d22FaceRight
			c.y = 3*d22Side + c.x
			c.x = 1
		case d22FaceLeft: // Wrap to D
			f = d22FaceRight
			c.y = 2*d22Side + (d22Side - c.x)
			c.x = 1
		default:
			panic("invalid wrap")
		}
	case c.x > 2*d22Side && c.y <= d22Side: // Side B
		switch f {
		case d22FaceUp: // Wrap to F
			f = d22FaceUp
			c.x = c.x - d22Side
			c.y = 3 * d22Side
		case d22FaceRight: // Wrap to E
			f = d22FaceLeft
			c.x = 2 * d22Side
			c.y = 3*d22Side - c.y
		case d22FaceDown: // Wrap to C
			f = d22FaceLeft
			c.y = c.x
			c.x = 2 * d22Side
		default:
			panic("invalid wrap")
		}
	case c.x <= 2*d22Side && c.y <= d22Side*2: // Side C
		switch f {
		case d22FaceLeft: // Wrap to D
			f = d22FaceDown
			c.x = c.y - d22Side
			c.y = 2*d22Side + 1
		case d22FaceRight: // Wrap to B
			f = d22FaceUp
			c.x = c.y
			c.y = d22Side
		default:
			panic("invalid wrap")
		}
	case c.x <= d22Side && c.y <= d22Side*3: // Side D
		switch f {
		case d22FaceUp: // Wrap to C
			f = d22FaceRight
			c.y = d22Side + c.x
			c.x = d22Side + 1
		case d22FaceLeft: // Wrap to A
			f = d22FaceRight
			c.y = d22Side - (c.x - 2*d22Side)
			c.x = d22Side + 1
		default:
			panic("invalid wrap")
		}
	case c.x <= 2*d22Side && c.y <= d22Side*3: // Side E
		switch f {
		case d22FaceRight: // Wrap to B
			f = d22FaceLeft
			c.y = d22Side - (c.y - 2*d22Side)
			c.x = 3 * d22Side
		case d22FaceDown: // Wrap to F
			f = d22FaceLeft
			c.y = 3*d22Side + c.x
			c.x = d22Side
		default:
			panic("invalid wrap")
		}
	case c.x <= d22Side && c.y <= d22Side*4: // Side F
		switch f {
		case d22FaceLeft: // Wrap to A
			f = d22FaceDown
			c.x = c.y - 3*d22Side
			c.y = 1
		case d22FaceDown: // Wrap to B
			f = d22FaceDown
			c.x = c.x + 2*d22Side
			c.y = 1
		case d22FaceRight: // Wrap to E
			f = d22FaceUp
			c.x = c.y - 2*d22Side
			c.y = 3 * d22Side
		default:
			panic("invalid wrap")
		}
	}
	return g.get(c) == d22TypeOpen, c, f
}

func (g *d22Grid) normalizeStep(coord d22Coord, facing d22Facing, walk int, wrapFn d22Wrapper) (d22Coord, d22Facing) {
	for i := 0; i < walk; i++ {
		// check next step
		c := facing.step(coord)
		tp := g.get(c)
		switch tp {
		case d22TypeOpen:
			coord = facing.step(coord)
		case d22TypeWall:
			// Return last step
			return coord, facing
		case d22TypeNothing:
			// Wrap
			possible, c, f := wrapFn(coord, facing)
			if possible {
				coord = c
				facing = f
			}
		}
	}
	return coord, facing
}

func (g *d22Grid) findStart() d22Coord {
	c := d22Coord{1, 1}
	for g.get(c) != d22TypeOpen {
		c = d22FaceRight.step(c)
	}
	return c
}

func (g *d22Grid) walk(start d22Coord, facing d22Facing, dirs []d22Direction, wrapFn d22Wrapper) (d22Coord, d22Facing) {
	c := start
	f := facing
	for _, dir := range dirs {
		if dir.num == 0 {
			f = f.turn(dir.turn)
			continue
		}

		c, f = g.normalizeStep(c, f, dir.num, wrapFn)
	}
	return c, f
}

func (d *Day22) getCode(coords d22Coord, facing d22Facing) int {
	return (1000 * coords.y) + (4 * coords.x) + int(facing)
}

func (d *Day22) SolveI(input string) any {
	grid, directions := d.getNotes(input)
	coord, facing := grid.walk(grid.findStart(), d22FaceRight, directions, grid.wrapLinear)
	return d.getCode(coord, facing)
}

func (d *Day22) SolveII(input string) any {
	grid, directions := d.getNotes(input)
	coord, facing := grid.walk(grid.findStart(), d22FaceRight, directions, grid.wrapCube)
	// < 103019
	return d.getCode(coord, facing)
}
