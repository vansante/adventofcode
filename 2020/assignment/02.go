package assignment

import (
	"fmt"
	"strings"
)

type Day02 struct{}

type d01Password struct {
	min      int
	max      int
	letter   string
	password string
}

func (p d01Password) valid1() bool {
	count := strings.Count(p.password, p.letter)
	return count >= p.min && count <= p.max
}

func (p d01Password) valid2() bool {
	var firstOK, secondOK bool
	if len(p.password) >= p.min {
		firstOK = p.password[p.min-1:p.min] == p.letter
	}
	if len(p.password) >= p.max {
		secondOK = p.password[p.max-1:p.max] == p.letter
	}
	return (firstOK && !secondOK) || (secondOK && !firstOK)
}

func (d *Day02) passwordFromLine(str string) d01Password {
	var p d01Password
	n, err := fmt.Sscanf(str, "%d-%d %s%s", &p.min, &p.max, &p.letter, &p.password)
	if n != 4 || err != nil {
		panic(fmt.Sprintf("n: %d: %v || %s", n, err, str))
	}
	p.letter = p.letter[:1]
	return p
}

func (d *Day02) processLines(lines []string) []d01Password {
	var passwords []d01Password
	for i := range lines {
		passwords = append(passwords, d.passwordFromLine(lines[i]))
	}
	return passwords
}

func (d *Day02) SolveI(input string) int64 {
	passwords := d.processLines(SplitLines(input))
	var valid int64
	for i := range passwords {
		if passwords[i].valid1() {
			valid++
		}
	}
	return valid
}

func (d *Day02) SolveII(input string) int64 {
	passwords := d.processLines(SplitLines(input))
	var valid int64
	for i := range passwords {
		if passwords[i].valid2() {
			valid++
		}
	}
	return valid
}
