package assignment

import (
	"fmt"
	"log"
	"math"
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

func (d *Day20) makeImage(size int) *d20Image {
	img := &d20Image{
		pix: make([][]bool, size),
	}
	for i := 0; i < size; i++ {
		img.pix[i] = make([]bool, size)
	}
	return img
}

func (img *d20Image) print() {
	for y := range img.pix {
		for x := range img.pix[y] {
			if img.pix[y][x] {
				print("#")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
}

func (img *d20Image) get(x, y int, defaultVal bool) bool {
	if y < 0 || y >= len(img.pix) {
		return defaultVal
	}
	if x < 0 || x >= len(img.pix[0]) {
		return defaultVal
	}
	return img.pix[y][x]
}

type d20Vector struct {
	x, y int
}

func (img *d20Image) getMonsterVectors() []d20Vector {
	monster := `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

	lines := strings.Split(monster, "\n")
	var vectors []d20Vector
	for y := range lines {
		for x, char := range strings.Split(lines[y], "") {
			if char == "#" {
				vectors = append(vectors, d20Vector{x, y})
			}
		}
	}
	return vectors
}

func (img *d20Image) findMonsters() (monsters, blips int) {
	vs := img.getMonsterVectors()

	for y := 0; y < len(img.pix); y++ {
		for x := 0; x < len(img.pix[y]); x++ {
			if img.pix[y][x] {
				blips++
			}

			found := true
			for _, v := range vs {
				found = found && img.get(x+v.x, y+v.y, false)
			}
			if found {
				monsters++
			}
		}
	}
	return monsters, blips - (monsters * len(vs))
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
	t.d20Image.print()
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

func (d *Day20) tileImage(tiles []*d20Tile) *d20Image {
	tileWidth := int(math.Sqrt(float64(len(tiles))))
	pixPerTile := len(tiles[0].pix) - 2

	img := d.makeImage(tileWidth * pixPerTile)

	y := 0
	cur := d.findTopLeft(tiles)
	for cur != nil {
		x := 0
		line := cur
		for line != nil {
			tileY := y
			for ty := 1; ty < len(line.pix)-1; ty++ {
				tileX := x
				for tx := 1; tx < len(line.pix)-1; tx++ {
					img.pix[tileY][tileX] = line.pix[ty][tx]
					tileX++
				}
				tileY++
			}

			x += len(line.pix) - 2
			line = line.neighbours[d20Right]
		}

		y += len(cur.pix) - 2
		cur = cur.neighbours[d20Bottom]
	}
	return img
}

func (d *Day20) SolveII(input string) int64 {
	tiles := d.getTiles(input)

	d.orientTiles(tiles)

	img := d.tileImage(tiles)

	img.print()

	for rot := 0; rot < 4; rot++ {
		for flip := 0; flip < 4; flip++ {
			monsters, blips := img.findMonsters()
			if monsters > 0 {
				return int64(blips)
			}

			if flip%2 == 0 {
				img.flipHorizontal()
			} else {
				img.flipVertical()
			}
		}
		img.rotate()
	}

	panic("no monsters found")
}
