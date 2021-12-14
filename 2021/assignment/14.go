package assignment

import (
	"math"
	"strings"
)

type Day14 struct{}

type d14Polymer struct {
	pairs  map[string]int64
	counts map[string]int64
}

func (p *d14Polymer) applyRules(rules []d14Rule) {
	nw := make(map[string]int64, len(p.pairs))

	for pair, freq := range p.pairs {
		if freq <= 0 {
			continue
		}
		nw[pair] = freq
	}
	for pair, freq := range p.pairs {
		if freq <= 0 {
			continue
		}
		for _, r := range rules {
			if pair != r.a+r.b {
				continue
			}
			// match, insert
			p.counts[r.ins] += freq
			nw[r.a+r.ins] += freq
			nw[r.ins+r.b] += freq
			nw[pair] -= freq
		}
	}
	p.pairs = nw
}

func (p *d14Polymer) minMax() (int64, int64) {
	min := int64(math.MaxInt)
	max := int64(math.MinInt)
	for _, sum := range p.counts {
		if sum <= 0 {
			continue
		}
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
	}
	return min, max
}

type d14Rule struct {
	a, b, ins string
}

func (d *Day14) getInput(input string) (*d14Polymer, []d14Rule) {
	lines := SplitLines(input)

	p := &d14Polymer{
		pairs:  make(map[string]int64),
		counts: make(map[string]int64),
	}
	line := lines[0]
	for i := 0; i < len(line)-1; i++ {
		p.pairs[line[i:i+2]]++
		p.counts[line[i:i+1]]++
	}
	p.counts[line[len(line)-1:]]++

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

	min, max := p.minMax()
	return max - min
}

func (d *Day14) SolveII(input string) int64 {
	p, rules := d.getInput(input)
	d.steps(p, 40, rules)

	min, max := p.minMax()
	return max - min
}
