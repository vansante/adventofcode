package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	regex = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func retrieveInputLines(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	return input
}

type value struct {
	idx int64
	val int64
}

type program struct {
	xMask    int64
	zeroMask int64
	oneMask  int64
	values   []value
}

func getPrograms(lines []string) []program {
	var pros []program
	var pro program
	for i := range lines {
		if strings.HasPrefix(lines[i], "mask = ") {
			maskStr := strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "1", "0")
			maskStr = strings.ReplaceAll(maskStr, "X", "1")
			xMask, err := strconv.ParseInt(maskStr, 2, 64)
			if err != nil {
				panic(err)
			}

			maskStr = strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "0", "Z")
			maskStr = strings.ReplaceAll(maskStr, "X", "0")
			maskStr = strings.ReplaceAll(maskStr, "1", "0")
			maskStr = strings.ReplaceAll(maskStr, "Z", "1")
			zeroMask, err := strconv.ParseInt(maskStr, 2, 64)
			if err != nil {
				panic(err)
			}

			maskStr = strings.ReplaceAll(strings.TrimPrefix(lines[i], "mask = "), "X", "0")
			oneMask, err := strconv.ParseInt(maskStr, 2, 64)
			if err != nil {
				panic(err)
			}

			pros = append(pros, pro)
			pro = program{
				xMask:    xMask,
				zeroMask: zeroMask,
				oneMask:  oneMask,
			}
			continue
		}
		matches := regex.FindStringSubmatch(lines[i])
		if len(matches) != 3 {
			panic(fmt.Sprintf("[%s]: %d", lines[i], len(matches)))
		}
		idx, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			panic(err)
		}
		val, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			panic(err)
		}
		pro.values = append(pro.values, value{idx, val})
	}
	pros = append(pros, pro)
	return pros[1:]
}

func (p program) executePtI(memory map[int64]int64) {
	for i := range p.values {
		val := p.values[i]
		v := val.val & ^p.zeroMask
		memory[val.idx] = v | p.oneMask
	}
}

func (p program) executePtII(memory map[int64]int64) {
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

func main() {
	wd, _ := os.Getwd()
	lines := retrieveInputLines(filepath.Join(wd, "14/input.txt"))

	pros := getPrograms(lines)

	memory := make(map[int64]int64)
	for i := range pros {
		pros[i].executePtI(memory)
	}

	sum := int64(0)
	for idx := range memory {
		sum += memory[idx]
	}

	fmt.Printf("Part I sum: %d\n\n", sum)

	memory = make(map[int64]int64)
	for i := range pros {
		pros[i].executePtII(memory)
	}

	sum = int64(0)
	for idx := range memory {
		sum += memory[idx]
	}

	fmt.Printf("Part II sum: %d\n\n", sum)
}
