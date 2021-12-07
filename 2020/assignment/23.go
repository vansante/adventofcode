package assignment

import (
	"strconv"
	"strings"
)

type Day23 struct{}

type d23Circle struct {
	cups     []int // CupLabel -> Next CupLabel
	min, max int
	pickup   [3]int
}

func (c *d23Circle) Result() string {
	v := ""
	next := c.cups[1]
	for {
		v += strconv.Itoa(next)
		next = c.cups[next]
		if next == 1 {
			return v
		}
	}
}

func (c *d23Circle) ResultInt() int64 {
	num, err := strconv.ParseInt(c.Result(), 10, 64)
	CheckErr(err)
	return num
}

func (c *d23Circle) CupsAfter1(n int) []int {
	next := c.cups[1]
	cups := make([]int, n)
	for i := 0; i < n; i++ {
		cups[i] = next
		next = c.cups[next]
	}
	return cups
}

func (c *d23Circle) Move() {
	currentCup := c.cups[0]
	c.pickup[0] = c.cups[currentCup]
	c.pickup[1] = c.cups[c.pickup[0]]
	c.pickup[2] = c.cups[c.pickup[1]]

	c.cups[currentCup] = c.cups[c.pickup[2]]

	dest := currentCup - 1
	for {
		if dest < c.min {
			dest = c.max
		}
		if !IntsContains(c.pickup[:], dest) {
			break
		}
		dest -= 1
	}

	c.cups[dest], c.cups[c.pickup[2]] = c.pickup[0], c.cups[dest]
	c.cups[0] = c.cups[c.cups[0]]
}

func (d *Day23) GetCircle(cups []int) *d23Circle {
	ci := &d23Circle{
		cups: make([]int, len(cups)+1),
		min:  IntsMin(cups),
		max:  IntsMax(cups),
	}

	for i := range cups {
		if i >= len(cups)-1 {
			break
		}
		ci.cups[cups[i]] = cups[i+1]
	}
	ci.cups[cups[len(cups)-1]] = cups[0] // Make it loop around
	ci.cups[0] = cups[0]                 // The zero index is a pointer to the current
	return ci
}

func (d *Day23) SolveI(input string) int64 {
	cups := MakeInts(strings.Split(input, ""))
	circle := d.GetCircle(cups)

	const rounds = 100
	for i := 0; i < rounds; i++ {
		circle.Move()
	}
	return circle.ResultInt()
}

func (d *Day23) SolveII(input string) int64 {
	const cupCount = 1_000_000
	const rounds = 10_000_000

	cups := MakeInts(strings.Split(input, ""))
	extCups := make([]int, cupCount)
	copy(extCups, cups)
	for i := len(cups); i < cupCount; i++ {
		extCups[i] = i + 1
	}

	circle := d.GetCircle(extCups)
	for i := 0; i < rounds; i++ {
		circle.Move()
	}

	cups = circle.CupsAfter1(2)
	return int64(cups[0]) * int64(cups[1])
}
