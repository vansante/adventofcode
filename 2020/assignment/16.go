package assignment

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day16 struct{}

type d16Property struct {
	name   string
	ranges []d16NumRange
}

type d16NumRange struct {
	min int
	max int
}

func (d *Day16) readProperties(lines []string) (remainder []string, props []d16Property) {
	for i := range lines {
		if lines[i] == "your ticket:" {
			return lines[i:], props
		}

		pieces := strings.Split(lines[i], ": ")
		if len(pieces) != 2 {
			panic(lines[i])
		}
		p := d16Property{
			name: strings.ReplaceAll(pieces[0], " ", "-"),
		}

		ranges := strings.Split(pieces[1], " or ")
		for j := range ranges {
			r := d16NumRange{}
			n, err := fmt.Sscanf(strings.TrimSpace(ranges[j]), "%d-%d", &r.min, &r.max)
			if n != 2 || err != nil {
				log.Panicf("[%s] %d matches, err: %v", ranges[j], n, err)
			}
			p.ranges = append(p.ranges, r)
		}
		props = append(props, p)
	}
	panic("no your ticket line")
}

func (d *Day16) readTickets(lines []string) (mine d16Ticket, nearby []d16Ticket) {
	isMine := true
	for i := range lines {
		if lines[i] == "your ticket:" {
			continue
		}
		if lines[i] == "nearby tickets:" {
			isMine = false
			continue
		}

		var t d16Ticket
		ticketNums := strings.Split(lines[i], ",")
		for j := range ticketNums {
			num, err := strconv.ParseInt(ticketNums[j], 10, 32)
			if err != nil {
				panic(err)
			}
			t = append(t, int(num))
		}

		if isMine {
			mine = t
		} else {
			nearby = append(nearby, t)
		}
	}
	return mine, nearby
}

func (p d16Property) valid(num int) bool {
	for _, r := range p.ranges {
		if num >= r.min && num <= r.max {
			return true
		}
	}
	return false
}

type d16Ticket []int

func (t d16Ticket) errorRate(props []d16Property) (rate int, validFor map[int][]string) {
	validFor = make(map[int][]string)
	for i, num := range t {
		var valid []string
		for _, p := range props {
			if p.valid(num) {
				valid = append(valid, p.name)
			}
		}
		if len(valid) == 0 {
			rate += num
			continue
		}
		validFor[i] = valid
	}
	return rate, validFor
}

func (d *Day16) findErrorRate(props []d16Property, tickets []d16Ticket) int {
	var rate int
	for i := range tickets {
		fieldRate, _ := tickets[i].errorRate(props)
		rate += fieldRate
	}
	return rate
}

func (d *Day16) filterInvalidTickets(props []d16Property, tickets []d16Ticket) (valid []d16Ticket, fieldIndex map[string]int) {
	validFor := make(map[int][]string)
	for i := range tickets {
		rate, fields := tickets[i].errorRate(props)
		if rate > 0 {
			continue // none valid
		}

		valid = append(valid, tickets[i])
		if i == 0 {
			validFor = fields
			continue
		}
		validFor = d.intersect(validFor, fields)
	}

	fieldIndex = make(map[string]int)
	for i := range validFor {
		if len(validFor[i]) != 1 {
			log.Panicf("invalid amount of possible fields for %d: %v", i, validFor[i])
		}
		fieldIndex[validFor[i][0]] = i
	}

	return valid, fieldIndex
}

func (d *Day16) intersect(a, b map[int][]string) map[int][]string {
	result := make(map[int][]string)
	for i := range a {
		for j := range a[i] {
			if len(b[i]) == 0 || d.contains(b[i], a[i][j]) {
				result[i] = append(result[i], a[i][j])
			}
		}
	}
	for i := range b {
		for j := range b[i] {
			if len(a[i]) == 0 {
				result[i] = append(result[i], b[i][j])
			}
		}
	}

	for {
		if !d.eliminateSingles(result) && !d.eliminateUniques(result) {
			break
		}
	} // Repeat until none remain

	return result
}

func (d *Day16) eliminateSingles(m map[int][]string) bool {
	removed := false
	for i := range m {
		if len(m[i]) != 1 {
			continue
		}

		for j := range m {
			if j == i {
				continue
			}
			m[j], removed = d.remove(m[j], m[i][0])
		}
		if removed {
			break
		}
	}
	return removed
}

func (d *Day16) eliminateUniques(m map[int][]string) bool {
	ctr := make(map[string]int)
	for i := range m {
		for f := range m[i] {
			ctr[m[i][f]]++
		}
	}

	for f := range ctr {
		if ctr[f] != 1 {
			continue
		}

		for i := range m {
			if d.contains(m[i], f) && len(m[i]) != 1 {
				m[i] = []string{f}
				return true
			}
		}
	}
	return false
}

func (d *Day16) contains(strs []string, str string) bool {
	for i := range strs {
		if strs[i] == str {
			return true
		}
	}
	return false
}

func (d *Day16) remove(strs []string, str string) ([]string, bool) {
	nw := make([]string, 0, len(strs))
	for i := range strs {
		if strs[i] == str {
			continue
		}
		nw = append(nw, strs[i])
	}
	return nw, len(nw) != len(strs)
}

func (d *Day16) SolveI(input string) int64 {
	lines := SplitLines(input)
	lines, props := d.readProperties(lines)
	_, nearby := d.readTickets(lines)

	return int64(d.findErrorRate(props, nearby))
}

func (d *Day16) SolveII(input string) int64 {
	lines := SplitLines(input)
	lines, props := d.readProperties(lines)
	mine, nearby := d.readTickets(lines)

	_, fields := d.filterInvalidTickets(props, nearby)

	var multi int64
	for f := range fields {
		if strings.HasPrefix(f, "departure") {
			if multi == 0 {
				multi = int64(mine[fields[f]])
			} else {
				multi *= int64(mine[fields[f]])
			}
		}
	}
	return multi
}
