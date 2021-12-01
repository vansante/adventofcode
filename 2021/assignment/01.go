package assignment

type Day01 struct {}

func (d *Day01) SolveI(input string) int64 {
	depths := MakeIntegers(SplitLines(input))

	var total int64
	for i := range depths {
		if i == 0 {
			continue
		}

		if depths[i-1] < depths[i] {
			total++
		}
	}
	return total
}

func (d *Day01) SolveII(input string) int64 {
	depths := MakeIntegers(SplitLines(input))

	var total int64
	for i := range depths {
		if i <= 2 {
			continue
		}

		sumI := depths[i-1] + depths[i-2] + depths[i-3]
		sumII := depths[i] + depths[i-1] + depths[i-2]

		if sumII > sumI {
			total++
		}
	}
	return total
}