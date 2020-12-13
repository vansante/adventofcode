package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func retrieveInputLines(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	return input
}

func getBuses(line string) []int64 {
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

func findNearestTime(tm int64, buses []int64) (time, bus int64) {
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

func findGoldenCoin(buses []int64) int64 {
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

func main() {
	wd, _ := os.Getwd()
	lines := retrieveInputLines(filepath.Join(wd, "13/input.txt"))

	targetTime, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		panic(err)
	}
	buses := getBuses(lines[1])
	time, bus := findNearestTime(targetTime, buses)
	fmt.Printf("Part I: Time: %d | WaitTime: %d | Bus: %d | Multiplied: %d\n\n", time, time-targetTime, bus, (time-targetTime)*bus)

	time = findGoldenCoin(buses)
	fmt.Printf("Part II: Time: %d\n\n", time)
}
