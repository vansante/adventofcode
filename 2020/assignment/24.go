package assignment

import "fmt"

type Day24 struct{}

type d24Vector struct {
	x, y, z int
}

var (
	// https://stackoverflow.com/questions/2049196/generating-triangular-hexagonal-coordinates-xyz
	d24VectorMap = map[string]d24Vector{
		"e":  {1, -1, 0},
		"se": {1, 0, -1},
		"sw": {0, 1, -1},
		"w":  {-1, 1, 0},
		"nw": {-1, 0, 1},
		"ne": {0, -1, 1},
	}
	d24Vectors []d24Vector
)

func init() {
	d24Vectors = make([]d24Vector, len(d24VectorMap))
	i := 0
	for k := range d24VectorMap {
		d24Vectors[i] = d24VectorMap[k]
		i++
	}
}

func (v *d24Vector) add(v2 d24Vector) {
	v.x += v2.x
	v.y += v2.y
	v.z += v2.z
}

type d24Grid struct {
	tiles [][][]bool
}

const (
	d24Size   = 160
	d24Center = d24Size / 2
)

func (g *d24Grid) flip(vs []d24Vector) {
	sum := d24Vector{d24Center, d24Center, d24Center}
	for _, v := range vs {
		sum.add(v)
	}
	g.tiles[sum.x][sum.y][sum.z] = !g.tiles[sum.x][sum.y][sum.z]
}

func (g *d24Grid) get(x, y, z int, defaultVal bool) bool {
	if x < 0 || x >= len(g.tiles) {
		return defaultVal
	}
	if y < 0 || y >= len(g.tiles[x]) {
		return defaultVal
	}
	if z < 0 || z >= len(g.tiles[x][y]) {
		return defaultVal
	}
	return g.tiles[x][y][z]
}

func (g *d24Grid) countBlackNeighbours(x, y, z int) int {
	sum := 0
	for i := range d24Vectors {
		v := d24Vector{x, y, z}
		v.add(d24Vectors[i])
		if g.get(v.x, v.y, v.z, false) {
			sum++
		}
	}
	return sum
}

func (g *d24Grid) countBlack() int64 {
	var sum int64
	for x := range g.tiles {
		for y := range g.tiles[x] {
			for z := range g.tiles[x][y] {
				if g.tiles[x][y][z] {
					sum++
				}
			}
		}
	}
	return sum
}

func (g *d24Grid) passDay() {
	flip := make([]d24Vector, 0, 128)

	for x := range g.tiles {
		for y := range g.tiles[x] {
			for z := range g.tiles[x][y] {
				neighbours := g.countBlackNeighbours(x, y, z)
				if g.tiles[x][y][z] && (neighbours == 0 || neighbours > 2) {
					flip = append(flip, d24Vector{x, y, z})
				}

				if !g.tiles[x][y][z] && neighbours == 2 {
					flip = append(flip, d24Vector{x, y, z})
				}
			}
		}
	}

	for _, f := range flip {
		g.tiles[f.x][f.y][f.z] = !g.tiles[f.x][f.y][f.z]
	}
}

func (g *d24Grid) passDays(n int) {
	for i := 0; i < n; i++ {
		g.passDay()
	}
}

func (g *d24Grid) init(size int) {
	g.tiles = make([][][]bool, size)
	for i := range g.tiles {
		g.tiles[i] = make([][]bool, size)
		for j := range g.tiles[i] {
			g.tiles[i][j] = make([]bool, size)
		}
	}
}

func (d *Day24) readVectors(input string) [][]d24Vector {
	lines := SplitLines(input)
	vs := make([][]d24Vector, len(lines))

	for i, line := range lines {
		for len(line) > 0 {
			v, ok := d24VectorMap[line[:1]]
			if ok {
				line = line[1:]
				vs[i] = append(vs[i], v)
				continue
			}
			v, ok = d24VectorMap[line[:2]]
			if !ok {
				panic(fmt.Sprintf("vector %s not found", line[:2]))
			}
			line = line[2:]
			vs[i] = append(vs[i], v)
		}
	}
	return vs
}

func (d *Day24) init(input string) d24Grid {
	directions := d.readVectors(input)

	g := d24Grid{}
	g.init(d24Size)

	for i := range directions {
		g.flip(directions[i])
	}
	return g
}

func (d *Day24) SolveI(input string) int64 {
	g := d.init(input)

	return g.countBlack()
}

func (d *Day24) SolveII(input string) int64 {
	g := d.init(input)
	g.passDays(100)
	return g.countBlack()
}
