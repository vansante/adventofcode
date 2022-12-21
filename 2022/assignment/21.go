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
	d21Root = "root"
)

type d21Monkey struct {
	name       string
	number     int64
	monkeys    [2]string
	monkeyNums [2]int64
	operator   string
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

func (d *Day21) getMonkeys(input string) (monkeys []*d21Monkey, nums map[string]int64) {
	lines := util.SplitLines(input)

	mp := make(map[string]int64)
	m := make([]*d21Monkey, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ":")
		m[i] = &d21Monkey{
			name:       split[0],
			number:     math.MaxInt,
			monkeyNums: [2]int64{math.MaxInt, math.MaxInt},
		}

		str := strings.TrimSpace(split[1])
		if unicode.IsDigit(rune(str[0])) {
			var err error
			m[i].number, err = strconv.ParseInt(str, 10, 32)
			util.CheckErr(err)

			mp[m[i].name] = m[i].number
			continue
		}

		spl := strings.Split(str, " ")
		m[i].monkeys[0] = spl[0]
		m[i].operator = spl[1]
		m[i].monkeys[1] = spl[2]
	}
	return m, mp
}

func (d *Day21) Solve(monkeys []*d21Monkey, nums map[string]int64) {
	for {
		for _, mon := range monkeys {
			if mon.number != math.MaxInt {
				continue
			}

			num1, ok1 := nums[mon.monkeys[0]]
			num2, ok2 := nums[mon.monkeys[1]]
			if !ok1 || !ok2 {
				continue
			}
			mon.monkeyNums[0] = num1
			mon.monkeyNums[1] = num2

			mon.number = mon.calculate()
			nums[mon.name] = mon.number
		}

		_, ok := nums[d21Root]
		if ok {
			break
		}
	}
}

func (d *Day21) SolveI(input string) any {
	monkeys, nums := d.getMonkeys(input)

	d.Solve(monkeys, nums)

	return nums[d21Root]
}

func (d *Day21) SolveII(input string) any {
	const identifier = "humn"

	return "Not Implemented Yet"
}
