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
		// <generator:add:days>
	}
)

func main() {
	fmt.Printf("Usage: %s <dayNumber>\n\n", os.Args[0])

	if len(os.Args) < 2 {
		panic("please provide day number argument")
	}

	dayNum, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("error parsing day number [%s]: %v", os.Args[1], err))
	}

	day, ok := days[int(dayNum)]
	if !ok {
		panic(fmt.Sprintf("day %d not found", dayNum))
	}

	inputs := findInputs(int(dayNum))

	separator := strings.Repeat("#", 12)
	for _, in := range inputs {
		fmt.Printf("%s %d Day %02d - Part I - %s %s\n", separator, 2022, dayNum, in.name, separator)
		start := time.Now()
		resultI := day.SolveI(in.content)
		fmt.Printf("Found answer in %v:\n%v\n", time.Since(start), resultI)

		fmt.Println()

		fmt.Printf("%s %d Day %02d - Part II - %s %s\n", separator, 2022, dayNum, in.name, separator)
		start = time.Now()
		resultII := day.SolveII(in.content)
		fmt.Printf("Found answer in %v:\n%v\n", time.Since(start), resultII)

		fmt.Println()
	}
}

type input struct {
	name    string
	content string
}

func findInputs(dayNum int) []input {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := fmt.Sprintf("%s/%02d", wd, dayNum)
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	results := make([]input, 0, 8)
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".txt") {
			continue
		}
		contents, err := os.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			panic(err)
		}
		results = append(results, input{
			name:    strings.TrimSuffix(entry.Name(), ".txt"),
			content: string(contents),
		})
	}
	return results
}
