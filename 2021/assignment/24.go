package assignment

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Day24 struct{}

type d24Memory [4]int64

type d24Program struct {
	mem   d24Memory
	instr []d24Instruction
}

func (d *d24Program) execute(from int, input int64) int {
	d.mem[0] = input
	line := 0
	for line = from; line < len(d.instr); line++ {
		if d.instr[line].operation == d24Inp {
			break
		}
		d.instr[line].execute(&d.mem)
	}
	return line
}

type d24Operation int8

const (
	d24Inp d24Operation = iota
	d24Add
	d24Mul
	d24Div
	d24Mod
	d24Eql
)

type d24Instruction struct {
	operation d24Operation
	a, b      int8
	bVal      int64
}

func (i *d24Instruction) execute(mem *d24Memory) {
	b := i.bVal
	if b == d24BNotSet {
		b = mem[i.b]
	}
	switch i.operation {
	case d24Add:
		mem[i.a] += b
	case d24Mul:
		mem[i.a] *= b
	case d24Div:
		mem[i.a] /= b
	case d24Mod:
		mem[i.a] %= b
	case d24Eql:
		if mem[i.a] == b {
			mem[i.a] = 1
			return
		}
		mem[i.a] = 0
	default:
		panic("invalid instruction")
	}
}

func (d *Day24) varToIdx(char string) int8 {
	switch char {
	case "w":
		return 0
	case "x":
		return 1
	case "y":
		return 2
	case "z":
		return 3
	}
	panic("invalid var")
}

func (d *Day24) strToOp(str string) d24Operation {
	switch str {
	case "inp":
		return d24Inp
	case "add":
		return d24Add
	case "mul":
		return d24Mul
	case "div":
		return d24Div
	case "mod":
		return d24Mod
	case "eql":
		return d24Eql
	}
	panic("invalid var")
}

const (
	d24BNotSet = math.MinInt64
	d24Digits  = 14
)

func (d *Day24) getInstructions(input string) []d24Instruction {
	lines := SplitLines(input)

	instr := make([]d24Instruction, 0, len(input))
	for _, line := range lines {
		op := d24Instruction{}
		if line[:3] == "inp" {
			opStr := ""
			aStr := ""
			n, err := fmt.Sscanf(line, "%s %s", &opStr, &aStr)
			CheckErr(err)
			if n != 2 {
				panic("invalid input")
			}
			op.operation = d.strToOp(opStr)
			op.a = d.varToIdx(aStr)

			instr = append(instr, op)
			continue
		}
		var opStr, aStr, bStr string
		n, err := fmt.Sscanf(line, "%s %s %s", &opStr, &aStr, &bStr)
		CheckErr(err)
		if n != 3 {
			panic("invalid input")
		}

		op.operation = d.strToOp(opStr)
		op.a = d.varToIdx(aStr)

		bVal, err := strconv.ParseInt(bStr, 10, 32)
		if err == nil {
			op.bVal = bVal
		} else {
			op.bVal = d24BNotSet
			op.b = d.varToIdx(bStr)
		}

		instr = append(instr, op)
	}
	return instr
}

type d24State struct {
	prevNums int64
	zVal     int64
}

func (d *Day24) highestNumber(sl []d24State) []d24State {
	m := make(map[int64]int64, len(sl))
	for i := range sl {
		cur, ok := m[sl[i].zVal]
		if ok && cur > sl[i].prevNums {
			continue
		}
		m[sl[i].zVal] = sl[i].prevNums
	}

	nw := make([]d24State, 0, len(m))
	for zVal, prev := range m {
		nw = append(nw, d24State{prevNums: prev, zVal: zVal})
	}
	return nw
}

func (d *Day24) lowestNumber(sl []d24State) []d24State {
	m := make(map[int64]int64, len(sl))
	for i := range sl {
		cur, ok := m[sl[i].zVal]
		if ok && cur < sl[i].prevNums {
			continue
		}
		m[sl[i].zVal] = sl[i].prevNums
	}

	nw := make([]d24State, 0, len(m))
	for zVal, prev := range m {
		nw = append(nw, d24State{prevNums: prev, zVal: zVal})
	}
	return nw
}

func (d *Day24) simulateStates(input string, stateFilter func(sl []d24State) []d24State) []int64 {
	instr := d.getInstructions(input)

	p := d24Program{
		instr: instr,
	}

	mems := make([]d24State, 1, 9)
	mems[0] = d24State{prevNums: 0, zVal: 0}
	line := 1
	for line < len(p.instr) {
		var nwLine int
		nwMems := make([]d24State, 0, len(mems))
		for i := range mems {
			for n := 1; n <= 9; n++ {
				p.mem = d24Memory{0, 0, 0, mems[i].zVal}
				nwLine = p.execute(line, int64(n))

				nwMems = append(nwMems, d24State{
					prevNums: mems[i].prevNums*10 + int64(n),
					zVal:     p.mem[3],
				})
			}
		}
		line = nwLine + 1
		mems = stateFilter(nwMems)
		fmt.Printf("Simulating %d states\n", len(mems))
	}

	zeroVals := make([]int64, 0)
	for i := range mems {
		if mems[i].zVal == 0 {
			zeroVals = append(zeroVals, mems[i].prevNums)
		}
	}
	return zeroVals
}

func (d *Day24) SolveI(input string) int64 {
	zeroVals := d.simulateStates(input, d.highestNumber)

	sort.Slice(zeroVals, func(i, j int) bool {
		return zeroVals[i] > zeroVals[j]
	})

	return zeroVals[0]
}

func (d *Day24) SolveII(input string) int64 {
	zeroVals := d.simulateStates(input, d.lowestNumber)

	sort.Slice(zeroVals, func(i, j int) bool {
		return zeroVals[i] < zeroVals[j]
	})

	return zeroVals[0]
}
