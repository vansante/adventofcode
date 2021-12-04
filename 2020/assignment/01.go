package assignment

type Day01 struct{}

const d01LookFor = 2020

func (d *Day01) SolveI(input string) int64 {
	nums := MakeIntegers(SplitLines(input))

	var foundTwo bool
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if !foundTwo && nums[i]+nums[j] == d01LookFor {
				return nums[i] * nums[j]
			}
		}
	}
	panic("no result")
}

func (d *Day01) SolveII(input string) int64 {
	nums := MakeIntegers(SplitLines(input))

	var foundThree bool
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if foundThree || nums[i]+nums[j] >= d01LookFor { // Some early elimination
				continue
			}

			for k := range nums {
				if i == k || j == k {
					continue
				}

				if nums[i]+nums[j]+nums[k] == d01LookFor {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}
	panic("no result")
}
