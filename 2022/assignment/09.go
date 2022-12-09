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

var d09Start = d08Coord{0, 5}

func (d *Day09) walkDirections(dirs []d09Direction, walker func(head, tail d08Coord)) {
	head := d09Start
	tail := d09Start

	for _, d := range dirs {
		for i := 0; i < d.amount; i++ {
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
			xDiff := tail.x - head.x
			yDiff := tail.y - head.y

			switch {
			// Diagonal differences:
			case util.Abs(yDiff) == 1 && util.Abs(xDiff) == 2:
				tail.x -= xDiff / 2
				tail.y -= yDiff
			case util.Abs(xDiff) == 1 && util.Abs(yDiff) == 2:
				tail.y -= yDiff / 2
				tail.x -= xDiff
			// Straight differences:
			case yDiff == 0 && util.Abs(xDiff) > 1:
				tail.x -= xDiff / 2
			case xDiff == 0 && util.Abs(yDiff) > 1:
				tail.y -= yDiff / 2
			}

			walker(head, tail)
		}
	}
}

func (d *Day09) printGrid(head, tail d08Coord, size int) {
	fmt.Println()
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			switch {
			case head.x == x && head.y == y:
				print("H")
			case tail.x == x && tail.y == y:
				print("T")
			default:
				print(".")
			}
		}
		print("\n")
	}
	fmt.Println()
}

func (d *Day09) SolveI(input string) any {
	dirs := d.getDirections(input)

	visited := make(map[d08Coord]struct{}, 2048)
	d.walkDirections(dirs, func(head, tail d08Coord) {
		//d.printGrid(head, tail, 6)

		visited[tail] = struct{}{}
	})
	return len(visited)
}

func (d *Day09) SolveII(input string) any {
	return "Not Implemented Yet"
}
