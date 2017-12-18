package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

var registers = make(map[string]int)
var instructions []Instruction

type Instruction struct {
	name      string
	registerA string
	registerB string
	useValue  bool
	value     int
}

func (ins *Instruction) execute(lastFreq int) (jump int, freq int) {
	switch ins.name {
	case "snd":
		freq = registers[ins.registerA]
	case "set":
		if ins.useValue {
			registers[ins.registerA] = ins.value
		} else {
			registers[ins.registerA] = registers[ins.registerB]
		}
	case "add":
		if ins.useValue {
			registers[ins.registerA] += ins.value
		} else {
			registers[ins.registerA] += registers[ins.registerB]
		}
	case "mul":
		if ins.useValue {
			registers[ins.registerA] *= ins.value
		} else {
			registers[ins.registerA] *= registers[ins.registerB]
		}
	case "mod":
		if ins.useValue {
			registers[ins.registerA] %= ins.value
		} else {
			registers[ins.registerA] %= registers[ins.registerB]
		}
	case "rcv":
		if registers[ins.registerA] != 0 {
			freq = lastFreq
		}
	case "jgz":
		if registers[ins.registerA] > 0 {
			if ins.useValue {
				jump = ins.value
			} else {
				jump = registers[ins.registerB]
			}
		}
	default:
		panic(ins.name)
	}
	return
}

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc18/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for i := range lines {
		instruction := instructionFromLine(lines[i])
		instructions = append(instructions, *instruction)
	}

	position := 0
	freq := 0
	for {
		ins := instructions[position]

		//fmt.Printf("%d %v\n%v\n", position, ins, registers)
		jump, newFreq := ins.execute(freq)
		if newFreq != 0 {
			freq = newFreq
		}

		if jump == 0 {
			position++
		} else {
			position += jump
		}

		if position < 0 || position >= len(instructions) {
			break
		}
	}

	fmt.Printf("Frequency: %d\n", freq)
}

func instructionFromLine(line string) (in *Instruction) {
	in = &Instruction{}

	switch line[:3] {
	case "snd", "rcv":
		_, err := fmt.Sscanf(line, "%s %s", &in.name, &in.registerA)
		if err != nil {
			panic(err)
		}
	default:
		in.useValue = true
		_, err := fmt.Sscanf(line, "%s %s %d", &in.name, &in.registerA, &in.value)
		if err != nil {
			_, err = fmt.Sscanf(line, "%s %s %s", &in.name, &in.registerA, &in.registerB)
			in.useValue = false
			if err != nil {
				panic(err)
			}
		}
	}
	return
}
