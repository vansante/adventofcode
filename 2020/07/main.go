package main

import (
	"fmt"
	"io/ioutil"
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

type contain struct {
	bagType string
	amount  int
}

func retrieveRules(lines []string) map[string][]contain {
	rules := make(map[string][]contain)
	for i := range lines {
		words := strings.Split(lines[i], " ")
		subject := words[0] + words[1]
		if words[4] == "no" {
			rules[subject] = nil
			continue
		}
		for wordPos := 4; wordPos < len(words); wordPos += 4 {
			amount, err := strconv.ParseInt(words[wordPos], 10, 32)
			if err != nil {
				panic(err)
			}
			contained := words[wordPos+1] + words[wordPos+2]
			rules[subject] = append(rules[subject], contain{
				bagType: contained,
				amount:  int(amount),
			})
			if strings.Contains(words[wordPos+3], ".") {
				break
			}
		}
	}
	return rules
}

func inverse(rules map[string][]contain) map[string][]string {
	out := make(map[string][]string)
	for key := range rules {
		for i := range rules[key] {
			bag := rules[key][i].bagType
			out[bag] = append(out[bag], key)
		}
	}
	return out
}

func findParents(bagType string, invert map[string][]string, results map[string]int) {
	for i := range invert[bagType] {
		currentBag := invert[bagType][i]
		results[currentBag]++

		if len(invert[bagType]) == 0 {
			continue
		}

		findParents(currentBag, invert, results)
	}
}

func findChildrenCount(bagType string, rules map[string][]contain) int {
	sum := 0
	bagRules := rules[bagType]
	for i := range bagRules {
		sum += bagRules[i].amount
		sum += bagRules[i].amount * findChildrenCount(bagRules[i].bagType, rules)
	}
	return sum
}

func main() {
	wd, _ := os.Getwd()
	input := retrieveInputLines(filepath.Join(wd, "07/input.txt"))
	rules := retrieveRules(input)
	invert := inverse(rules)

	const target = "shinygold"

	canContain := make(map[string]int)
	findParents(target, invert, canContain)

	fmt.Printf("How many bag colors can eventually contain at least one shiny gold bag? %d\n\n", len(canContain))

	sum := findChildrenCount(target, rules)

	fmt.Printf("How many individual bags are required inside your single shiny gold bag? %d\n\n", sum)
}
