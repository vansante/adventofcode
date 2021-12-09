package assignment

import (
	"strconv"
	"strings"
)

type Day07 struct{}

type d07contain struct {
	bagType string
	amount  int
}

func (d *Day07) retrieveRules(lines []string) map[string][]d07contain {
	rules := make(map[string][]d07contain)
	for i := range lines {
		words := strings.Split(lines[i], " ")
		subject := words[0] + words[1]
		if words[4] == "no" {
			rules[subject] = nil
			continue
		}
		for wordPos := 4; wordPos < len(words); wordPos += 4 {
			amount, err := strconv.ParseInt(words[wordPos], 10, 32)
			CheckErr(err)
			contained := words[wordPos+1] + words[wordPos+2]
			rules[subject] = append(rules[subject], d07contain{
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

func (d *Day07) inverse(rules map[string][]d07contain) map[string][]string {
	out := make(map[string][]string)
	for key := range rules {
		for i := range rules[key] {
			bag := rules[key][i].bagType
			out[bag] = append(out[bag], key)
		}
	}
	return out
}

func (d *Day07) findParents(bagType string, invert map[string][]string, results map[string]int) {
	for i := range invert[bagType] {
		currentBag := invert[bagType][i]
		results[currentBag]++

		if len(invert[bagType]) == 0 {
			continue
		}

		d.findParents(currentBag, invert, results)
	}
}

func (d *Day07) findChildrenCount(bagType string, rules map[string][]d07contain) int {
	sum := 0
	bagRules := rules[bagType]
	for i := range bagRules {
		sum += bagRules[i].amount
		sum += bagRules[i].amount * d.findChildrenCount(bagRules[i].bagType, rules)
	}
	return sum
}

func (d *Day07) SolveI(input string) int64 {
	rules := d.retrieveRules(SplitLines(input))
	invert := d.inverse(rules)

	const target = "shinygold"

	canContain := make(map[string]int)
	d.findParents(target, invert, canContain)

	return int64(len(canContain))
}

func (d *Day07) SolveII(input string) int64 {
	rules := d.retrieveRules(SplitLines(input))

	const target = "shinygold"

	sum := d.findChildrenCount(target, rules)

	return int64(sum)
}
