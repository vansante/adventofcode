package assignment

import (
	"fmt"
	"log"
	"strings"
)

type Day08 struct{}

func (d *Day08) retrieveInstructions(in string) []d08Operand {
	split := strings.Split(in, "\n")

	var input []d08Operand
	for i := range split {
		line := strings.ReplaceAll(strings.TrimSpace(split[i]), "+", "")
		if line == "" {
			continue
		}

		op := d08Operand{}
		n, err := fmt.Sscanf(line, "%s %d", &op.operator, &op.argument)
		if err != nil || n != 2 {
			log.Panicf("[%s] error parsing operand: %v | %d", line, err, n)
		}
		input = append(input, op)
	}
	return input
}

type d08Operand struct {
	operator string
	argument int
}

type d08Computer struct {
	idx          int
	memory       []d08Operand
	accumulator  int64
	instrCounter map[int]int
	haltOn       int
}

func (c *d08Computer) run() {
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

func (d *Day08) SolveI(input string) int64 {
	instr := d.retrieveInstructions(input)

	c := d08Computer{
		memory:       instr,
		instrCounter: make(map[int]int),
		haltOn:       2,
	}
	c.run()

	return c.accumulator
}

func (d *Day08) SolveII(input string) int64 {
	instr := d.retrieveInstructions(input)

	for i := range instr {
		prevOp := instr[i].operator
		switch instr[i].operator {
		case "jmp":
			instr[i].operator = "nop"
		case "nop":
			instr[i].operator = "jmp"
		}

		c := d08Computer{
			memory:       instr,
			idx:          0,
			instrCounter: make(map[int]int),
			haltOn:       2,
		}
		c.run()
		if c.idx == len(c.memory) {
			return c.accumulator
		}
		instr[i].operator = prevOp
	}

	panic("no result")
}
