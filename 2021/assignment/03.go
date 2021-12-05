package assignment

import (
	"math"
	"strconv"
)

type Day03 struct{}

type d03Number int16

func (n *d03Number) SetBit(idx int, one bool) {
	if !one {
		return
	}
	*n = *n + d03Number(math.Pow(2, float64(idx)))
}

func (n *d03Number) Bit(idx int) bool {
	if idx < 0 || idx > 15 {
		panic("invalid index")
	}
	pow2 := int16(math.Pow(2, float64(idx)))
	return int16(*n)&pow2 > 0
}

func (d *Day03) GetNumbers(lines []string) []d03Number {
	ns := make([]d03Number, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.ParseInt(line, 2, 16)
		CheckErr(err)
		ns = append(ns, d03Number(n))
	}
	return ns
}

func (d *Day03) countOnes(ns []d03Number, idx int) int {
	countOne := 0
	for _, n := range ns {
		if n.Bit(idx) {
			countOne++
		}
	}
	return countOne
}

func (d *Day03) filter(ns []d03Number, keep func(n d03Number) bool) []d03Number {
	filtered := make([]d03Number, 0, len(ns))
	for i := range ns {
		if keep(ns[i]) {
			filtered = append(filtered, ns[i])
		}
	}
	return filtered
}

func (d *Day03) SolveI(input string) int64 {
	ns := d.GetNumbers(SplitLines(input))

	var gamma, epsilon d03Number
	for i := 0; i < 12; i++ {
		ones := d.countOnes(ns, i)
		zeroes := len(ns) - ones
		gamma.SetBit(i, ones*2 > len(ns))
		epsilon.SetBit(i, ones > 0 && zeroes*2 > len(ns))
	}

	return int64(gamma) * int64(epsilon)
}

func (d *Day03) SolveII(input string) int64 {
	ns := d.GetNumbers(SplitLines(input))

	oxygen := ns
	co2 := make([]d03Number, len(ns))
	copy(co2, ns)

	for i := 11; i >= 0; i-- {
		if len(oxygen) > 1 {
			ones := d.countOnes(oxygen, i)
			zeroes := len(oxygen) - ones

			oxygen = d.filter(oxygen, func(n d03Number) bool {
				if ones >= zeroes {
					return n.Bit(i)
				}
				return !n.Bit(i)
			})
		}

		if len(co2) > 1 {
			ones := d.countOnes(co2, i)
			zeroes := len(co2) - ones

			co2 = d.filter(co2, func(n d03Number) bool {
				if ones == 0 {
					return true
				}
				if ones >= zeroes {
					return !n.Bit(i)
				}
				return n.Bit(i)
			})
		}
	}

	if len(oxygen) != 1 {
		panic("too many oxygen numbers")
	}

	if len(co2) != 1 {
		panic("too many co2 numbers")
	}

	return int64(oxygen[0]) * int64(co2[0])
}
