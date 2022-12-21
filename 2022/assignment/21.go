package assignment

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/vansante/adventofcode/2022/util"
)

type Day21 struct{}

const (
	d21Root  = "root"
	d21Human = "humn"
)

type d21Monkey struct {
	name       string
	number     int64
	monkeys    [2]string
	monkeyNums [2]int64
	operator   string
}

func (m *d21Monkey) reverseCalc() int64 {
	switch m.operator {
	case "+":
		return m.monkeyNums[0] + m.monkeyNums[1]
	case "-":
		return m.monkeyNums[0] - m.monkeyNums[1]
	case "*":
		return m.monkeyNums[0] / m.monkeyNums[1]
	case "/":
		return m.monkeyNums[0] * m.monkeyNums[1]
	}
	panic("unknown operator")
}

func (m *d21Monkey) calculate() int64 {
	switch m.operator {
	case "+":
		return m.monkeyNums[0] + m.monkeyNums[1]
	case "-":
		return m.monkeyNums[0] - m.monkeyNums[1]
	case "*":
		return m.monkeyNums[0] * m.monkeyNums[1]
	case "/":
		return m.monkeyNums[0] / m.monkeyNums[1]
	}
	panic("unknown operator")
}

func (d *Day21) getMonkeys(input string, skipHuman bool) *d21Monkeys {
	lines := util.SplitLines(input)

	m := &d21Monkeys{
		list: make([]*d21Monkey, len(lines)),
		mp:   make(map[string]*d21Monkey, len(lines)),
	}

	for i, line := range lines {
		split := strings.Split(line, ":")

		if skipHuman && split[0] == d21Human {
			continue
		}

		mon := &d21Monkey{
			name:       split[0],
			number:     math.MaxInt,
			monkeyNums: [2]int64{math.MaxInt, math.MaxInt},
		}
		m.list[i] = mon

		str := strings.TrimSpace(split[1])
		if unicode.IsDigit(rune(str[0])) {
			var err error
			mon.number, err = strconv.ParseInt(str, 10, 32)
			util.CheckErr(err)

			m.mp[mon.name] = mon
			continue
		}

		spl := strings.Split(str, " ")
		mon.monkeys[0] = spl[0]
		mon.operator = spl[1]
		mon.monkeys[1] = spl[2]
	}
	return m
}

func (m *d21Monkeys) solve(name string) int64 {
	for {
		for _, mon := range m.list {
			if mon.number != math.MaxInt {
				continue
			}

			num1, ok1 := m.mp[mon.monkeys[0]]
			num2, ok2 := m.mp[mon.monkeys[1]]
			if !ok1 || !ok2 {
				continue
			}
			mon.monkeyNums[0] = num1.number
			mon.monkeyNums[1] = num2.number

			mon.number = mon.calculate()
			m.mp[mon.name] = mon
		}

		val, ok := m.mp[name]
		if ok {
			return val.number
		}
	}
}

type d21Monkeys struct {
	list []*d21Monkey
	mp   map[string]*d21Monkey
}

func (d *Day21) equality() {

}

func (d *Day21) SolveI(input string) any {
	monkeys := d.getMonkeys(input, false)
	return monkeys.solve(d21Root)
}

func (d *Day21) SolveII(input string) any {
	//monkeys := d.getMonkeys(input, true)

	return "Not Implemented Yet"
}
