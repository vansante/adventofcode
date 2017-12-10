package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc10.2/input.txt")
	if err != nil {
		panic(err)
	}

	var list []int
	for i := 0; i <= 255; i++ {
		list = append(list, i)
	}

	var lengths []int
	for i := range input {
		lengths = append(lengths, int(input[i]))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	fmt.Printf("Lengths: %#v\n", lengths)

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

	var denseHash []int
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

	fmt.Printf("Dense hash: %#v\n", denseHash)
	fmt.Printf("Hash: ")
	for i := range denseHash {
		fmt.Printf("%x", denseHash[i]/16)
		fmt.Printf("%x", denseHash[i]%16)
	}
	fmt.Printf("\n")
}
