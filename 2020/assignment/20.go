package assignment

import (
	"fmt"
	"log"
	"strings"
)

type Day20 struct{}

const (
	d20Size = 10
	d20Last = d20Size - 1
)

const (
	d20Top = iota
	d20Right
	d20Bottom
	d20Left
)

type d20Image struct {
	pix [][]bool
}

func (img *d20Image) rotate() {
	nw := make([][]bool, len(img.pix))
	for i := 0; i < len(img.pix); i++ {
		nw[i] = make([]bool, len(img.pix[i]))
		for j := 0; j < len(img.pix[i]); j++ {
			nw[i][j] = img.pix[len(img.pix[i])-1-j][i]
		}
	}
	img.pix = nw
}

func (img *d20Image) flipHorizontal() {
	nw := make([][]bool, len(img.pix))
	for y := range img.pix {
		nw[y] = make([]bool, len(img.pix[y]))
		for x := range img.pix[y] {
			nw[y][x] = img.pix[y][len(img.pix[y])-1-x]
		}
	}
	img.pix = nw
}

func (img *d20Image) flipVertical() {
	nw := make([][]bool, len(img.pix))

	for y := range img.pix {
		nw[y] = make([]bool, len(img.pix[y]))
		for x := range img.pix[y] {
			nw[y][x] = img.pix[len(img.pix[y])-1-y][x]
		}
	}
	img.pix = nw
}

type d20Tile struct {
	d20Image

	id   int64
	lock bool

	neighbours [4]*d20Tile
}

func (d *Day20) opposite(which int) int {
	switch which {
	case d20Top:
		return d20Bottom
	case d20Right:
		return d20Left
	case d20Bottom:
		return d20Top
	case d20Left:
		return d20Right
	}
	panic("which border?")
}

func (t *d20Tile) String() string {
	return fmt.Sprintf("T-%d %v", t.id, t.lock)
}

func (t *d20Tile) print() {
	fmt.Printf("--- Tile %d: ---\n", t.id)
	for y := range t.pix {
		for x := range t.pix[y] {
			if t.pix[y][x] {
				print("#")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
}

func (t *d20Tile) getBorder(which int) d20Border {
	switch which {
	case d20Top:
		return t.pix[0]
	case d20Right:
		b := make(d20Border, d20Size)
		for y := 0; y < d20Size; y++ {
			b[y] = t.pix[y][d20Last]
		}
		return b
	case d20Bottom:
		return t.pix[d20Last]
	case d20Left:
		b := make(d20Border, d20Size)
		for y := 0; y < d20Size; y++ {
			b[y] = t.pix[y][0]
		}
		return b
	}
	panic("which border?")
}

func (t *d20Tile) getBorders() []d20Border {
	return []d20Border{
		t.getBorder(d20Top),
		t.getBorder(d20Right),
		t.getBorder(d20Bottom),
		t.getBorder(d20Left),
	}
}

type d20Border []bool

func (b d20Border) equals(b2 d20Border) bool {
	for i := range b {
		if b[i] != b2[i] {
			return false
		}
	}
	return true
}

func (d *Day20) getTiles(input string) []*d20Tile {
	split := strings.Split(input, "\n\n")

	tiles := make([]*d20Tile, 0, len(split))
	for i := range split {
		tileStr := SplitLines(split[i])
		tile := &d20Tile{}

		n, err := fmt.Sscanf(tileStr[0], "Tile %d:", &tile.id)
		if err != nil || n != 1 {
			log.Panicf("[%s] error parsing id line: %v | %d", tileStr[0], err, n)
		}

		tile.pix = make([][]bool, len(tileStr)-1)
		for y, line := range tileStr[1:] {
			tile.pix[y] = make([]bool, len(tileStr)-1)
			for x, val := range strings.Split(line, "") {
				switch val {
				case ".":
				case "#":
					tile.pix[y][x] = true
				default:
					panic("invalid value")
				}
			}
		}
		tiles = append(tiles, tile)
	}
	return tiles
}

func (d *Day20) findTopLeft(tiles []*d20Tile) *d20Tile {
	for _, t := range tiles {
		if t.neighbours[d20Top] == nil && t.neighbours[d20Left] == nil {
			return t
		}
	}
	panic("no left top corner found")
}

func (d *Day20) matchBorder(tile *d20Tile, side int, tiles []*d20Tile) *d20Tile {
	b1 := tile.getBorder(side)

	for _, t := range tiles {
	tileRotateLoop:
		for rot := 0; rot < 4; rot++ {
			for flip := 0; flip < 4; flip++ {
				if t.getBorder(d.opposite(side)).equals(b1) {
					return t
				}

				if t.lock {
					break tileRotateLoop
				}
				if flip%2 == 0 {
					t.flipHorizontal()
				} else {
					t.flipVertical()
				}
			}
			t.rotate()
		}
	}
	return nil
}

func (d *Day20) orientTiles(tiles []*d20Tile) *d20Tile {
	tiles[0].lock = true

	oriented := true
	for oriented {
		oriented = false

		for _, t1 := range tiles {
			if !t1.lock {
				continue
			}

			for side := 0; side < 4; side++ {
				if t1.neighbours[side] != nil {
					continue
				}

				match := d.matchBorder(t1, side, tiles)
				if match != nil {
					match.lock = true
					t1.neighbours[side] = match
					match.neighbours[d.opposite(side)] = t1
					oriented = true
				}
			}
		}
	}
	return nil
}

func (d *Day20) SolveI(input string) int64 {
	tiles := d.getTiles(input)

	d.orientTiles(tiles)

	multi := int64(1)
	for _, t := range tiles {
		sum := 0
		for _, neighbour := range t.neighbours {
			if neighbour != nil {
				sum++
			}
		}
		if sum == 2 {
			multi *= t.id
		}
	}
	return multi
}

func (d *Day20) SolveII(input string) int64 {
	tiles := d.getTiles(input)

	d.orientTiles(tiles)
	tl := d.findTopLeft(tiles)

	fmt.Println(tl)

	return 0
}
