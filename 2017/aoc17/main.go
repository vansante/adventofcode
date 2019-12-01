package main

import (
	"fmt"
)

func main() {
	input := 371

	var list []int
	list = append(list, 0)

	curLen := len(list)
	curPos := 0
	i := 0
	for ; i < 2017; i++ {
		curPos = (curPos + input) % curLen
		curPos++
		list = append(list[:curPos], append([]int{curLen}, list[curPos:]...)...)
		curLen++
	}

	fmt.Printf("Position: %d, Value at position: %d, Value after position: %d\n", curPos, list[curPos], list[curPos+1])

	val := 0
	for ; i < 50*1000*1000; i++ {
		curPos = 1 + (curPos + input) % curLen
		if curPos == 1 {
			val = curLen
		}
		curLen++
	}

	fmt.Printf("Value after 0: %d\n", val)
}
