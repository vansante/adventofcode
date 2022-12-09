package assignment

import (
	"fmt"

	"github.com/vansante/adventofcode/2022/util"
)

type Day09 struct{}

type d09Direction struct {
	dir    string
	amount int
}

type d09Coord struct {
	x, y int
}

func (d *Day09) getDirections(input string) []d09Direction {
	lines := util.SplitLines(input)
	dirs := make([]d09Direction, len(lines))
	for i, line := range lines {
		n, err := fmt.Sscanf(line, "%s %d", &dirs[i].dir, &dirs[i].amount)
		util.CheckErr(err)
		if n != 2 {
			panic("invalid scan")
		}
	}

	return dirs
}

func (d *Day09) walkDirections(dirs []d09Direction, start d09Coord, ropeSize int, walker func(rope []d09Coord)) {
	if ropeSize < 2 {
		panic("rope too short")
	}

	rope := make([]d09Coord, ropeSize)
	for i := range rope {
		rope[i] = start
	}

	for _, d := range dirs {
		for i := 0; i < d.amount; i++ {
			head := &rope[0]
			// Move head:
			switch d.dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "D":
				head.y++
			case "U":
				head.y--
			default:
				panic("invalid direction")
			}

			// Move tail
			h := head
			for j := 1; j < len(rope); j++ {
				t := &rope[j]
				xDiff := t.x - h.x
				yDiff := t.y - h.y

				switch {
				// Diagonal differences:
				case util.Abs(yDiff) == 2 && util.Abs(xDiff) == 2:
					t.x -= xDiff / 2
					t.y -= yDiff / 2
				case util.Abs(yDiff) == 1 && util.Abs(xDiff) == 2:
					t.x -= xDiff / 2
					t.y -= yDiff
				case util.Abs(xDiff) == 1 && util.Abs(yDiff) == 2:
					t.y -= yDiff / 2
					t.x -= xDiff
				// Straight differences:
				case yDiff == 0 && util.Abs(xDiff) > 1:
					t.x -= xDiff / 2
				case xDiff == 0 && util.Abs(yDiff) > 1:
					t.y -= yDiff / 2
				}

				if j == len(rope)-1 {
					break
				}
				h = &rope[j]
			}
			walker(rope)
		}
	}
}

func (d *Day09) printGrid(rope []d09Coord, size int) {
	fmt.Println()
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			printed := false
			for i := range rope {
				if rope[i].x != x || rope[i].y != y {
					continue
				}
				if printed {
					continue
				}

				printed = true
				if i == 0 {
					print("H")
				} else {
					fmt.Printf("%d", i)
				}
			}
			if !printed {
				print(".")
			}
		}
		print("\n")
	}
	fmt.Println()
}

func (d *Day09) SolveI(input string) any {
	dirs := d.getDirections(input)

	visited := make(map[d09Coord]struct{}, 2048)
	d.walkDirections(dirs, d09Coord{0, 5}, 2, func(rope []d09Coord) {
		//d.printGrid(rope, 6)

		visited[rope[1]] = struct{}{}
	})
	return len(visited)
}

func (d *Day09) SolveII(input string) any {
	dirs := d.getDirections(input)

	visited := make(map[d09Coord]struct{}, 2048)
	d.walkDirections(dirs, d09Coord{12, 15}, 10, func(rope []d09Coord) {
		//d.printGrid(rope, 27)

		visited[rope[9]] = struct{}{}
	})
	return len(visited)

}
