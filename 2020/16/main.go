package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func retrieveInputLines(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	return input
}

type property struct {
	name   string
	ranges []numRange
}

type numRange struct {
	min int
	max int
}

func readProperties(lines []string) (remainder []string, props []property) {
	for i := range lines {
		if lines[i] == "your ticket:" {
			return lines[i:], props
		}

		pieces := strings.Split(lines[i], ": ")
		if len(pieces) != 2 {
			panic(lines[i])
		}
		p := property{
			name: strings.ReplaceAll(pieces[0], " ", "-"),
		}

		ranges := strings.Split(pieces[1], " or ")
		for j := range ranges {
			r := numRange{}
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

func readTickets(lines []string) (mine ticket, nearby []ticket) {
	isMine := true
	for i := range lines {
		if lines[i] == "your ticket:" {
			continue
		}
		if lines[i] == "nearby tickets:" {
			isMine = false
			continue
		}

		var t ticket
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

func (p property) valid(num int) bool {
	for _, r := range p.ranges {
		if num >= r.min && num <= r.max {
			return true
		}
	}
	return false
}

type ticket []int

func (t ticket) errorRate(props []property) (rate int, validFor map[int][]string) {
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

func findErrorRate(props []property, tickets []ticket) int {
	var rate int
	for i := range tickets {
		fieldRate, _ := tickets[i].errorRate(props)
		rate += fieldRate
	}
	return rate
}

func filterInvalidTickets(props []property, tickets []ticket) (valid []ticket, fieldIndex map[string]int) {
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
		validFor = intersect(validFor, fields)
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

func intersect(a, b map[int][]string) map[int][]string {
	result := make(map[int][]string)
	for i := range a {
		for j := range a[i] {
			if len(b[i]) == 0 || contains(b[i], a[i][j]) {
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
		if !eliminateSingles(result) && !eliminateUniques(result) {
			break
		}
	} // Repeat until none remain

	return result
}

func eliminateSingles(m map[int][]string) bool {
	removed := false
	for i := range m {
		if len(m[i]) != 1 {
			continue
		}

		for j := range m {
			if j == i {
				continue
			}
			m[j], removed = remove(m[j], m[i][0])
		}
		if removed {
			break
		}
	}
	return removed
}

func eliminateUniques(m map[int][]string) bool {
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
			if contains(m[i], f) && len(m[i]) != 1 {
				m[i] = []string{f}
				return true
			}
		}
	}
	return false
}

func contains(strs []string, str string) bool {
	for i := range strs {
		if strs[i] == str {
			return true
		}
	}
	return false
}

func remove(strs []string, str string) ([]string, bool) {
	nw := make([]string, 0, len(strs))
	for i := range strs {
		if strs[i] == str {
			continue
		}
		nw = append(nw, strs[i])
	}
	return nw, len(nw) != len(strs)
}

func main() {
	wd, _ := os.Getwd()
	lines := retrieveInputLines(filepath.Join(wd, "16/input.txt"))

	lines, props := readProperties(lines)
	mine, nearby := readTickets(lines)

	fmt.Printf("Part I: Error rate is %d\n\n", findErrorRate(props, nearby))

	_, fields := filterInvalidTickets(props, nearby)

	fmt.Printf("Part II: The fields are:\n%v\n\n", fields)

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

	fmt.Printf("The fields with departure multiplied are: %d\n\n", multi)
}
