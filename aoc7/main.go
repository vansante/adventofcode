package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Program struct {
	name     string
	weight   int
	parent   string
	children []string
}

var programs = make(map[string]*Program)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc7/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for i := range lines {
		program := programFromLine(lines[i])
		programs[program.name] = program
	}

	for name := range programs {
		for i := range programs[name].children {
			childName := programs[name].children[i]
			parentProg, ok := programs[childName]
			if !ok {
				panic("Not found")
			}
			parentProg.parent = name
		}
	}

	for name := range programs {
		if programs[name].parent == "" {
			fmt.Printf("Program with no parent: %s\n\n", name)
		}

		children := programs[name].children

		if ok, weights := consistentChildrenWeight(children); ok {
			consistent := true
			for i := range children {
				childProgram := programs[children[i]]
				if ok, _ := consistentChildrenWeight(childProgram.children); ok {
					consistent = false
				}
			}

			if consistent {
				//diff := weights[0] - weights[len(weights)-1]
				fmt.Printf("Inconsistent parent program: %#v, %#v \n\n", programs[name], weights)
				for i := range children {
					weight := programWeightRecursive(programs[children[i]])
					fmt.Printf("Child: %#v (%d)\n", programs[children[i]], weight)
				}
			}
		}
	}
}

func consistentChildrenWeight(children []string) (consistent bool, weights []int) {
	var all []int
	for i := range children {
		all = append(all, programWeightRecursive(programs[children[i]]))
	}

	sort.Ints(all)
	return len(all) > 2 && all[0] != all[len(all)-1], all
}

func programWeightRecursive(program *Program) (weight int) {
	weight += program.weight
	for i := range program.children {
		weight += programWeightRecursive(programs[program.children[i]])
	}
	return
}

func programFromLine(line string) (prg *Program) {
	prg = &Program{}

	parts := strings.Split(line, " -> ")
	if len(parts) < 1 {
		panic("No data")
	}

	nameParts := strings.Split(parts[0], " (")
	prg.name = nameParts[0]

	weightStr := strings.Trim(nameParts[1], ")\n")

	var err error
	prg.weight, err = strconv.Atoi(weightStr)
	if err != nil {
		fmt.Printf("Line: %s, nameParts: %#v", line, nameParts)
		panic(err)
	}

	if len(parts) > 1 {
		childrenParts := strings.Split(parts[1], ",")

		for i := range childrenParts {
			prg.children = append(prg.children, strings.TrimSpace(childrenParts[i]))
		}
	}
	return
}
