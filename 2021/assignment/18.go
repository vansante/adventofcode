package assignment

import (
	"math"
	"strconv"
	"strings"
)

type Day18 struct{}

const (
	d18Open   = "["
	d18Close  = "]"
	d18Comma  = ","
	d18NotSet = math.MinInt
)

type d18Pair struct {
	lft, rgt       *d18Pair
	lftVal, rgtVal int
	parent         *d18Pair
}

func (d *Day18) parseLine(line string) *d18Pair {
	stack := make([]*d18Pair, 0, 128)
	var last *d18Pair
	for len(line) > 0 {
		switch line[:1] {
		case d18Open:
			nw := &d18Pair{lftVal: d18NotSet, rgtVal: d18NotSet}
			if len(stack) > 0 {
				cur := stack[len(stack)-1]
				if cur.lft == nil && cur.lftVal == d18NotSet {
					cur.lft = nw
				} else {
					cur.rgt = nw
				}
				nw.parent = cur
			}
			stack = append(stack, nw)

			line = line[1:]
			continue
		case d18Close:
			var cur *d18Pair
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if len(stack) == 0 {
				last = cur
			}
			if cur.lftVal == d18NotSet && cur.lft == nil {
				panic("left not set")
			}
			if cur.rgtVal == d18NotSet && cur.rgt == nil {
				panic("right not set")
			}

			line = line[1:]
			continue
		case d18Comma:
			line = line[1:]
			continue
		}

		commaIdx := strings.Index(line, ",")
		closeIdx := strings.Index(line, "]")
		if commaIdx < 0 && closeIdx < 0 {
			panic("wrong input")
		}

		if len(stack) == 0 {
			panic("empty stack")
		}

		cur := stack[len(stack)-1]
		if commaIdx > 0 && (commaIdx < closeIdx || closeIdx == -1) {
			val, err := strconv.ParseInt(line[:commaIdx], 10, 32)
			CheckErr(err)
			cur.lftVal = int(val)

			line = line[commaIdx:]
		} else if closeIdx > 0 && (closeIdx < commaIdx || commaIdx == -1) {
			val, err := strconv.ParseInt(line[:closeIdx], 10, 32)
			CheckErr(err)
			cur.rgtVal = int(val)

			line = line[closeIdx:]
		}
	}

	if len(stack) > 0 {
		panic("invalid input")
	}
	return last
}

func (d *Day18) getPairs(input string) []*d18Pair {
	lines := SplitLines(input)
	pairs := make([]*d18Pair, len(lines))
	for i, line := range lines {
		pairs[i] = d.parseLine(line)
	}
	return pairs
}

func (p *d18Pair) magnitude() int64 {
	mag := int64(0)
	if p.lft != nil {
		mag += 3 * p.lft.magnitude()
	} else {
		mag += int64(3 * p.lftVal)
	}

	if p.rgt != nil {
		mag += 2 * p.rgt.magnitude()
	} else {
		mag += int64(2 * p.rgtVal)
	}
	return mag
}

func (p *d18Pair) print() {
	print(d18Open)
	if p.lft != nil {
		p.lft.print()
	} else {
		print(p.lftVal)
	}
	print(d18Comma)
	if p.rgt != nil {
		p.rgt.print()
	} else {
		print(p.rgtVal)
	}
	print(d18Close)
}

func (p *d18Pair) add(pair *d18Pair) *d18Pair {
	nw := &d18Pair{
		lft:    p,
		rgt:    pair,
		lftVal: d18NotSet,
		rgtVal: d18NotSet,
	}
	pair.parent = nw
	p.parent = nw
	nw.reduce()
	return nw
}

func (p *d18Pair) reduce() {
	for {
		if p.explode() {
			continue
		}
		if p.split() {
			continue
		}
		break
	}
}

func (p *d18Pair) walkInOrder(depth int, walker func(p *d18Pair, depth int)) {
	if p.lft != nil {
		p.lft.walkInOrder(depth+1, walker)
	}
	walker(p, depth)
	if p.rgt != nil {
		p.rgt.walkInOrder(depth+1, walker)
	}
}

func (p *d18Pair) walkRight(depth int, walker func(p *d18Pair, depth int)) {
	if p.rgt != nil {
		p.rgt.walkRight(depth+1, walker)
	}
	walker(p, depth)
	if p.lft != nil {
		p.lft.walkRight(depth+1, walker)
	}
}

func (p *d18Pair) explode() bool {
	exploded := false
	p.walkInOrder(0, func(cur *d18Pair, depth int) {
		if exploded {
			return
		}
		if cur.lft != nil || cur.rgt != nil {
			return
		}
		if depth < 4 {
			return
		}

		cur.addLeft(cur.lftVal)
		cur.addRight(cur.rgtVal)
		cur.detach()

		exploded = true
		return
	})
	return exploded
}

func (p *d18Pair) addLeft(val int) {
	// the pair's left value is added to the first regular number to the left of the exploding pair (if any)
	cur := p.parent
	child := p
	for cur != nil {
		if cur.lft == child {
			child = cur
			cur = cur.parent
			continue
		}

		if cur.lft != nil {
			set := false
			cur.lft.walkRight(0, func(p *d18Pair, _ int) {
				if !set && p.rgt == nil {
					p.rgtVal += val
					set = true
				}
				if !set && p.lft == nil {
					p.lftVal += val
					set = true
				}
			})
			if set {
				return
			}
		}

		if cur.lft == nil {
			cur.lftVal += val
			return
		}
		child = cur
		cur = cur.parent
	}
}

func (p *d18Pair) addRight(val int) {
	// the pair's left value is added to the first regular number to the left of the exploding pair (if any)
	cur := p.parent
	child := p
	for cur != nil {
		if cur.rgt == nil {
			cur.rgtVal += val
			return
		}
		if cur.rgt == child {
			child = cur
			cur = cur.parent
			continue
		}

		if cur.rgt != nil {
			set := false
			cur.rgt.walkInOrder(0, func(p *d18Pair, _ int) {
				if !set && p.lft == nil {
					p.lftVal += val
					set = true
				}
			})
			if set {
				return
			}
		}

		if cur.lft == nil {
			cur.lftVal += val
			return
		}
		child = cur
		cur = cur.parent
	}
}

func (p *d18Pair) detach() {
	if p.parent == nil {
		panic("cannot detach without parent")
	}
	parent := p.parent
	p.parent = nil
	if parent.lft == p {
		parent.lft = nil
		parent.lftVal = 0
		return
	} else if parent.rgt == p {
		parent.rgt = nil
		parent.rgtVal = 0
		return
	}
	panic("no left or right set")
}

func (p *d18Pair) split() bool {
	split := false
	p.walkInOrder(0, func(cur *d18Pair, depth int) {
		if split {
			return
		}
		if cur.lftVal >= 10 {
			roundedDown := cur.lftVal / 2
			cur.lft = &d18Pair{
				lftVal: roundedDown,
				rgtVal: cur.lftVal - roundedDown,
				parent: cur,
			}
			cur.lftVal = d18NotSet
			split = true
			return
		}
		if cur.rgtVal >= 10 {
			roundedDown := cur.rgtVal / 2
			cur.rgt = &d18Pair{
				lftVal: roundedDown,
				rgtVal: cur.rgtVal - roundedDown,
				parent: cur,
			}
			cur.rgtVal = d18NotSet
			split = true
			return
		}
	})
	return split
}

func (d *Day18) SolveI(input string) int64 {
	pairs := d.getPairs(input)

	sum := pairs[0]
	for i, p := range pairs {
		if i == 0 {
			continue
		}

		sum = sum.add(p)
	}
	sum.print()
	println()
	return sum.magnitude()
}

func (d *Day18) SolveII(input string) int64 {
	lines := SplitLines(input)
	max := int64(math.MinInt)
	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}
			p1 := d.parseLine(lines[i])
			p2 := d.parseLine(lines[j])

			sum := p1.add(p2)
			mag := sum.magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	return max
}
