package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type Program struct {
	lastFreq   int64
	registers  map[string]int64
	inQueue    chan int64
	outQueue   chan int64
	valuesSent int
	waiting    bool
}

type Instruction struct {
	name      string
	registerA string
	registerB string
	aIsValue  bool
	bIsValue  bool
	valueA    int64
	valueB    int64
}

func (prog *Program) execute(ins *Instruction) (jump int64, freq int64) {
	switch ins.name {
	case "snd":
		if prog.outQueue != nil {
			prog.outQueue <- prog.registers[ins.registerA]
			prog.valuesSent++
		} else {
			freq = prog.registers[ins.registerA]
		}
	case "set":
		if ins.bIsValue {
			prog.registers[ins.registerA] = ins.valueB
		} else {
			prog.registers[ins.registerA] = prog.registers[ins.registerB]
		}
	case "add":
		if ins.bIsValue {
			prog.registers[ins.registerA] += ins.valueB
		} else {
			prog.registers[ins.registerA] += prog.registers[ins.registerB]
		}
	case "mul":
		if ins.bIsValue {
			prog.registers[ins.registerA] *= ins.valueB
		} else {
			prog.registers[ins.registerA] *= prog.registers[ins.registerB]
		}
	case "mod":
		if ins.bIsValue {
			prog.registers[ins.registerA] %= ins.valueB
		} else {
			prog.registers[ins.registerA] %= prog.registers[ins.registerB]
		}
	case "rcv":
		if prog.inQueue != nil {
			prog.waiting = true
			prog.registers[ins.registerA] = <-prog.inQueue
			prog.waiting = false
		} else {
			if prog.registers[ins.registerA] != 0 {
				freq = prog.lastFreq
			}
		}
	case "jgz":
		compareVal := prog.registers[ins.registerA]
		if ins.aIsValue {
			compareVal = ins.valueA
		}
		if compareVal > 0 {
			if ins.bIsValue {
				jump = ins.valueB
			} else {
				jump = prog.registers[ins.registerB]
			}
		}
	default:
		panic(ins.name)
	}
	return
}

func (prog *Program) run(instructions []Instruction) (frequency int64) {
	var position int
	for {
		ins := instructions[position]

		jump, newFreq := prog.execute(&ins)
		if newFreq != 0 {
			prog.lastFreq = newFreq
			if ins.name == "rcv" {
				break
			}
		}

		if jump == 0 {
			position++
		} else {
			position += int(jump)
		}

		if position < 0 || position >= len(instructions) {
			break
		}
	}
	frequency = prog.lastFreq
	return
}

func main() {
	input, err := ioutil.ReadFile("D:/go/src/adventofcode2017/aoc18/input.txt")
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	var instructions []Instruction
	for i := range lines {
		instructions = append(instructions, *instructionFromLine(lines[i]))
	}

	program := Program{
		registers: make(map[string]int64),
	}
	frequency := program.run(instructions)

	fmt.Printf("Frequency: %d\n", frequency)

	queue0 := make(chan int64, 100000)
	queue1 := make(chan int64, 100000)

	program0 := Program{
		registers: make(map[string]int64),
		inQueue:   queue1,
		outQueue:  queue0,
	}
	program1 := Program{
		registers: make(map[string]int64),
		inQueue:   queue0,
		outQueue:  queue1,
	}
	program1.registers["p"] = 1

	//wg := sync.WaitGroup{}
	//wg.Add(2)
	go func() {
		program0.run(instructions)
		//wg.Done()
	}()
	go func() {
		program1.run(instructions)
		//wg.Done()
	}()
	// FIXME: Find an elegant solution to detect they are both waiting
	time.Sleep(5 * time.Second)

	fmt.Printf("Prog 1 times sent: %d\n", program1.valuesSent)
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
		in.aIsValue = true
		in.bIsValue = true
		_, err := fmt.Sscanf(line, "%s %d %d", &in.name, &in.valueA, &in.valueB)
		if err != nil {
			_, err := fmt.Sscanf(line, "%s %s %d", &in.name, &in.registerA, &in.valueB)
			in.aIsValue = false
			if err != nil {
				_, err = fmt.Sscanf(line, "%s %s %s", &in.name, &in.registerA, &in.registerB)
				in.bIsValue = false
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return
}
