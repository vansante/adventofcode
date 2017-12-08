package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"fmt"
)

type Comparator struct {
	execute func(regValue, compareValue int) bool
}

var comparators = make(map[string]*Comparator)

func init() {
	comparators[">"] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue > compareValue
		},
	}
	comparators[">="] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue >= compareValue
		},
	}
	comparators["<"] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue < compareValue
		},
	}
	comparators["<="] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue <= compareValue
		},
	}
	comparators["=="] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue == compareValue
		},
	}
	comparators["!="] = &Comparator{
		execute: func(regValue, compareValue int) bool {
			return regValue != compareValue
		},
	}
}

type Instruction struct {
	register        string
	increment       bool
	incrementValue  int
	comparator      Comparator
	compareRegister string
	compareValue    int
}

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc8/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	var instructions []Instruction
	for i := range lines {
		instruction := instructionFromLine(lines[i])
		instructions = append(instructions, *instruction)
	}

	highest := math.MinInt32
	registers := make(map[string]int)
	for i := range instructions {
		ins := instructions[i]

		if ins.comparator.execute(registers[ins.compareRegister], ins.compareValue) {
			if ins.increment {
				registers[ins.register] += ins.incrementValue
			} else {
				registers[ins.register] -= ins.incrementValue
			}

			if highest < registers[ins.register] {
				highest = registers[ins.register]
			}
		}
	}

	fmt.Printf("Highest while processing: %d\n", highest)

	highest = math.MinInt32
	for key := range registers {
		if highest < registers[key] {
			highest = registers[key]
		}
	}

	fmt.Printf("Highest: %d\n", highest)
}

func instructionFromLine(line string) (in *Instruction) {
	parts := strings.Split(line, " ")
	if len(parts) < 7 {
		panic("Not enough parts")
	}
	incVal, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	compVal, err := strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}

	in = &Instruction{
		register:        parts[0],
		increment:       parts[1] == "inc",
		incrementValue:  incVal,
		comparator:      *comparators[parts[5]],
		compareRegister: parts[4],
		compareValue:    compVal,
	}
	return
}
