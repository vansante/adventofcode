package main

import (
	"fmt"
	"math/bits"
)

type Point struct {
	x int
	y int
}

func main() {
	input := "xlqgujun"

	grid := make([][]int, 128)
	usedBits := 0
	groupId := 1
	for i := 0; i < len(grid); i++ {
		curInput := []byte(fmt.Sprintf("%s-%d", input, i))

		denseHash := KnotHash(curInput)
		for j := range denseHash {
			usedBits += bits.OnesCount(uint(denseHash[j]))
			for b := 7; b >= 0; b-- {
				mask := 1 << uint(b)
				var num int
				//fmt.Printf("%d ", exp)
				if denseHash[j]&mask != 0 {
					num = groupId
					groupId++
				}
				grid[i] = append(grid[i], num)
			}
		}
	}
	fmt.Printf("Used bits: %#v\n", usedBits)

	for i := 0; i < len(grid)*2; i++ {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == 0 {
					continue
				}
				points := getAdjacentPoints(x, y, len(grid)-1)

				ownVal := grid[y][x]
				for _, p := range points {
					val := grid[p.y][p.x]
					if val == 0 {
						continue
					}

					if val < ownVal {
						grid[y][x] = val
					} else if val > ownVal {
						grid[p.y][p.x] = ownVal
					}
				}
			}
		}
	}

	//for i := range grid {
	//	for j := range grid[i] {
	//		fmt.Printf("%4d ", grid[i][j])
	//	}
	//	fmt.Printf("\n")
	//}

	groupMap := make(map[int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			val := grid[y][x]
			groupMap[val] = true
		}
	}

	// FIXME: Subtract 1 for the zero group
	fmt.Printf("Total groups: %#v\n", len(groupMap)-1)
}

func getAdjacentPoints(x, y, max int) []Point {
	var points []Point
	if x > 0 {
		points = append(points, Point{x - 1, y})
	}
	if y > 0 {
		points = append(points, Point{x, y - 1})
	}
	if x < max {
		points = append(points, Point{x + 1, y})
	}
	if y < max {
		points = append(points, Point{x, y + 1})
	}

	return points
}

func KnotHash(input []byte) (denseHash []int) {
	var list []int
	for i := 0; i <= 255; i++ {
		list = append(list, i)
	}

	var lengths []int
	for i := range input {
		lengths = append(lengths, int(input[i]))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	currentPosition := 0
	skipSize := 0
	listLength := len(list)
	for r := 0; r < 64; r++ {
		for i := range lengths {
			curLength := lengths[i]
			startIdx := currentPosition
			endIdx := (currentPosition + curLength - 1) % listLength

			for j := 0; j < curLength/2; j++ {
				list[startIdx], list[endIdx] = list[endIdx], list[startIdx]

				startIdx = (startIdx + 1) % listLength
				endIdx--
				if endIdx < 0 {
					endIdx += listLength
				}
			}

			currentPosition = (currentPosition + curLength + skipSize) % listLength
			skipSize++
		}
	}

	var curVal int
	for i := range list {
		if i == 0 {
			curVal = list[i]
			continue
		}
		if i%16 == 0 {
			denseHash = append(denseHash, curVal)
			curVal = list[i]
			continue
		}
		curVal ^= list[i]
	}
	denseHash = append(denseHash, curVal)
	return
}
