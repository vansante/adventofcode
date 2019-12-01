package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc16/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	instructionStrings := strings.Split(inputStr, ",")

	var instructions []Instruction
	for i := range instructionStrings {
		instructions = append(instructions, InstructionFromString(instructionStrings[i]))
	}

	var programs Programs
	programs.Init()
	programs.PerformInstructions(instructions, "")
	fmt.Println(programs.String())

	programs.Init()
	start := programs.String()

	numInstructions := 0
	for i := 0; ; i++ {
		repeat := programs.PerformInstructions(instructions, start)
		if repeat >= 0 {
			numInstructions = (i * len(instructions)) + repeat
			fmt.Printf("Found start position after %d instructions\n", numInstructions)
			break
		}
	}

	programs.Init()
	totalInstructions := 1000*1000*1000 * len(instructions)
	remainingInstructions := totalInstructions % numInstructions
	fmt.Printf("After discounting duplicate rounds, %d instructions left\n", remainingInstructions)

	for i := remainingInstructions; i > 0; {
		if i >= len(instructions) {
			programs.PerformInstructions(instructions, "")
		} else {
			programs.PerformInstructions(instructions[:i], "")
		}
		i -= len(instructions)
	}

	fmt.Println(programs.String())
}

type Programs struct {
	List  []string
	Index map[string]int
	Start int
}

func (p *Programs) Init() {
	p.List = make([]string, 16)
	p.Index = make(map[string]int, 16)
	p.Start = 0

	for i := 0; i < 16; i++ {
		letter := string(97+i)
		p.List[i] = letter
		p.Index[letter] = i
	}
}

func (p *Programs) PerformInstructions(instructions []Instruction, findSequence string) (repeat int) {
	for i := range instructions {
		p.perform(&instructions[i])

		if findSequence != "" && p.String() == findSequence {
			return i + 1
		}
	}
	return -1
}

func (p *Programs) perform(ins *Instruction) {
	list := p.List
	idx := p.Index

	switch ins.Type {
	case "s":
		p.Start = p.Start - ins.Length
		if p.Start < 0 {
			p.Start += len(p.List)
		}
	case "x":
		posA := (ins.PosA + p.Start) % len(list)
		posB := (ins.PosB + p.Start) % len(list)
		idx[list[posA]], idx[list[posB]] = idx[list[posB]], idx[list[posA]]
		list[posA], list[posB] = list[posB], list[posA]
	case "p":
		list[idx[ins.ProgramA]], list[idx[ins.ProgramB]] = list[idx[ins.ProgramB]], list[idx[ins.ProgramA]]
		idx[ins.ProgramA], idx[ins.ProgramB] = idx[ins.ProgramB], idx[ins.ProgramA]
	default:
		panic(ins)
	}
}

func (p *Programs) String() (str string) {
	for i := p.Start; i < p.Start + len(p.List); i++ {
		idx := i % len(p.List)
		str += p.List[idx]
	}
	return
}

func InstructionFromString(instStr string) (inst Instruction) {
	inst.Type = instStr[:1]

	switch inst.Type {
	case "s":
		num, err := strconv.Atoi(instStr[1:])
		if err != nil {
			panic(err)
		}
		inst.Length = num
	case "x":
		var err error
		splitIdx := strings.Index(instStr, "/")
		inst.PosA, err = strconv.Atoi(instStr[1:splitIdx])
		if err != nil {
			panic(err)
		}
		inst.PosB, err = strconv.Atoi(instStr[splitIdx+1:])
		if err != nil {
			panic(err)
		}
	case "p":
		splitIdx := strings.Index(instStr, "/")
		inst.ProgramA, inst.ProgramB = instStr[1:splitIdx], instStr[splitIdx+1:]
	default:
		panic(inst)
	}
	return
}

type Instruction struct {
	Type     string
	Length   int
	ProgramA string
	ProgramB string
	PosA     int
	PosB     int
}
