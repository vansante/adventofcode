package assignment

import (
	"fmt"
	"sort"
	"strings"
)

type Day10 struct{}

type d10Chunk struct {
	Type     string
	Contains []*d10Chunk
}

func (c *d10Chunk) findChunk(char string) *d10Chunk {
	for i := len(c.Contains) - 1; i >= 0; i-- {
		if c.Contains[i].Type == char {
			return c.Contains[i]
		}
	}
	panic("not found")
}

var d20Opposite = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func (d *Day10) parseLine(line string, onInvalid func(string), onIncomplete func(string, []*d10Chunk)) *d10Chunk {
	var stack = []*d10Chunk{
		{
			Type:     "ROOT",
			Contains: make([]*d10Chunk, 0),
		},
	}
	for _, char := range strings.Split(line, "") {
		switch char {
		case "(", "[", "{", "<":
			nw := &d10Chunk{
				Type:     char,
				Contains: make([]*d10Chunk, 0),
			}
			cur := stack[len(stack)-1]
			cur.Contains = append(cur.Contains, nw)
			stack = append(stack, nw)
		case ")", "]", "}", ">":
			cur := stack[len(stack)-1]
			if d20Opposite[char] != cur.Type {
				fmt.Printf("%s | invalid char %s, expected %s\n", line, char, d20Opposite[cur.Type])
				if onInvalid != nil {
					onInvalid(d20Opposite[char])
				}
				return nil
			}
			if len(stack) == 1 {
				panic(fmt.Sprintf("invalid line: %s", line))
			}
			stack = stack[:len(stack)-1]

		default:
			panic("unknown character")
		}
	}
	if len(stack) != 1 {
		fmt.Printf("%s | incomplete\n", line)
		if onIncomplete != nil {
			onIncomplete(line, stack)
		}
	}
	return stack[0]
}

func (d *Day10) GetChunks(input string, onInvalid func(string), onIncomplete func(string, []*d10Chunk)) []*d10Chunk {
	lines := SplitLines(input)
	chnks := make([]*d10Chunk, 0, len(lines))
	for _, line := range lines {
		root := d.parseLine(line, onInvalid, onIncomplete)
		if root == nil {
			continue
		}
		chnks = append(chnks, root)
	}

	return chnks
}

func (d *Day10) SolveI(input string) int64 {
	sum := int64(0)
	d.GetChunks(input, func(s string) {
		switch s {
		case "(":
			sum += 3
		case "[":
			sum += 57
		case "{":
			sum += 1197
		case "<":
			sum += 25137
		}
	}, nil)
	return sum
}

func (d *Day10) SolveII(input string) int64 {
	scores := make([]int64, 0)
	d.GetChunks(input, nil, func(line string, stack []*d10Chunk) {
		score := int64(0)
		for i := len(stack) - 1; i > 0; i-- {
			chnk := stack[i]
			switch chnk.Type {
			case "(":
				score *= 5
				score += 1
			case "[":
				score *= 5
				score += 2
			case "{":
				score *= 5
				score += 3
			case "<":
				score *= 5
				score += 4
			}
		}
		//fmt.Printf("%s | %d short | %d score\n", line, len(stack)-1, score)
		scores = append(scores, score)
	})

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}
