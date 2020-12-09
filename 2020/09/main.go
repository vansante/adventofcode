package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func retrieveInputNumbers(file string) []int64 {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []int64
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		input = append(input, num)
	}
	return input
}

func findNumberInSums(number int64, previousNumbers []int64) bool {
	for i := range previousNumbers {
		for j := range previousNumbers {
			if previousNumbers[i]+previousNumbers[j] == number {
				return true
			}
		}
	}
	return false
}

func main() {
	wd, _ := os.Getwd()
	numbers := retrieveInputNumbers(filepath.Join(wd, "09/input.txt"))

	const preamble = 25

	var invalidNumber int64
	for i := range numbers {
		if i <= preamble {
			continue
		}
		if !findNumberInSums(numbers[i], numbers[i-preamble-1:i]) {
			invalidNumber = numbers[i]
			break
		}
	}
	fmt.Printf("Part I: Found %d\n\n", invalidNumber)

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

	fmt.Printf("Part II: Found lowest: %d + highest: %d == %d\n\n", lowest, highest, lowest+highest)
}
