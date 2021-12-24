package assignment

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day24 struct{}

type d24Memory [4]int64

type d24Program struct {
	mem     d24Memory
	inputFn func() int64
	inputs  d24Inputs
	instr   []d24Instruction
}

func (d *d24Program) execute() {
	for _, i := range d.instr {
		i.execute(d)
	}
}

type d24Inputs []int64

type d24Instruction struct {
	operation string
	a, b      int8
	bVal      int64
}

func (i *d24Instruction) execute(pro *d24Program) {
	b := i.bVal
	if b == d24BNotSet {
		b = pro.mem[i.b]
	}
	switch i.operation {
	case "inp":
		if pro.inputFn != nil {
			pro.mem[i.a] = pro.inputFn()
			return
		}
		pro.mem[i.a], pro.inputs = pro.inputs[len(pro.inputs)-1], pro.inputs[:len(pro.inputs)-1]
		return
	case "add":
		pro.mem[i.a] += b
		return
	case "mul":
		pro.mem[i.a] *= b
		return
	case "div":
		pro.mem[i.a] /= b
		return
	case "mod":
		pro.mem[i.a] %= b
	case "eql":
		if pro.mem[i.a] == b {
			pro.mem[i.a] = 1
			return
		}
		pro.mem[i.a] = 0
		return
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
			aStr := ""
			n, err := fmt.Sscanf(line, "%s %s", &op.operation, &aStr)
			CheckErr(err)
			if n != 2 {
				panic("invalid input")
			}
			op.a = d.varToIdx(aStr)

			instr = append(instr, op)
			continue
		}
		var aStr, bStr string
		n, err := fmt.Sscanf(line, "%s %s %s", &op.operation, &aStr, &bStr)
		CheckErr(err)
		if n != 3 {
			panic("invalid input")
		}
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

func (d *Day24) pow10(n int) int64 {
	res := int64(10)
	for i := 0; i < n-1; i++ {
		res *= 10
	}
	return res
}

func (d *Day24) SolveI(input string) int64 {
	instr := d.getInstructions(input)

	min, err := strconv.ParseInt(strings.Repeat("1", d24Digits), 10, 64)
	CheckErr(err)
	max, err := strconv.ParseInt(strings.Repeat("9", d24Digits), 10, 64)
	CheckErr(err)

	p := d24Program{
		instr: instr,
	}
	zIndex := d.varToIdx("z")
	for num := max; num >= min; num-- {
		p.mem = d24Memory{}
		pos := d24Digits
		p.inputFn = func() int64 {
			// https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number/46755013
			n := num % d.pow10(pos)
			res := n / d.pow10(pos-1)
			pos--
			return res
		}
		p.execute()
		if p.mem[zIndex] != 0 {
			continue // invalid
		}
		return num
	}
	panic("no result")
}

func (d *Day24) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}
