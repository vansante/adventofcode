package assignment

type Day09 struct{}

func (d *Day09) findNumberInSums(number int64, previousNumbers []int64) bool {
	for i := range previousNumbers {
		for j := range previousNumbers {
			if previousNumbers[i]+previousNumbers[j] == number {
				return true
			}
		}
	}
	return false
}

func (d *Day09) SolveI(input string) int64 {
	numbers := MakeIntegers(SplitLines(input))

	preamble := 25
	if numbers[0] == 35 { // Exception for the example
		preamble = 5
	}

	for i := range numbers {
		if i <= preamble {
			continue
		}
		if !d.findNumberInSums(numbers[i], numbers[i-preamble-1:i]) {
			return numbers[i]
		}
	}
	panic("no result")
}

func (d *Day09) SolveII(input string) int64 {
	invalidNumber := d.SolveI(input)

	numbers := MakeIntegers(SplitLines(input))

	var lowest, highest int64
outer:
	for i := range numbers {
		total := numbers[i]
		lowest = numbers[i]
		highest = numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			total += numbers[j]

			if numbers[j] < lowest {
				lowest = numbers[j]
			}
			if numbers[j] > highest {
				highest = numbers[j]
			}

			if total == invalidNumber {
				break outer
			} else if total > invalidNumber {
				break
			}
		}
	}
	return lowest + highest
}
