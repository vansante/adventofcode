package assignment

import (
	"fmt"
	"log"
	"strings"
)

type Day20 struct{}

type d20Tile struct {
	id  int
	pix [][]bool
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

func (t *d20Tile) rotate() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, len(t.pix)),
	}
	for i := 0; i < len(t.pix); i++ {
		nw.pix[i] = make([]bool, len(t.pix))
		for j := 0; j < len(t.pix); j++ {
			nw.pix[i][j] = t.pix[len(t.pix)-1-j][i]
		}
	}
	return nw
}

func (t *d20Tile) flipHorizontal() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, len(t.pix)),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, len(t.pix))
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[y][len(t.pix)-x-1]
		}
	}
	return nw
}

func (t *d20Tile) flipVertical() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, len(t.pix)),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, len(t.pix))
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[len(t.pix)-y-1][x]
		}
	}
	return nw
}

func (d *Day20) getTiles(input string) []d20Tile {
	split := strings.Split(input, "\n\n")

	tiles := make([]d20Tile, 0, len(split))
	for i := range split {
		tileStr := SplitLines(split[i])
		tile := d20Tile{}

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

func (d *Day20) SolveI(input string) int64 {
	tiles := d.getTiles(input)

	//fmt.Println(tiles)
	tiles[0].print()
	tiles[0].rotate().print()
	return 0
}

func (d *Day20) SolveII(input string) int64 {
	return 0
}
