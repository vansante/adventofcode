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

func (d *Day23) costs(str string) int64 {
	sum := int64(0)
	costs := strings.Split(str, "")
	for i := 0; i < len(costs)-1; i += 2 {
		count, err := strconv.ParseInt(costs[i+1], 10, 32)
		CheckErr(err)
		sum += d22Costs[costs[i]] * count
	}
	return sum
}

func (d *Day23) SolveI(input string) int64 {
	costs := ""
	//#############
	//#...........#
	//###B#A#B#C###
	//  #C#D#D#A#
	//  #########

	costs += "C2A5"

	//#############
	//#A........C.#
	//###B#.#B#.###
	//  #C#D#D#A#
	//  #########

	costs += "A9D8"

	//#############
	//#AA.......C.#
	//###B#.#B#.###
	//  #C#.#D#D#
	//  #########

	costs += "B5B4"

	//#############
	//#AA.......C.#
	//###.#B#.#.###
	//  #C#B#D#D#
	//  #########

	costs += "D5C5"

	//#############
	//#AA.........#
	//###.#B#.#D###
	//  #C#B#C#D#
	//  #########

	costs += "C7"

	//#############
	//#AA.........#
	//###.#B#C#D###
	//  #.#B#C#D#
	//  #########

	costs += "A3A3"

	//#############
	//#...........#
	//###A#B#C#D###
	//  #A#B#C#D#
	//  #########

	return d.costs(costs)
}

func (d *Day23) SolveII(input string) int64 {
	costs := ""
	//#############
	//#...........#
	//###B#A#B#C###
	//  #D#C#B#A#
	//  #D#B#A#C#
	//  #C#D#D#A#
	//  #########

	// TODO: FIXME: Implement me!
	return d.costs(costs)
}
