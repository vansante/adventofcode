package assignment

import (
	"fmt"
	"strings"

	"github.com/vansante/adventofcode/2022/util"
)

type Day03 struct{}

type d03Item int

func (d *Day03) fromString(s string) d03Item {
	val := int([]rune(s)[0])
	if val < 97 { // Capital
		return d03Item(val - 64 + 26)
	}
	return d03Item(val - 96) // lowercase
}

type d03Rucksack struct {
	str  []string
	a, b []d03Item
}

func (r *d03Rucksack) findDuplicates() int64 {
	dupes := util.SliceIntersect(r.a, r.b)
	dupes = util.RemoveSliceDuplicates(dupes)
	if len(dupes) != 1 {
		panic(fmt.Sprintf("%d duplicates", len(dupes)))
	}
	return int64(util.SumSlice(dupes))
}

func (d *Day03) getRucksacks(input string) []d03Rucksack {
	lines := util.SplitLines(input)
	items := make([]d03Rucksack, len(lines))
	for i := range lines {
		line := lines[i]
		items[i].str = strings.Split(line, "")

		for _, s := range items[i].str[:len(line)/2] {
			items[i].a = append(items[i].a, d.fromString(s))
		}
		for _, s := range items[i].str[len(line)/2:] {
			items[i].b = append(items[i].b, d.fromString(s))
		}

	}
	return items
}

func (d *Day03) SolveI(input string) any {
	sacks := d.getRucksacks(input)

	var sum int64
	for _, sack := range sacks {
		sum += sack.findDuplicates()
	}
	return sum
}

func (d *Day03) SolveII(input string) any {
	sacks := d.getRucksacks(input)

	var sum int64
	for i := 0; i < len(sacks)-2; i += 3 {
		dupes := util.SliceIntersect(sacks[i].str, sacks[i+1].str)
		dupes = util.SliceIntersect(dupes, sacks[i+2].str)
		dupes = util.RemoveSliceDuplicates(dupes)
		if len(dupes) != 1 {
			panic(fmt.Sprintf("%d duplicates", len(dupes)))
		}
		sum += int64(d.fromString(dupes[0]))
	}
	return sum
}
