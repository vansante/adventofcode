package main

import (
	"fmt"
	"os"
	"path"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"github.com/vansante/adventofcode/2022/util"

	"github.com/vansante/adventofcode/2022/assignment"
)

var (
	days = map[int]assignment.Assignment{
		1: &assignment.Day01{},
		// <generator:add:days>
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

	profile := util.SliceContains(os.Args, "<profile>")
	for _, in := range inputs {
		var cpuFile *os.File
		if profile {
			var err error
			cpuFile, err = os.CreateTemp(os.TempDir(), fmt.Sprintf("day_%02d_%s_cpu.pprof", dayNum, in.name))
			if err != nil {
				panic(err)
			}

			err = pprof.StartCPUProfile(cpuFile)
			if err != nil {
				_ = cpuFile.Close()
				panic(err)
			}
		}

		if !strings.HasPrefix(in.name, "2_") {
			fmt.Printf("Solving 2022 day %d first assignment with '%s'\n", dayNum, in.name)
			start := time.Now()
			resultI := day.SolveI(in.content)
			fmt.Printf("Solved first assignment: %d\n", resultI)
			fmt.Printf("Time taken: %v\n", time.Since(start))
		}

		if !strings.HasPrefix(in.name, "1_") {
			fmt.Printf("Solving 2022 day %d second assignment with '%s'\n", dayNum, in.name)
			start := time.Now()
			resultII := day.SolveII(in.content)
			fmt.Printf("Solved second assignment: %d\n", resultII)
			fmt.Printf("Time taken: %v\n", time.Since(start))
		}

		if profile {
			pprof.StopCPUProfile()
			fmt.Printf("Wrote cpu profile to %s\n", cpuFile.Name())
			_ = cpuFile.Close()
		}
		fmt.Println()
	}
}

type input struct {
	name    string
	content string
}

func findInputs(dayNum int, inputArg string) []input {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := fmt.Sprintf("%s/%02d", wd, dayNum)
	if inputArg != "" && inputArg[:1] != "<" {
		contents, err := os.ReadFile(path.Join(dir, inputArg+".txt"))
		if err != nil {
			panic(err)
		}
		return []input{{
			name:    inputArg,
			content: string(contents),
		}}
	}

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
