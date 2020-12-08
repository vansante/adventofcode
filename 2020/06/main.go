package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type answers string
type group struct {
	answers []answers
	sum     map[string]int
}

func retrieveGroups(file string) []group {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n\n")

	var groups []group
	for i := range split {
		g := group{
			sum: make(map[string]int),
		}
		parts := strings.Split(split[i], "\n")
		for i := range parts {
			data := strings.TrimSpace(parts[i])
			if data == "" {
				continue
			}
			g.answers = append(g.answers, answers(data))
		}
		if len(g.answers) > 0 {
			g.summarize()
			groups = append(groups, g)
		}
	}
	return groups
}

func (g *group) summarize() {
	for i := range g.answers {
		a := g.answers[i]
		for _, answer := range strings.Split(string(a), "") {
			g.sum[answer]++
		}
	}
}

func main() {
	wd, _ := os.Getwd()
	groups := retrieveGroups(filepath.Join(wd, "06/input.txt"))

	yesAnswers := 0
	everyoneYesAnswers := 0
	for i := range groups {
		//fmt.Println(groups[i].sum)
		for _, sum := range groups[i].sum {
			yesAnswers++
			if sum == len(groups[i].answers) {
				everyoneYesAnswers++
			}
		}
	}

	fmt.Printf("Part I: Total questions with a yes answer: %d\n", yesAnswers)
	fmt.Printf("Part II: Total questions with everyone yes answer: %d\n", everyoneYesAnswers)

}
