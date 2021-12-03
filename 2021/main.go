package main

import (
	"fmt"
	"github.com/vansante/adventofcode/2021/assignment"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	days = map[int]assignment.Assignment{
		1: &assignment.Day01{},
		2: &assignment.Day02{},
		3: &assignment.Day03{},
	}
)

func main() {
	fmt.Printf("Usage: %s <dayNumber> [<inputName>]\n\n", os.Args[0])

	if len(os.Args) < 2 {
		panic("please provide day number argument")
	}

	dayNum, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("error parsing day number [%s]: %v", os.Args[1], err))
	}

	inputArg := ""
	if len(os.Args) >= 3 {
		inputArg = os.Args[2]
	}

	day, ok := days[int(dayNum)]
	if !ok {
		panic(fmt.Sprintf("day %d not found", dayNum))
	}

	inputs := findInputs(int(dayNum), inputArg)

	for name, input := range inputs {
		fmt.Printf("Solving 2021 day %d first assignment with '%s'\n", dayNum, name)
		start := time.Now()
		resultI := day.SolveI(input)
		fmt.Printf("Solved first assignment: %d\n", resultI)
		fmt.Printf("Time taken: %v\n", time.Since(start))
		fmt.Println()

		fmt.Printf("Solving 2021 day %d second assignment with '%s'\n", dayNum, name)
		start = time.Now()
		resultII := day.SolveII(input)
		fmt.Printf("Solved second assignment: %d\n", resultII)
		fmt.Printf("Time taken: %v\n", time.Since(start))
		fmt.Println()
	}
}

func findInputs(dayNum int, inputArg string) map[string]string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := fmt.Sprintf("%s/%02d", wd, dayNum)
	if inputArg != "" {
		contents, err := os.ReadFile(path.Join(dir, inputArg+".txt"))
		if err != nil {
			panic(err)
		}
		return map[string]string{inputArg: string(contents)}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	results := make(map[string]string)
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".txt") {
			continue
		}
		contents, err := os.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			panic(err)
		}
		results[strings.TrimSuffix(entry.Name(), ".txt")] = string(contents)
	}
	return results
}
