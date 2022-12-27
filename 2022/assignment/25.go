package assignment

import (
	"github.com/vansante/adventofcode/2022/util"
)

type Day25 struct{}

func snafu2dec(in string) int64 {
	var sum int64

	factor := int64(1)
	for i := len(in) - 1; i >= 0; i-- {
		switch in[i] {
		case '=':
			sum -= factor * 2
		case '-':
			sum -= factor
		case '0':
			// Nothing to add or subtract
		case '1':
			sum += factor
		case '2':
			sum += factor * 2
		default:
			panic("invalid character")
		}
		factor *= 5
	}

	return sum
}

func dec2snafu(num int64) string {
	snaf := make([]byte, 0, 32)

	for num > 0 {
		remain := num % 5
		num /= 5

		if remain > 2 {
			num++
			remain -= 5
		}

		switch remain {
		case -2:
			snaf = append(snaf, '=')
		case -1:
			snaf = append(snaf, '-')
		case 0:
			snaf = append(snaf, '0')
		case 1:
			snaf = append(snaf, '1')
		case 2:
			snaf = append(snaf, '2')
		default:
			panic("invalid")
		}
	}
	return string(util.SliceReverse(snaf))
}

func (d *Day25) SolveI(input string) any {
	lines := util.SplitLines(input)
	var sum int64
	for _, snafu := range lines {
		sum += snafu2dec(snafu)
	}
	return dec2snafu(sum)
}

func (d *Day25) SolveII(input string) any {
	return "Not Implemented Yet"
}
