package assignment

import (
	"fmt"
	"math"
	"strings"
)

type Day14 struct{}

type d14Polymer struct {
	el []string
}

func (p *d14Polymer) insert(idx int, el string) {
	p.el = append(p.el[:idx], append([]string{el}, p.el[idx:]...)...)
}

func (p *d14Polymer) count() map[string]int64 {
	sums := make(map[string]int64)
	for _, el := range p.el {
		sums[el]++
	}
	return sums
}

func (p *d14Polymer) applyRules(rules []d14Rule) {
	for i := 0; i < len(p.el)-1; i++ {
		cur := p.el[i]
		nxt := p.el[i+1]
		for _, r := range rules {
			if r.a != cur || r.b != nxt {
				continue
			}
			// match, insert
			p.insert(i+1, r.ins)
			i++ // Increment i so we wont work on the inserted char
			break
		}
	}
}

type d14Rule struct {
	a, b, ins string
}

func (d *Day14) getInput(input string) (*d14Polymer, []d14Rule) {
	lines := SplitLines(input)

	p := &d14Polymer{
		el: strings.Split(lines[0], ""),
	}

	rules := make([]d14Rule, 0, len(lines))
	for _, line := range lines[1:] {
		chars := strings.Split(line, "")
		rules = append(rules, d14Rule{
			a:   chars[0],
			b:   chars[1],
			ins: chars[6],
		})
	}
	return p, rules
}

func (d *Day14) steps(p *d14Polymer, n int, rules []d14Rule) {
	for i := 0; i < n; i++ {
		p.applyRules(rules)
	}
}

func (d *Day14) SolveI(input string) int64 {
	p, rules := d.getInput(input)
	d.steps(p, 10, rules)

	sums := p.count()
	min := int64(math.MaxInt)
	max := int64(math.MinInt)
	for _, sum := range sums {
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min, max)
	return max - min
}

func (d *Day14) SolveII(input string) int64 {
	return 0
}
