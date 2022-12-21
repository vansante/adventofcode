package assignment

import (
	"fmt"
	"log"
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

func (m *d21Monkey) calculate(a, b int64) int64 {
	switch m.operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("unknown operator")
}

func (m *d21Monkey) equalizeLeft(left, equalTo int64) int64 {
	switch m.operator {
	case "+":
		return equalTo - left
	case "-":
		return -equalTo + left
	case "*":
		return equalTo / left
	case "/":
		return left / equalTo
	}
	panic("unknown operator")
}

func (m *d21Monkey) equalizeRight(right, equalTo int64) int64 {
	switch m.operator {
	case "+":
		return equalTo - right
	case "-":
		return equalTo + right
	case "*":
		return equalTo / right
	case "/":
		return equalTo * right
	}
	panic("unknown operator")
}

func (d *Day21) getMonkeys(input string, skipHuman bool) *d21Monkeys {
	lines := util.SplitLines(input)

	m := &d21Monkeys{
		mp: make(map[string]*d21Monkey, len(lines)),
	}

	for _, line := range lines {
		split := strings.Split(line, ":")

		if skipHuman && split[0] == d21Human {
			continue
		}

		mon := &d21Monkey{
			name:       split[0],
			number:     math.MaxInt,
			monkeyNums: [2]int64{math.MaxInt, math.MaxInt},
		}
		m.mp[mon.name] = mon

		str := strings.TrimSpace(split[1])
		if unicode.IsDigit(rune(str[0])) {
			var err error
			mon.number, err = strconv.ParseInt(str, 10, 32)
			util.CheckErr(err)
			continue
		}

		spl := strings.Split(str, " ")
		mon.monkeys[0] = spl[0]
		mon.operator = spl[1]
		mon.monkeys[1] = spl[2]
	}
	return m
}

type d21Monkeys struct {
	mp map[string]*d21Monkey
}

func (m *d21Monkeys) calculate(name string) int64 {
	mon := m.mp[name]
	if mon.number != math.MaxInt {
		return mon.number
	}

	return mon.calculate(
		m.calculate(mon.monkeys[0]),
		m.calculate(mon.monkeys[1]),
	)
}

func (m *d21Monkeys) findMonkey(root, name string) bool {
	if root == name {
		return true
	}
	mon, ok := m.mp[root]
	if !ok {
		return false
	}
	if mon.name == name {
		return true
	}
	return m.findMonkey(mon.monkeys[0], name) || m.findMonkey(mon.monkeys[1], name)
}

func (m *d21Monkeys) mustEqual(root, human string, result int64) int64 {
	mon, ok := m.mp[root]
	if !ok {
		if root == human {
			return result
		}
		log.Panicf("root monkey %s not found", root)
	}

	if mon.number == math.MaxInt {
		if m.findMonkey(mon.monkeys[0], d21Human) {
			// Human is in the left branch
			right := m.calculate(mon.monkeys[1])
			return m.mustEqual(mon.monkeys[0], human, mon.equalizeRight(right, result))
		}
		// Human is in the right branch
		left := m.calculate(mon.monkeys[0])
		return m.mustEqual(mon.monkeys[1], human, mon.equalizeLeft(left, result))
	}

	panic(fmt.Sprintf("wrong monkey %s", mon.name))
}

func (m *d21Monkeys) findHumanValue() int64 {
	mon, ok := m.mp[d21Root]
	if !ok {
		panic("root monkey not found")
	}

	var result int64
	var branchRoot string
	if m.findMonkey(mon.monkeys[0], d21Human) {
		// Human is in the left branch
		result = m.calculate(mon.monkeys[1])
		branchRoot = mon.monkeys[0]
	} else {
		// Human is in the right branch
		result = m.calculate(mon.monkeys[0])
		branchRoot = mon.monkeys[1]
	}

	return m.mustEqual(branchRoot, d21Human, result)
}

func (d *Day21) SolveI(input string) any {
	monkeys := d.getMonkeys(input, false)
	return monkeys.calculate(d21Root)
}

func (d *Day21) SolveII(input string) any {
	monkeys := d.getMonkeys(input, true)
	return monkeys.findHumanValue()
}
