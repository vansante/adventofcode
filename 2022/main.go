package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/vansante/adventofcode/2022/assignment"
)

var (
	days = map[int]assignment.Assignment{
		1:  &assignment.Day01{},
		2:  &assignment.Day02{},
		3:  &assignment.Day03{},
		4:  &assignment.Day04{},
		5:  &assignment.Day05{},
		6:  &assignment.Day06{},
		7:  &assignment.Day07{},
		8:  &assignment.Day08{},
		9:  &assignment.Day09{},
		10: &assignment.Day10{},
		11: &assignment.Day11{},
		12: &assignment.Day12{},
		13: &assignment.Day13{},
		14: &assignment.Day14{},
		15: &assignment.Day15{},
		16: &assignment.Day16{},
		17: &assignment.Day17{},
		18: &assignment.Day18{},
		19: &assignment.Day19{},
		// <generator:add:days>
	}
)

func main() {
	fmt.Printf("Usage: %s [<dayNumber>]\n\n", os.Args[0])

	if len(os.Args) <= 1 {
		for i := 1; i <= 25; i++ {
			if _, ok := days[i]; !ok {
				return
			}
			run(i)
		}
		return
	}

	dayNum, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("error parsing day number [%s]: %v", os.Args[1], err))
	}

	run(int(dayNum))
}

func run(dayNum int) {
	input := getInput(dayNum)

	separator := strings.Repeat("#", 12)
	fmt.Printf("%s %d Day %02d - Part I  %s\n", separator, 2022, dayNum, separator)
	start := time.Now()
	resultI := days[dayNum].SolveI(input)
	fmt.Printf("Found answer in %v:\n%v\n", time.Since(start), resultI)

	fmt.Println()

	fmt.Printf("%s %d Day %02d - Part II %s\n", separator, 2022, dayNum, separator)
	start = time.Now()
	resultII := days[dayNum].SolveII(input)
	fmt.Printf("Found answer in %v:\n%v\n", time.Since(start), resultII)

	fmt.Println()
}

func getInput(dayNum int) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := fmt.Sprintf("%s/%02d", wd, dayNum)
	contents, err := os.ReadFile(path.Join(dir, "input.txt"))
	if err != nil {
		panic(err)
	}

	return string(contents)
}
