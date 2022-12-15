package assignment

import (
	"fmt"

	"github.com/vansante/adventofcode/2022/util"
)

type Day04 struct{}

type d04Section struct {
	start, end int64
}

func (s *d04Section) canMerge(other d04Section) *d04Section {
	if (s.start <= other.start && s.end >= other.start) ||
		(s.start <= other.end && s.end > other.end) {

		return &d04Section{
			start: util.Min(s.start, other.start),
			end:   util.Max(s.end, other.end),
		}
	}

	return nil
}

func (s *d04Section) contains(other d04Section) bool {
	return s.start <= other.start && s.end >= other.end
}

type d04Elve struct {
	s1, s2 d04Section
	merged *d04Section
}

func (e *d04Elve) checkMerge() {
	e.merged = e.s1.canMerge(e.s2)
	if e.merged == nil {
		e.merged = e.s2.canMerge(e.s1)
	}
}

func (d *Day04) getElves(input string) []d04Elve {
	lines := util.SplitLines(input)
	elves := make([]d04Elve, len(lines))
	for i, line := range lines {
		e := d04Elve{}
		n, err := fmt.Sscanf(line, "%d-%d,%d-%d", &e.s1.start, &e.s1.end, &e.s2.start, &e.s2.end)
		util.CheckErr(err)
		if n != 4 {
			panic(fmt.Sprintf("unexpected match count %d", n))
		}
		e.checkMerge()

		elves[i] = e
	}
	return elves
}

func (d *Day04) SolveI(input string) any {
	elves := d.getElves(input)

	var contained int64
	for _, e := range elves {
		if e.s1.contains(e.s2) || e.s2.contains(e.s1) {
			contained++
		}
	}
	return contained
}

func (d *Day04) SolveII(input string) any {
	elves := d.getElves(input)

	var overlap int64
	for _, e := range elves {
		if e.merged != nil {
			overlap++
		}
	}

	return overlap
}
