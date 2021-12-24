package assignment

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day24 struct{}

type d24Memory [4]int64

type d24State struct {
	instr int
	input int64
	zVal  int64
}

type d24StateResult struct {
	zVal int64
	line int
}

type d24Program struct {
	mem    d24Memory
	instr  []d24Instruction
	states map[d24State]d24StateResult
}

func (d *d24Program) execute(from int, input int64) (zVal int64, line int) {
	d.mem[0] = input
	state := d24State{
		instr: from,
		input: input,
		zVal:  d.mem[3],
	}

	res, ok := d.states[state]
	if ok {
		d.mem[3] = res.zVal
		return res.zVal, res.line
	}

	for line = from; line < len(d.instr); line++ {
		if d.instr[line].operation == d24Inp {
			break
		}
		d.instr[line].execute(&d.mem)
	}
	d.states[state] = d24StateResult{
		zVal: d.mem[3],
		line: line,
	}
	return d.mem[3], line
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

		op.bVal, err = strconv.ParseInt(bStr, 10, 32)
		if err != nil {
			op.bVal = d24BNotSet
			op.b = d.varToIdx(bStr)
		}

		instr = append(instr, op)
	}
	return instr
}

var d24Pow10s = [d24Digits + 1]int64{}

func init() {
	for i := range d24Pow10s {
		d24Pow10s[i] = d24Pow10(i)
	}
}

func d24Pow10(n int) int64 {
	res := int64(1)
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

func (d *Day24) numAt(num int64, pos int) int8 {
	// https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number/46755013
	n := num % d24Pow10s[pos]
	res := n / d24Pow10s[pos-1]
	return int8(res)
}

type d24Input [d24Digits]int8

func (d *Day24) SolveI(input string) int64 {
	instr := d.getInstructions(input)

	min, err := strconv.ParseInt(strings.Repeat("1", d24Digits), 10, 64)
	CheckErr(err)
	max, err := strconv.ParseInt(strings.Repeat("9", d24Digits), 10, 64)
	CheckErr(err)

	p := d24Program{
		instr:  instr,
		states: make(map[d24State]d24StateResult, 1024*1024),
	}

	in := d24Input{}

	for num := max; num >= min; num-- {
		p.mem = d24Memory{}
		pos := d24Digits - 1
		for i := d24Digits - 1; i >= 0; i-- {
			in[i] = d.numAt(num, i+1)
		}

		line := 1 // skip first input
		var zVal int64
		for line < len(instr) {
			zVal, line = p.execute(line, int64(in[pos]))
			line++ // skip input instruction
			pos--
		}

		if zVal == 0 {
			return num
		}
	}
	panic("no result")
}

func (d *Day24) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}
