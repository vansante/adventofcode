package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day05 struct{}

type d05Stack []string

func (s d05Stack) reverse() d05Stack {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

func (d *Day05) getStacks(input string) ([]d05Stack, []string) {
	lines := strings.Split(input, "\n")

	stacks := make([]d05Stack, 10)

	highest := 0
	for i, line := range lines {
		if line == "" {
			s := stacks[:highest+1]
			for i := range s {
				s[i] = s[i].reverse()
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
			for i := 0; i < len(crates)/2; i++ {
				crates[i], crates[len(crates)-1-i] = crates[len(crates)-1-i], crates[i]
			}
		}

		stacks[to-1] = append(stacks[to-1], crates...)
	}
}

func (d *Day05) SolveI(input string) int64 {
	stacks, lines := d.getStacks(input)
	d.doMoves(stacks, lines, true)

	fmt.Println()
	for i := range stacks {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()

	return 0
}

func (d *Day05) SolveII(input string) int64 {
	stacks, lines := d.getStacks(input)
	d.doMoves(stacks, lines, false)

	fmt.Println()
	for i := range stacks {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()

	return 0
}
