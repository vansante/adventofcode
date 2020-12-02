package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
		if line != "" {
			input = append(input, line)
		}
	}
	return input
}

type password struct {
	min      int
	max      int
	letter   string
	password string
}

func (p password) valid1() bool {
	count := strings.Count(p.password, p.letter)
	return count >= p.min && count <= p.max
}

func (p password) valid2() bool {
	var firstOK, secondOK bool
	if len(p.password) >= p.min {
		firstOK = p.password[p.min-1:p.min] == p.letter
	}
	if len(p.password) >= p.max {
		secondOK = p.password[p.max-1:p.max] == p.letter
	}
	return (firstOK && !secondOK) || (secondOK && !firstOK)
}

func passwordFromLine(str string) password {
	var p password
	n, err := fmt.Sscanf(str, "%d-%d %s%s", &p.min, &p.max, &p.letter, &p.password)
	if n != 4 || err != nil {
		panic(fmt.Sprintf("n: %d: %v || %s", n, err, str))
	}
	p.letter = p.letter[:1]
	return p
}

func processLines(lines []string) []password {
	var passwords []password
	for i := range lines {
		passwords = append(passwords, passwordFromLine(lines[i]))
	}
	return passwords
}

func main() {
	wd, _ := os.Getwd()
	passwords := processLines(retrieveInputLines(filepath.Join(wd, "02/input.txt")))

	valid1 := 0
	valid2 := 0
	for i := range passwords {
		if passwords[i].valid1() {
			valid1++
		}
		if passwords[i].valid2() {
			valid2++
		}
	}
	fmt.Printf("Valid passwords (1): %d / %d\n", valid1, len(passwords))
	fmt.Printf("Valid passwords (2): %d / %d\n", valid2, len(passwords))
}
