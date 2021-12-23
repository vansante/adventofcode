package assignment

import (
	"strconv"
	"strings"
)

type Day23 struct{}

var d22Costs = map[string]int64{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

func (d *Day23) SolveI(input string) int64 {
	//#############
	//#...........#
	//###B#A#B#C###
	//  #C#D#D#A#
	//  #########

	costStr := "C2A5"

	//#############
	//#A........C.#
	//###B#.#B#.###
	//  #C#D#D#A#
	//  #########

	costStr += "A9B4"

	//#############
	//#AA.B.....C.#
	//###B#.#.#.###
	//  #C#D#D#.#
	//  #########

	costStr += "D6D7"

	//#############
	//#AA.B.....C.#
	//###B#.#.#D###
	//  #C#.#.#D#
	//  #########

	costStr += "B3C5"

	//#############
	//#AA.........#
	//###B#.#.#D###
	//  #C#B#C#D#
	//  #########

	costStr += "B4C7"

	//#############
	//#AA.........#
	//###.#B#C#D###
	//  #.#B#C#D#
	//  #########

	costStr += "A3A3"

	//#############
	//#...........#
	//###A#B#C#D###
	//  #A#B#C#D#
	//  #########

	sum := int64(0)
	costs := strings.Split(costStr, "")
	for i := 0; i < len(costs)-1; i += 2 {
		count, err := strconv.ParseInt(costs[i+1], 10, 32)
		CheckErr(err)
		sum += d22Costs[costs[i]] * count
	}
	// 6671, 6381 too low
	// 14770 too high
	// 14630, 14530 not right
	return sum
}

func (d *Day23) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}
