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

type d20Tile struct {
	id  int64
	pix [][]bool
}

func (t *d20Tile) String() string {
	return fmt.Sprintf("T-%d", t.id)
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
		pix: make([][]bool, d20Size),
	}
	for i := 0; i < d20Size; i++ {
		nw.pix[i] = make([]bool, d20Size)
		for j := 0; j < d20Size; j++ {
			nw.pix[i][j] = t.pix[d20Last-j][i]
		}
	}
	return nw
}

func (t *d20Tile) flipHorizontal() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, d20Size),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, d20Size)
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[y][d20Last-x]
		}
	}
	return nw
}

func (t *d20Tile) flipVertical() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, d20Size),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, d20Size)
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[d20Last-y][x]
		}
	}
	return nw
}

func (t *d20Tile) rowMatch(other *d20Tile, y int) bool {
	for x := range t.pix[y] {
		if t.pix[y][x] != other.pix[d20Last-y][x] {
			return false
		}
	}
	return true
}

func (t *d20Tile) colMatch(other *d20Tile, x int) bool {
	for y := range t.pix {
		if t.pix[y][x] != other.pix[y][d20Last-x] {
			return false
		}
	}
	return true
}

func (t *d20Tile) getBorders() []d20Border {
	b := make([]d20Border, 4)
	b[1] = make(d20Border, d20Size)
	b[3] = make(d20Border, d20Size)
	for y := 0; y < d20Size; y++ {
		b[1][y] = t.pix[y][d20Last]
		b[3][y] = t.pix[y][0]
	}
	b[0] = t.pix[0]
	b[2] = t.pix[d20Last]
	return b
}

type d20Border []bool

func (b d20Border) flip() d20Border {
	nw := make(d20Border, len(b))
	for i := range b {
		nw[len(b)-i-1] = b[i]
	}
	return nw
}

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

func (d *Day20) StitchImage(tiles []d20Tile) {
	for _, t := range tiles {
		tileBorders := t.getBorders()

		for _, t2 := range tiles {
			if t.id == t2.id {
				continue
			}
			tile2Borders := t2.getBorders()
			for _, tb := range tileBorders {
				for _, tb2 := range tile2Borders {
					if !tb.equals(tb2) && !tb.flip().equals(tb2) {
						continue
					}

				}
			}
		}
	}
}

func (d *Day20) SolveI(input string) int64 {
	tiles := d.getTiles(input)

	multi := int64(1)
	for _, t := range tiles {
		tileBorders := t.getBorders()

		neighbourSum := 0
		for _, t2 := range tiles {
			if t.id == t2.id {
				continue
			}
			tile2Borders := t2.getBorders()
			for _, tb := range tileBorders {
				for _, tb2 := range tile2Borders {
					if tb.equals(tb2) || tb.flip().equals(tb2) {
						neighbourSum++
					}
				}
			}
		}

		switch neighbourSum {
		case 0, 1:
			panic("invalid amount of neighbours")
		case 2:
			multi *= t.id
		}
	}
	return multi
}

func (d *Day20) SolveII(input string) int64 {
	return 0
}
