package assignment

import (
	"container/ring"
	"fmt"

	"github.com/vansante/adventofcode/2022/util"
)

type Day20 struct{}

type d20Number struct {
	num int
	pos int
}

func (n d20Number) equals(other d20Number) bool {
	return n.num == other.num && n.pos == other.pos
}

func (d *Day20) getNumbers(input string, multiplier int) *ring.Ring {
	var rng *ring.Ring
	ints := util.ParseInts(util.SplitLines(input))
	ints = util.SliceReverse(ints)

	for i, num := range ints {
		lnk := &ring.Ring{Value: d20Number{
			num: num * multiplier,
			pos: i,
		}}
		if rng == nil {
			rng = lnk
			continue
		}
		rng = lnk.Link(rng)
	}
	return rng
}

func (d *Day20) slice(nums *ring.Ring) []d20Number {
	s := make([]d20Number, 0, 10_000)
	nums.Do(func(a any) {
		s = append(s, a.(d20Number))
	})
	return s
}

func (d *Day20) print(nums *ring.Ring) {
	nums.Do(func(a any) {
		fmt.Print(a.(d20Number).num, ", ")
	})
	fmt.Println()
}

func (d *Day20) mix(nums *ring.Ring, num d20Number, listLen int) {
	if num.num == 0 {
		return // do nothing
	}

	remove := nums
	for {
		if remove.Value.(d20Number).equals(num) {
			break
		}
		remove = remove.Next()
	}

	prev := remove.Prev()
	removed := prev.Unlink(1)
	prev.Move(num.num % (listLen - 1)).Link(removed)
}

func (d *Day20) getCoordinates(nums *ring.Ring) int64 {
	zero := nums
	for {
		if zero.Value.(d20Number).num == 0 {
			break
		}
		zero = zero.Next()
	}

	var sum int64
	for i := 0; i < 3001; i++ {
		switch i {
		case 1000, 2000, 3000:
			sum += int64(zero.Value.(d20Number).num)
		}
		zero = zero.Next()
	}
	return sum
}

func (d *Day20) SolveI(input string) any {
	nums := d.getNumbers(input, 1)
	length := nums.Len()
	for _, num := range d.slice(nums) {
		d.mix(nums, num, length)
	}

	return d.getCoordinates(nums)
}

func (d *Day20) SolveII(input string) any {
	const decryptionKey = 811589153

	nums := d.getNumbers(input, decryptionKey)

	numOrder := d.slice(nums)
	length := nums.Len()
	for i := 0; i < 10; i++ {
		for _, num := range numOrder {
			d.mix(nums, num, length)
		}
	}

	return d.getCoordinates(nums)
}
