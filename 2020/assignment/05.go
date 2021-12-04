package assignment

import (
	"log"
	"strings"
)

type Day05 struct{}

func (d *Day05) retrieveTickets(input string) []d05Ticket {
	split := strings.Split(input, "\n")

	var tickets []d05Ticket
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		if len(line) != 10 {
			log.Panicf("[%s] invalid ticket", line)
		}
		tickets = append(tickets, strings.Split(line, ""))
	}
	return tickets
}

type d05Ticket []string

func (t d05Ticket) findRow() int {
	directives := t[:7]
	upper := 127
	lower := 0
	current := 64
	for i := range directives {
		switch directives[i] {
		case "F":
			upper -= current
		case "B":
			lower += current
		default:
			log.Panicf("[%s] invalid row id: %s", t, directives[i])
		}
		current /= 2
	}
	return lower
}

func (t d05Ticket) findColumn() int {
	directives := t[7:]
	upper := 7
	lower := 0
	current := 4
	for i := range directives {
		switch directives[i] {
		case "L":
			upper -= current
		case "R":
			lower += current
		default:
			log.Panicf("[%s] invalid column id: %s", t, directives[i])
		}
		current /= 2
	}
	return lower
}

func (t d05Ticket) id() int {
	return t.findRow()*8 + t.findColumn()
}

type d05Row [8]bool
type d05Plane [128]d05Row

func (d *Day05) SolveI(input string) int64 {
	tickets := d.retrieveTickets(input)

	var p d05Plane
	highestID := 0

	for i := range tickets {
		p[tickets[i].findRow()][tickets[i].findColumn()] = true

		if tickets[i].id() > highestID {
			highestID = tickets[i].id()
		}
	}

	return int64(highestID)
}

func (d *Day05) SolveII(input string) int64 {
	tickets := d.retrieveTickets(input)

	var p d05Plane
	for i := range tickets {
		p[tickets[i].findRow()][tickets[i].findColumn()] = true
	}

	for i := range p {
		if i <= 4 || i == 127 {
			continue
		}
		for j := range p[i] {
			if !p[i][j] {
				return int64(i*8 + j)
			}
		}
	}

	panic("no result")
}
