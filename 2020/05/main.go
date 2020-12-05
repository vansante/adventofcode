package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func retrieveTickets(file string) []ticket {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var tickets []ticket
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

type ticket []string

func (t ticket) findRow() int {
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

func (t ticket) findColumn() int {
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

func (t ticket) id() int {
	return t.findRow()*8 + t.findColumn()
}

type row [8]bool
type plane [128]row

func main() {
	wd, _ := os.Getwd()
	tickets := retrieveTickets(filepath.Join(wd, "05/input.txt"))

	var p plane

	highestID := 0
	for i := range tickets {
		//fmt.Printf("%s: row: %d column: %d id: %d\n", tickets[i], tickets[i].findRow(), tickets[i].findColumn(), tickets[i].id())

		p[tickets[i].findRow()][tickets[i].findColumn()] = true

		if tickets[i].id() > highestID {
			highestID = tickets[i].id()
		}
	}

	fmt.Printf("The highest seat ID: %d\n\n", highestID)

	fmt.Println("Finding empty seats:")
	for i := range p {
		if i <= 4 || i == 127 {
			continue
		}
		for j := range p[i] {
			if !p[i][j] {
				fmt.Printf("empty seat at row: %d column: %d id: %d\n", i, j, i*8+j)
				return
			}
		}
	}
}
