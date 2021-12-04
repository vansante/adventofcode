package assignment

import (
	"strconv"
	"strings"
)

type Day04 struct{}

type day04board struct {
	rows [][]day04number
	mark bool
}

type day04number struct {
	val  int64
	mark bool
}

func (b *day04board) Mark(num int64) {
	for y := range b.rows {
		for x := range b.rows[y] {
			if b.rows[y][x].val == num {
				b.rows[y][x].mark = true
			}
		}
	}
}

func (b *day04board) Check() bool {
	b.mark = b.checkRows() || b.checkCols()
	return b.mark
}

func (b *day04board) checkRows() bool {
	for y := range b.rows {
		if b.checkRow(y) {
			return true
		}
	}
	return false
}

func (b *day04board) checkRow(y int) bool {
	for x := range b.rows[y] {
		if !b.rows[y][x].mark {
			return false
		}
	}
	return true
}

func (b *day04board) checkCols() bool {
	for x := 0; x < 5; x++ {
		if b.checkCol(x) {
			return true
		}
	}
	return false
}

func (b *day04board) checkCol(x int) bool {
	for y := range b.rows {
		if !b.rows[y][x].mark {
			return false
		}
	}
	return true
}

func (b *day04board) unmarkedSum() int64 {
	var sum int64
	for y := range b.rows {
		for x := range b.rows[y] {
			if !b.rows[y][x].mark {
				sum += b.rows[y][x].val
			}
		}
	}
	return sum
}

func (d *Day04) CountMarked(boards []day04board) int {
	count := 0
	for _, b := range boards {
		if b.mark {
			count++
		}
	}
	return count
}

func (d *Day04) GetInput(input string) (numbers []day04number, boards []day04board) {
	lines := strings.Split(input, "\n")

	var b *day04board
	for i := range lines {
		if i == 0 {
			numbers = d.GetNumbers(lines[0], ",")
			continue
		}

		if lines[i] == "" {
			if b != nil {
				boards = append(boards, *b)
			}
			b = &day04board{}
			continue
		}

		b.rows = append(b.rows, d.GetNumbers(lines[i], " "))
	}
	return numbers, boards
}

func (d *Day04) GetNumbers(input, split string) (numbers []day04number) {
	input = strings.TrimSpace(strings.ReplaceAll(input, "  ", " "))
	numStrs := strings.Split(input, split)
	for _, numStr := range numStrs {
		num, err := strconv.ParseInt(numStr, 10, 8)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, day04number{val: num})
	}
	return numbers
}

func (d *Day04) SolveI(input string) int64 {
	ns, bs := d.GetInput(input)

	for _, n := range ns {
		for _, b := range bs {
			b.Mark(n.val)
			if b.Check() {
				return n.val * b.unmarkedSum()
			}
		}
	}
	panic("no solved boards")
}

func (d *Day04) SolveII(input string) int64 {
	ns, bs := d.GetInput(input)

	for _, n := range ns {
		var lastMarked day04board
		for i := range bs {
			bs[i].Mark(n.val)
			if !bs[i].mark && bs[i].Check() {
				lastMarked = bs[i]
			}
		}
		if d.CountMarked(bs) == len(bs) {
			return lastMarked.unmarkedSum() * n.val
		}
	}
	panic("no result")
}
