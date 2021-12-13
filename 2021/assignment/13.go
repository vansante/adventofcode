package assignment

import (
	"fmt"
	"log"
	"strings"
)

type Day13 struct{}

const d13GridSize = 2_000

type d13Vector struct {
	x, y int
}

type d13Fold struct {
	axis string
	n    int
}

type d13Grid struct {
	y      []d13Line
	width  int
	height int
}

type d13Line struct {
	x []bool
}

func (g *d13Grid) print() {
	for y := range g.y {
		if y > g.height {
			continue
		}
		for x := range g.y[y].x {
			if x > g.width {
				continue
			}
			if g.y[y].x[x] {
				print("#")
			} else {
				print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *d13Grid) markVector(v d13Vector) {
	g.y[v.y].x[v.x] = true
	if v.y > g.height {
		g.height = v.y + 1
	}
	if v.x > g.width {
		g.width = v.x + 1
	}
}

func (g *d13Grid) fold(f d13Fold) {
	switch f.axis {
	case "x":
		g.foldX(f.n)
	case "y":
		g.foldY(f.n)
	default:
		panic("invalid axis")
	}
}

func (g *d13Grid) foldX(fold int) {
	if fold > g.width {
		panic("too big")
	}
	for y := range g.y {
		for x := fold; x < g.width; x++ {
			if !g.y[y].x[x] {
				continue
			}
			newX := fold - (x - fold)
			g.y[y].x[newX] = true
		}
	}
	g.width = fold
}

func (g *d13Grid) foldY(fold int) {
	if fold > g.height {
		panic("too big")
	}
	for y := fold; y < g.height; y++ {
		for x := range g.y[y].x {
			if !g.y[y].x[x] {
				continue
			}
			newY := fold - (y - fold)
			g.y[newY].x[x] = true
		}
	}
	g.height = fold
}

func (g *d13Grid) count() int64 {
	sum := int64(0)
	for y := 0; y <= g.height; y++ {
		for x := 0; x <= g.width; x++ {
			if g.y[y].x[x] {
				sum++
			}
		}
	}
	return sum
}

func (d *Day13) makeGrid() *d13Grid {
	g := &d13Grid{}
	g.y = make([]d13Line, d13GridSize)
	for y := range g.y {
		g.y[y].x = make([]bool, d13GridSize)
	}
	return g
}

func (d *Day13) ReadInput(input string) ([]d13Vector, []d13Fold) {
	spl := strings.Split(input, "\n\n")
	if len(spl) != 2 {
		panic("invalid input")
	}
	vectorStr, foldStr := spl[0], spl[1]
	vectors := make([]d13Vector, 0)
	for _, line := range SplitLines(vectorStr) {
		var x, y int
		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil || n != 2 {
			log.Panicf("[%s] error parsing line: %v | %d", line, err, n)
		}
		vectors = append(vectors, d13Vector{x, y})
	}

	folds := make([]d13Fold, 0)
	for _, line := range SplitLines(foldStr) {
		var num int
		var axis string
		n, err := fmt.Sscanf(strings.Replace(line, "=", " ", 1), "fold along %s %d", &axis, &num)
		if err != nil || n != 2 {
			log.Panicf("[%s] error parsing line: %v | %d", line, err, n)
		}
		folds = append(folds, d13Fold{axis, num})
	}
	return vectors, folds
}

func (d *Day13) SolveI(input string) int64 {
	vecs, folds := d.ReadInput(input)
	g := d.makeGrid()
	for _, v := range vecs {
		g.markVector(v)
	}

	g.fold(folds[0])
	return g.count()
}

func (d *Day13) SolveII(input string) int64 {
	vecs, folds := d.ReadInput(input)
	g := d.makeGrid()
	for _, v := range vecs {
		g.markVector(v)
	}

	for _, f := range folds {
		g.fold(f)
	}
	g.print()
	return 0 // The output is a string
}
