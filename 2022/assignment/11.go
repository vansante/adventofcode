package assignment

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day11 struct{}

type d11Monkey struct {
	id              int64
	items           []int64
	operation       string
	op              func(int64) int64
	testDivisor     int64
	trueMonkey      int64
	falseMonkey     int64
	inspectionCount int64
}

// Hardcode operations, for ease
func (m *d11Monkey) setOp() {
	switch m.operation {
	case "new = old * 17":
		m.op = func(i int64) int64 {
			return i * 17
		}
	case "new = old + 8":
		m.op = func(i int64) int64 {
			return i + 8
		}
	case "new = old + 6":
		m.op = func(i int64) int64 {
			return i + 6
		}
	case "new = old * 19":
		m.op = func(i int64) int64 {
			return i * 19
		}
	case "new = old + 7":
		m.op = func(i int64) int64 {
			return i + 7
		}
	case "new = old * old":
		m.op = func(i int64) int64 {
			return i * i
		}
	case "new = old + 1":
		m.op = func(i int64) int64 {
			return i + 1
		}
	case "new = old + 2":
		m.op = func(i int64) int64 {
			return i + 2
		}
	case "new = old + 3":
		m.op = func(i int64) int64 {
			return i + 3
		}
	default:
		log.Panicf("unknown operation: %s", m.operation)
	}
}

func (m *d11Monkey) receive(item int64) {
	m.items = append(m.items, item)
}

func (m *d11Monkey) inspect(item int64, monkeys *d11Monkeys, modulus int64) {
	m.inspectionCount++

	item = m.op(item)
	if modulus > 0 {
		item %= modulus
	} else {
		item = item / 3
	}

	if item%m.testDivisor == 0 {
		monkeys.get(m.trueMonkey).receive(item)
		return
	}
	monkeys.get(m.falseMonkey).receive(item)
}

func (m *d11Monkey) turn(monkeys *d11Monkeys, modulus int64) {
	for _, item := range m.items {
		m.inspect(item, monkeys, modulus)
	}
	m.items = m.items[:0] // Empty, all thrown
}

type d11Monkeys struct {
	ids []int64
	m   map[int64]*d11Monkey
}

func (m *d11Monkeys) add(monkey *d11Monkey) {
	m.ids = append(m.ids, monkey.id)
	m.m[monkey.id] = monkey
}

func (m *d11Monkeys) get(id int64) *d11Monkey {
	return m.m[id]
}

func (m *d11Monkeys) loop(walker func(m *d11Monkey)) {
	for _, id := range m.ids {
		walker(m.get(id))
	}
}

func (d *Day11) getMonkeys(input string) *d11Monkeys {
	lines := util.SplitLines(input)

	monkeys := &d11Monkeys{
		ids: make([]int64, 0, 128),
		m:   make(map[int64]*d11Monkey, 128),
	}
	var m *d11Monkey
	for _, line := range lines {
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "Monkey "):
			if m != nil {
				monkeys.add(m)
			}
			id, err := strconv.ParseInt(line[7:8], 10, 32)
			util.CheckErr(err)
			m = &d11Monkey{
				id:    id,
				items: make([]int64, 0, 100),
			}
		case strings.HasPrefix(line, "Starting items: "):
			m.items = util.ParseInt64s(strings.Split(line[16:], ", "))
		case strings.HasPrefix(line, "Operation: "):
			m.operation = line[11:]
			m.setOp()
		case strings.HasPrefix(line, "Test: divisible by "):
			var err error
			m.testDivisor, err = strconv.ParseInt(line[19:], 10, 32)
			util.CheckErr(err)
		case strings.HasPrefix(line, "If true: throw to monkey "):
			var err error
			m.trueMonkey, err = strconv.ParseInt(line[25:], 10, 32)
			util.CheckErr(err)
		case strings.HasPrefix(line, "If false: throw to monkey "):
			var err error
			m.falseMonkey, err = strconv.ParseInt(line[26:], 10, 32)
			util.CheckErr(err)
		default:
			log.Panicf("invalid line: %s", line)
		}
	}
	if m != nil {
		monkeys.add(m)
	}
	return monkeys
}

func (m *d11Monkeys) round(modulus int64) {
	m.loop(func(monkey *d11Monkey) {
		monkey.turn(m, modulus)
	})
}

func (m *d11Monkeys) sortByInspections() {
	sort.Slice(m.ids, func(i, j int) bool {
		return m.get(m.ids[i]).inspectionCount > m.get(m.ids[j]).inspectionCount
	})
}

func (d *Day11) SolveI(input string) any {
	monkeys := d.getMonkeys(input)

	for i := 0; i < 20; i++ {
		monkeys.round(-1)
	}

	monkeys.sortByInspections()
	return monkeys.get(monkeys.ids[0]).inspectionCount * monkeys.get(monkeys.ids[1]).inspectionCount
}

func (d *Day11) SolveII(input string) any {
	monkeys := d.getMonkeys(input)
	modulus := int64(1)
	monkeys.loop(func(m *d11Monkey) {
		modulus *= m.testDivisor
	})

	for i := 0; i < 10_000; i++ {
		monkeys.round(modulus)
	}

	monkeys.sortByInspections()
	return monkeys.get(monkeys.ids[0]).inspectionCount * monkeys.get(monkeys.ids[1]).inspectionCount
}
