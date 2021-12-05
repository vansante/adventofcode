package assignment

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day14 struct{}

var d14regex = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

type d14Value struct {
	idx int64
	val int64
}

type d14Program struct {
	xMask    int64
	zeroMask int64
	oneMask  int64
	values   []d14Value
}

func (d *Day14) getPrograms(lines []string) []d14Program {
	var pros []d14Program
	var pro d14Program
	for i := range lines {
		if strings.HasPrefix(lines[i], "mask = ") {
			maskStr := strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "1", "0")
			maskStr = strings.ReplaceAll(maskStr, "X", "1")
			xMask, err := strconv.ParseInt(maskStr, 2, 64)
			CheckErr(err)

			maskStr = strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "0", "Z")
			maskStr = strings.ReplaceAll(maskStr, "X", "0")
			maskStr = strings.ReplaceAll(maskStr, "1", "0")
			maskStr = strings.ReplaceAll(maskStr, "Z", "1")
			zeroMask, err := strconv.ParseInt(maskStr, 2, 64)
			CheckErr(err)

			maskStr = strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "X", "0")
			oneMask, err := strconv.ParseInt(maskStr, 2, 64)
			CheckErr(err)

			pros = append(pros, pro)
			pro = d14Program{
				xMask:    xMask,
				zeroMask: zeroMask,
				oneMask:  oneMask,
			}
			continue
		}
		matches := d14regex.FindStringSubmatch(lines[i])
		if len(matches) != 3 {
			panic(fmt.Sprintf("[%s]: %d", lines[i], len(matches)))
		}
		idx, err := strconv.ParseInt(matches[1], 10, 64)
		CheckErr(err)
		val, err := strconv.ParseInt(matches[2], 10, 64)
		CheckErr(err)
		pro.values = append(pro.values, d14Value{idx, val})
	}
	pros = append(pros, pro)
	return pros[1:]
}

func (p d14Program) executePtI(memory map[int64]int64) {
	for i := range p.values {
		val := p.values[i]
		v := val.val & ^p.zeroMask
		memory[val.idx] = v | p.oneMask
	}
}

func (p d14Program) executePtII(memory map[int64]int64) {
	for i := range p.values {
		val := p.values[i]

		var indexes []int
		for i := 0; i < 36; i++ {
			if int64(math.Pow(2, float64(i)))&p.xMask == 0 {
				continue
			}
			indexes = append(indexes, i)
		}

		max := int64(math.Pow(2, float64(len(indexes))))
		for i := int64(0); i < max; i++ {
			addr := val.idx & ^p.xMask
			addr = addr | p.oneMask
			for j, idx := range indexes {
				pow := int64(math.Pow(2, float64(j)))
				if i&pow != 0 {
					addr += int64(math.Pow(2, float64(idx)))
				}
			}
			memory[addr] = val.val
		}
	}
}

func (d *Day14) SolveI(input string) int64 {
	pros := d.getPrograms(SplitLines(input))

	memory := make(map[int64]int64)
	for i := range pros {
		pros[i].executePtI(memory)
	}

	sum := int64(0)
	for idx := range memory {
		sum += memory[idx]
	}
	return sum
}

func (d *Day14) SolveII(input string) int64 {
	pros := d.getPrograms(SplitLines(input))

	memory := make(map[int64]int64)
	for i := range pros {
		pros[i].executePtII(memory)
	}

	sum := int64(0)
	for idx := range memory {
		sum += memory[idx]
	}
	return sum
}
