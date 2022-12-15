package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day10 struct{}

type d10Instr struct {
	instr string
	value int64
}

func (d *Day10) getInstructions(input string) []d10Instr {
	lines := util.SplitLines(input)

	instr := make([]d10Instr, len(lines))
	for i, line := range lines {
		if line == "noop" {
			instr[i].instr = "noop"
			continue
		}

		n, err := fmt.Sscanf(line, "%s %d", &instr[i].instr, &instr[i].value)
		util.CheckErr(err)
		if n != 2 {
			panic("invalid format")
		}
	}
	return instr
}

type stateFn func(cycle int, xRegister int64)

func (d *Day10) eval(instr []d10Instr, state, afterState stateFn) int64 {
	cycle := 0
	xRegister := int64(1)

	state(cycle, xRegister)
	afterState(cycle, xRegister)

	for _, in := range instr {
		switch in.instr {
		case "noop":
			cycle++
			state(cycle, xRegister)
			afterState(cycle, xRegister)
		case "addx":
			cycle++
			state(cycle, xRegister)
			afterState(cycle, xRegister)

			cycle++
			state(cycle, xRegister)
			xRegister += in.value
			afterState(cycle, xRegister)
		}
	}
	return xRegister
}

func (d *Day10) SolveI(input string) any {
	instr := d.getInstructions(input)

	sum := int64(0)
	d.eval(instr, func(cycle int, xRegister int64) {
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			sum += int64(cycle) * xRegister
		}
	}, func(cycle int, xRegister int64) {})

	return sum
}

const d10Width, d10Height = 40, 6

type d10Screen [d10Width * d10Height]bool

func (s *d10Screen) get(x, y int) bool {
	return s[y*d10Width+x]
}

func (s *d10Screen) set(x, y int, on bool) {
	s[y*d10Width+x] = on
}

func (s *d10Screen) String() string {
	str := strings.Builder{}
	for y := 0; y < d10Height; y++ {
		for x := 0; x < d10Width; x++ {
			if s.get(x, y) {
				str.WriteString("▓")
				continue
			}
			str.WriteString("░")
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (s *d10Screen) print() {
	fmt.Println()
	print(s.String())
	fmt.Println()
}

func (d *Day10) SolveII(input string) any {
	instr := d.getInstructions(input)

	screen := d10Screen{}
	d.eval(instr, func(cycle int, xRegister int64) {}, func(cycle int, xRegister int64) {
		screenCycle := cycle % len(screen)

		if cycle == d10Width*d10Height {
			screen.print()
		}

		switch int64(cycle % d10Width) {
		case xRegister - 1, xRegister, xRegister + 1:
			screen[screenCycle] = true
		default:
			screen[screenCycle] = false
		}
	})

	return screen.String()
}
