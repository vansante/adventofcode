package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day05 struct{}

type d05Stack []string

func (s d05Stack) read() string {
	return s[len(s)-1]
}

func (d *Day05) getStacks(input string) ([]d05Stack, []string) {
	lines := strings.Split(input, "\n")

	stacks := make([]d05Stack, 10)

	highest := 0
	for i, line := range lines {
		if line == "" {
			s := stacks[:highest+1]
			for i := range s {
				s[i] = util.SliceReverse(s[i])
			}
			return s, lines[i+1:]
		}
		expectCrate := false

		for j, char := range line {
			ch := string(char)
			switch ch {
			case "[":
				expectCrate = true
				continue
			case "]":
				expectCrate = false
				continue
			}
			if !expectCrate {
				continue
			}

			stack := j / 4
			highest = util.Max(stack, highest)
			stacks[stack] = append(stacks[stack], ch)
		}
	}
	panic("unexpected input")
}

func (d *Day05) doMoves(stacks []d05Stack, instructions []string, withReverse bool) {
	for _, instr := range instructions {
		if instr == "" {
			return
		}

		var amount, from, to int
		n, err := fmt.Sscanf(instr, "move %d from %d to %d", &amount, &from, &to)
		util.CheckErr(err)
		if n != 3 {
			panic("unknown matches")
		}

		var crates []string
		f := stacks[from-1]
		crates, stacks[from-1] = f[len(f)-amount:], f[:len(f)-amount]

		if withReverse {
			// reverse the crates
			crates = util.SliceReverse(crates)
		}

		stacks[to-1] = append(stacks[to-1], crates...)
	}
}

func (d *Day05) SolveI(input string) any {
	stacks, lines := d.getStacks(input)
	d.doMoves(stacks, lines, true)

	var s string
	for _, stack := range stacks {
		s += stack.read()
	}
	return s
}

func (d *Day05) SolveII(input string) any {
	stacks, lines := d.getStacks(input)
	d.doMoves(stacks, lines, false)

	var s string
	for _, stack := range stacks {
		s += stack.read()
	}
	return s
}
