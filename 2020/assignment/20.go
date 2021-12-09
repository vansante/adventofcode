package assignment

import (
	"fmt"
	"log"
	"strings"
)

type Day20 struct{}

const (
	d09Size = 10
	d09Last = d09Size - 1
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
		pix: make([][]bool, d09Size),
	}
	for i := 0; i < d09Size; i++ {
		nw.pix[i] = make([]bool, d09Size)
		for j := 0; j < d09Size; j++ {
			nw.pix[i][j] = t.pix[d09Last-j][i]
		}
	}
	return nw
}

func (t *d20Tile) flipHorizontal() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, d09Size),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, d09Size)
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[y][d09Last-x]
		}
	}
	return nw
}

func (t *d20Tile) flipVertical() *d20Tile {
	nw := &d20Tile{
		id:  t.id,
		pix: make([][]bool, d09Size),
	}

	for y := range t.pix {
		nw.pix[y] = make([]bool, d09Size)
		for x := range t.pix[y] {
			nw.pix[y][x] = t.pix[d09Last-y][x]
		}
	}
	return nw
}

func (t *d20Tile) rowMatch(other *d20Tile, y int) bool {
	for x := range t.pix[y] {
		if t.pix[y][x] != other.pix[d09Last-y][x] {
			return false
		}
	}
	return true
}

func (t *d20Tile) colMatch(other *d20Tile, x int) bool {
	for y := range t.pix {
		if t.pix[y][x] != other.pix[y][d09Last-x] {
			return false
		}
	}
	return true
}

func (t *d20Tile) findNeighbours(others []*d20Tile) (top, right, bottom, left *d20Tile) {
	bla := 0
	for i := range others {
		oth := others[i]

		if t.id == oth.id {
			continue // dont match ourselves :x
		}

		for rotOth := 0; rotOth < 4; rotOth++ {
			curOth := oth
			for flipOth := 0; flipOth < 4; flipOth++ {
				if t.rowMatch(curOth, 0) {
					if top != nil && top.id != curOth.id {
						panic("more than 1 top")
					}
					top = curOth
				}
				if t.rowMatch(curOth, d09Last) {
					if bottom != nil && bottom.id != curOth.id {
						panic("more than 1 bottom")
					}
					bottom = curOth
				}
				if t.colMatch(curOth, 0) {
					if left != nil && left.id != curOth.id {
						panic("more than 1 left")
					}
					left = curOth
				}
				if t.colMatch(curOth, d09Last) {
					if right != nil && right.id != curOth.id {
						panic("more than 1 right")
					}
					right = curOth
				}

				if flipOth%2 == 0 {
					curOth = curOth.flipHorizontal()
				} else {
					curOth = curOth.flipVertical()
				}
				bla++
			}
			curOth = curOth.rotate()
		}
	}

	return top, right, bottom, left
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

func (d *Day20) SolveI(input string) int64 {
	tiles := d.getTiles(input)

	for _, t := range tiles {
		top, right, bottom, left := t.findNeighbours(tiles)

		fmt.Println(t, "Top: ", top)
		fmt.Println(t, "Right: ", right)
		fmt.Println(t, "Bottom: ", bottom)
		fmt.Println(t, "Left: ", left)
		//break
	}

	return 0
}

func (d *Day20) SolveII(input string) int64 {
	return 0
}
