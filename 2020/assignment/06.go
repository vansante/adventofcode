package assignment

import (
	"strings"
)

type Day06 struct{}

type d06Answers string

type d06Groups struct {
	answers []d06Answers
	sum     map[string]int
}

func (d *Day06) retrieveGroups(input string) []d06Groups {
	split := strings.Split(input, "\n\n")

	var groups []d06Groups
	for i := range split {
		g := d06Groups{
			sum: make(map[string]int),
		}
		parts := strings.Split(split[i], "\n")
		for i := range parts {
			data := strings.TrimSpace(parts[i])
			if data == "" {
				continue
			}
			g.answers = append(g.answers, d06Answers(data))
		}
		if len(g.answers) > 0 {
			g.summarize()
			groups = append(groups, g)
		}
	}
	return groups
}

func (g *d06Groups) summarize() {
	for i := range g.answers {
		a := g.answers[i]
		for _, answer := range strings.Split(string(a), "") {
			g.sum[answer]++
		}
	}
}

func (d *Day06) SolveI(input string) int64 {
	groups := d.retrieveGroups(input)

	var yes int64
	for i := range groups {
		for range groups[i].sum {
			yes++
		}
	}

	return yes
}

func (d *Day06) SolveII(input string) int64 {
	groups := d.retrieveGroups(input)

	var allYes int64
	for i := range groups {
		for _, sum := range groups[i].sum {
			if sum == len(groups[i].answers) {
				allYes++
			}
		}
	}
	return allYes
}
