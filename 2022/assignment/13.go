package assignment

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/vansante/adventofcode/2022/util"
)

type Day13 struct{}

func (d *Day13) getPairs(input string) []d13Pair {
	lines := strings.Split(input, "\n")

	p := &d13Pair{}
	pairs := make([]d13Pair, 0, 128)
	for i, line := range lines {
		if line == "" {
			pairs = append(pairs, *p)
			p = &d13Pair{}
			continue
		}

		switch i % 3 {
		case 0: // left packet
			p.left = d.parsePacket(line)
		case 1: // right packet
			p.right = d.parsePacket(line)
		default:
			panic("invalid line")
		}
	}
	return pairs
}

type d13Pair struct {
	left, right *d13Packet
}

func (p *d13Pair) print() {
	fmt.Println("L:", p.left.String())
	fmt.Println("R:", p.right.String())
	fmt.Println()
}

func (p *d13Pair) inOrder() bool {
	score := p.left.inOrder(p.right)
	return score >= 0
}

func (d *Day13) parsePacket(line string) *d13Packet {
	stack := make([]*d13Packet, 1, 64)
	root := &d13Packet{}
	stack[0] = root
	p := root
	for i := 0; i < len(line); i++ {
		char := line[i]

		switch char {
		case '[':
			nw := &d13Packet{}
			p.values = append(p.values, d13Value{pkt: nw})
			stack = append(stack, p)
			p = nw
		case ']':
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
		case ',':
			// Next value
		default:
			if !unicode.IsDigit(rune(char)) {
				log.Panicf("unexpected non number %v", rune(char))
			}

			if unicode.IsDigit(rune(line[i+1])) {
				num, err := strconv.ParseInt(line[i:i+2], 10, 32)
				util.CheckErr(err)
				p.values = append(p.values, d13Value{num: int(num)})
				i++
				continue
			}
			num, err := strconv.ParseInt(string(line[i]), 10, 32)
			util.CheckErr(err)
			p.values = append(p.values, d13Value{num: int(num)})
		}
	}
	// Strip the root wrapper packet
	return root.values[0].pkt
}

type d13Packet struct {
	values []d13Value
}

func (p *d13Packet) inOrder(rgt *d13Packet) int {
	for i := range p.values {
		if i >= len(rgt.values) {
			return -1
		}

		order := p.values[i].inOrder(&rgt.values[i])
		if order < 0 {
			return order
		}
		if order > 0 {
			return order
		}
	}

	return len(rgt.values) - len(p.values)
}

func (p *d13Packet) String() string {
	str := make([]string, 0, 128)
	for i := range p.values {
		str = append(str, p.values[i].String())
	}

	return fmt.Sprintf("[%s]", strings.Join(str, ","))
}

type d13Value struct {
	num int
	pkt *d13Packet
}

func (v *d13Value) inOrder(rgt *d13Value) int {
	if v.pkt != nil && rgt.pkt != nil {
		return v.pkt.inOrder(rgt.pkt)
	}

	if v.pkt != nil && rgt.pkt == nil {
		tmp := &d13Value{pkt: &d13Packet{values: []d13Value{
			{num: rgt.num},
		}}}
		return v.inOrder(tmp)
	}

	if v.pkt == nil && rgt.pkt != nil {
		tmp := &d13Value{pkt: &d13Packet{values: []d13Value{
			{num: v.num},
		}}}
		return tmp.inOrder(rgt)
	}

	return rgt.num - v.num
}

func (v *d13Value) String() string {
	if v.pkt != nil {
		return v.pkt.String()
	}
	return strconv.FormatInt(int64(v.num), 10)
}

func (d *Day13) SolveI(input string) any {
	pairs := d.getPairs(input)
	sum := 0
	for i := range pairs {
		if pairs[i].inOrder() {
			sum += i + 1
		}
	}
	return sum
}

func (d *Day13) getPackets(pairs []d13Pair) []*d13Packet {
	pkts := make([]*d13Packet, 2*len(pairs))
	for i, pair := range pairs {
		pkts[i*2] = pair.left
		pkts[i*2+1] = pair.right
	}
	return pkts
}

func (d *Day13) SolveII(input string) any {
	pkts := d.getPackets(d.getPairs(input))

	var first, second = "[[2]]", "[[6]]"

	pkts = append(pkts, d.parsePacket(first))
	pkts = append(pkts, d.parsePacket(second))

	sort.Slice(pkts, func(i, j int) bool {
		return pkts[i].inOrder(pkts[j]) > 0
	})

	var pos1, pos2 int
	for i := range pkts {
		if pkts[i].String() == first {
			pos1 = i + 1
		}
		if pkts[i].String() == second {
			pos2 = i + 1
		}
	}

	return pos1 * pos2
}
