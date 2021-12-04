package assignment

import (
	"strconv"
	"strings"
)

type Day13 struct{}

func (d *Day13) getBuses(line string) []int64 {
	buses := strings.Split(line, ",")
	var nums []int64
	for i := range buses {
		if buses[i] == "x" {
			nums = append(nums, -1)
			continue
		}

		num, err := strconv.ParseInt(buses[i], 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func (d *Day13) findNearestTime(tm int64, buses []int64) (time, bus int64) {
	for ; ; tm++ {
		for _, bus := range buses {
			if bus == -1 {
				continue
			}
			if tm%bus == 0 {
				return tm, bus
			}
		}
	}
}

func (d *Day13) findGoldenCoin(buses []int64) int64 {
	time := int64(1)
	step := int64(1)
	for i, bus := range buses {
		if bus == -1 {
			continue
		}

		for (time+int64(i))%bus != 0 {
			time += step
		}
		step *= bus
	}
	return time
}

func (d *Day13) SolveI(input string) int64 {
	lines := SplitLines(input)

	targetTime, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		panic(err)
	}

	buses := d.getBuses(lines[1])
	time, bus := d.findNearestTime(targetTime, buses)

	return (time - targetTime) * bus
}

func (d *Day13) SolveII(input string) int64 {
	lines := SplitLines(input)

	buses := d.getBuses(lines[1])
	return d.findGoldenCoin(buses)
}
