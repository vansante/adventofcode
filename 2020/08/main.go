package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func retrieveInstructions(file string) []operand {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []operand
	for i := range split {
		line := strings.ReplaceAll(strings.TrimSpace(split[i]), "+", "")
		if line == "" {
			continue
		}

		op := operand{}
		n, err := fmt.Sscanf(line, "%s %d", &op.operator, &op.argument)
		if err != nil || n != 2 {
			log.Panicf("[%s] error parsing operand: %v | %d", line, err, n)
		}
		input = append(input, op)
	}
	return input
}

type operand struct {
	operator string
	argument int
}

type computer struct {
	idx          int
	memory       []operand
	accumulator  int64
	instrCounter map[int]int
	haltOn       int
}

func (c *computer) run() {
	for c.idx < len(c.memory) {
		c.instrCounter[c.idx]++
		if c.haltOn > 1 && c.instrCounter[c.idx] > c.haltOn {
			break
		}

		op := c.memory[c.idx]

		switch op.operator {
		case "nop": // No operation
		case "jmp":
			c.idx += op.argument
			continue // Do not go past the c.idx++
		case "acc":
			c.accumulator += int64(op.argument)
		}
		c.idx++
	}
}

func main() {
	wd, _ := os.Getwd()
	instr := retrieveInstructions(filepath.Join(wd, "08/input.txt"))

	c := computer{
		memory:       instr,
		instrCounter: make(map[int]int),
		haltOn:       2,
	}
	c.run()

	fmt.Printf("Part I: The accumulator has: %d\n\n", c.accumulator)

	for i := range instr {
		prevOp := instr[i].operator
		switch instr[i].operator {
		case "jmp":
			instr[i].operator = "nop"
		case "nop":
			instr[i].operator = "jmp"
		}

		c := computer{
			memory:       instr,
			idx:          0,
			instrCounter: make(map[int]int),
			haltOn:       2,
		}
		c.run()
		if c.idx == len(c.memory) {
			fmt.Println("Found an instruction ending at the last operand")
			fmt.Printf("Part II: The accumulator has: %d\n\n", c.accumulator)
			break
		}
		instr[i].operator = prevOp
	}
}
