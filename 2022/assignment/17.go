package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day17 struct{}

type d17Shape struct {
	width, height int
	xy            [4][4]bool
}

var d17Shapes = []d17Shape{
	{ // --
		width:  4,
		height: 1,
		xy: [4][4]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
	},
	{ // +
		width:  3,
		height: 3,
		xy: [4][4]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		},
	},
	{ // inverse L
		width:  3,
		height: 3,
		xy: [4][4]bool{
			{true, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, false, false},
		},
	},
	{ // |
		width:  1,
		height: 4,
		xy: [4][4]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		},
	},
	{ // Square
		width:  2,
		height: 2,
		xy: [4][4]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, false, false},
			{false, false, false, false},
		},
	},
}

const d17Width = 7

type d17Direction uint8

func (d *d17Direction) String() string {
	switch *d {
	case d17Left:
		return "L"
	case d17Right:
		return "R"
	}
	return "?"
}

const (
	d17Left d17Direction = iota + 1
	d17Right
)

func (d *Day17) getDirections(input string) []d17Direction {
	chars := strings.TrimSpace(input)
	dirs := make([]d17Direction, len(chars))
	for i := range chars {
		switch chars[i] {
		case '<':
			dirs[i] = d17Left
		case '>':
			dirs[i] = d17Right
		default:
			panic("unknown character")
		}
	}
	return dirs
}

type d17Coord struct {
	x, y int
}

type d17Cave struct {
	rocks      map[d17Coord]struct{}
	maxHeight  int
	maxHeights [d17Width]int
	shapeIdx   int
	dirs       []d17Direction
	dirIdx     int
}

func (d *Day17) makeCave(dirs []d17Direction) d17Cave {
	return d17Cave{
		rocks:     make(map[d17Coord]struct{}),
		maxHeight: 0,
		dirs:      dirs,
	}
}

func (c *d17Cave) state() d17State {
	s := d17State{
		shapeIdx: c.shapeIdx,
		dirIdx:   c.dirIdx,
	}
	min := util.MinSlice(c.maxHeights[:])
	for i := range s.heights {
		s.heights[i] = c.maxHeights[i] - min
	}
	return s
}

func (c *d17Cave) shape() d17Shape {
	shape := d17Shapes[c.shapeIdx]
	c.shapeIdx++
	c.shapeIdx %= len(d17Shapes)
	return shape
}

func (c *d17Cave) direction() d17Direction {
	dir := c.dirs[c.dirIdx]
	c.dirIdx++
	c.dirIdx %= len(c.dirs)
	return dir
}

func (c *d17Cave) print() {
	for y := c.maxHeight + 4; y >= 1; y-- {
		fmt.Println()
		fmt.Printf("%03d |", y)
		for x := 0; x < d17Width; x++ {
			if _, ok := c.rocks[d17Coord{x, y}]; ok {
				print("#")
				continue
			}
			print(".")
		}
		print("|")
	}
	fmt.Println()
	fmt.Println("   ", strings.Repeat("=", d17Width+2))
}

func (c *d17Cave) set(x, y int) {
	if x < 0 || x >= d17Width {
		panic("invalid x")
	}
	if y <= 0 {
		panic("invalid y")
	}
	if c.get(x, y) {
		panic("already occupied")
	}

	c.maxHeight = util.Max(c.maxHeight, y)
	c.maxHeights[x] = util.Max(c.maxHeights[x], y)
	c.rocks[d17Coord{x, y}] = struct{}{}
}

func (c *d17Cave) get(x, y int) bool {
	if x < 0 || x >= d17Width {
		return true
	}
	if y <= 0 {
		return true
	}

	_, ok := c.rocks[d17Coord{x, y}]
	return ok
}

func (c *d17Cave) hasCollision(s d17Shape, posX, posY int) bool {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			// Skip parts of shape not occupied
			if !s.xy[y][x] {
				continue
			}
			// If shape and cave occupied, collision
			if c.get(x+posX, y+posY) {
				return true
			}
		}
	}
	return false
}

func (c *d17Cave) dropShape() {
	s := c.shape()

	x := 2
	y := c.maxHeight + 4

	if c.hasCollision(s, x, y) {
		panic("collision at start?")
	}

	for i := 0; ; i++ {
		// First move left/right
		switch c.direction() {
		case d17Left:
			if !c.hasCollision(s, x-1, y) {
				x--
				//print("L")
			}
		case d17Right:
			if !c.hasCollision(s, x+1, y) {
				x++
				//print("R")
			}
		}
		// Now drop
		if c.hasCollision(s, x, y-1) {
			c.setShape(s, x, y)
			return // Done!
		}
		y--
	}
}

func (c *d17Cave) setShape(s d17Shape, posX, posY int) {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			// Skip parts of shape not occupied
			if !s.xy[y][x] {
				continue
			}
			c.set(x+posX, y+posY)
		}
	}
}

type d17State struct {
	shapeIdx int
	dirIdx   int
	heights  [d17Width]int
}

func (c *d17Cave) findRepetition() int {
	states := make(map[d17State]struct{}, 50_000)
	for i := 0; ; i++ {
		state := c.state()
		if _, ok := states[state]; ok {
			return i
		}

		states[state] = struct{}{}
		c.dropShape()
	}
}

func (c *d17Cave) tetris(rocks int) {
	for i := 0; i < rocks; i++ {
		c.dropShape()
	}
}

func (d *Day17) SolveI(input string) any {
	dirs := d.getDirections(input)

	c := d.makeCave(dirs)
	c.tetris(2022)

	return c.maxHeight
}

func (d *Day17) SolveII(input string) any {
	const total = 1_000_000_000_000

	dirs := d.getDirections(input)
	c := d.makeCave(dirs)

	rep := c.findRepetition()
	startHeight := c.maxHeight
	rep = c.findRepetition()
	repetitionHeight := c.maxHeight - startHeight

	amount := total / int64(rep)
	height := int64(repetitionHeight) * amount
	remainder := total % int64(rep)

	c = d.makeCave(dirs)
	c.tetris(int(remainder))

	return height + int64(c.maxHeight)
}
