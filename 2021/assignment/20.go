package assignment

import (
	"math"
	"strings"
)

type Day20 struct{}

const (
	d20StartIdx = 10_000
)

type d20EnhanceAlg struct {
	bits [512]bool
}

type d20Image struct {
	pixels                 map[uint64]bool
	minX, minY, maxX, maxY int
}

func (i *d20Image) coordToIdx(x, y int) uint64 {
	return uint64(y)<<32 | uint64(x)
}

func (i *d20Image) idxToCoord(idx uint64) (x, y int) {
	x = int(idx & 0xffffffff)
	y = int(idx >> 32)
	return x, y
}

func (i *d20Image) setPixel(x, y int, val bool) {
	if x > i.maxX {
		i.maxX = x
	}
	if x < i.minX {
		i.minX = x
	}
	if y > i.maxY {
		i.maxY = y
	}
	if y < i.minY {
		i.minY = y
	}
	i.pixels[i.coordToIdx(x, y)] = val
}

func (i *d20Image) pixel(x, y int) bool {
	return i.pixels[i.coordToIdx(x, y)]
}

func (i *d20Image) print() {
	for y := i.minY; y <= i.maxY; y++ {
		for x := i.minX; x <= i.maxX; x++ {
			if i.pixel(x, y) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

type d20Vector struct {
	x, y int
}

var (
	d20Vectors = []d20Vector{
		{-1, -1}, // top left
		{0, -1},  // top
		{1, -1},  // top right
		{-1, 0},  // left
		{0, 0},   // self
		{1, 0},   // right
		{-1, 1},  // bottom left
		{0, 1},   // bottom
		{1, 1},   // bottom right
	}
)

func (i *d20Image) enhance(alg d20EnhanceAlg) d20Image {
	nw := d20Image{
		pixels: make(map[uint64]bool, len(i.pixels)*2),
		minX:   math.MaxInt,
		minY:   math.MaxInt,
	}

	for y := i.minY - 1; y <= i.maxY+1; y++ {
		for x := i.minX - 1; x <= i.maxX+1; x++ {
			var num int
			for n, v := range d20Vectors {
				bit := 0
				if i.pixel(x+v.x, y+v.y) {
					bit = 1
				}
				num += bit << (8 - n)
			}
			nw.setPixel(x, y, alg.bits[num])
		}
	}
	return nw
}

func (i *d20Image) countLit() int64 {
	sum := int64(0)
	for _, px := range i.pixels {
		if px {
			sum++
		}
	}
	return sum
}

func (d *Day20) readInput(input string) (d20EnhanceAlg, d20Image) {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) != 2 {
		panic("invalid input")
	}
	alg := d20EnhanceAlg{}
	chars := strings.Split(blocks[0], "")
	for i, char := range chars {
		switch char {
		case "#":
			alg.bits[i] = true
		case ".":
		case "\n":
			break
		default:
			panic("unknown character")
		}
	}

	img := d20Image{
		pixels: make(map[uint64]bool, 1024),
		minX:   math.MaxInt,
		minY:   math.MaxInt,
	}
	lines := SplitLines(blocks[1])
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			switch char {
			case ".":
				img.setPixel(d20StartIdx+x, d20StartIdx+y, false)
			case "#":
				img.setPixel(d20StartIdx+x, d20StartIdx+y, true)
			default:
				panic("unknown character")
			}
		}
	}
	return alg, img
}

func (d *Day20) SolveI(input string) int64 {
	alg, img := d.readInput(input)

	img.print()
	img = img.enhance(alg)
	img.print()
	img = img.enhance(alg)
	img.print()
	//for idx, bool := range img.pixels {
	//if !bool {
	//	continue
	//}
	//x, y := img.idxToCoord(idx)
	//fmt.Println(x, y, bool)
	//}
	return img.countLit()
}

func (d *Day20) SolveII(input string) int64 {
	return 0
}
