package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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

func findDifferences(current int64, numbers []int64) (diff1, diff2, diff3 int) {
	for i := range numbers {
		var diff int64
		switch numbers[i] {
		case current + 1:
			diff1++
			diff = 1
		case current + 2:
			diff2++
			diff = 2
		case current + 3:
			diff3++
			diff = 3
		default:
			continue
		}
		sub1, sub2, sub3 := findDifferences(current+diff, numbers)
		diff1 += sub1
		diff2 += sub2
		diff3 += sub3
		break
	}
	return diff1, diff2, diff3
}

func findCombinations(current int64, numbers []int64, values map[int64]int64) int64 {
	if len(numbers) == 0 {
		return 1
	}
	if val, ok := values[current]; ok {
		return val
	}

	var val int64
	for i := range numbers {
		if numbers[i] <= current || numbers[i] > current+3 {
			continue
		}
		val += findCombinations(numbers[i], numbers[i+1:], values)
	}
	values[current] = val
	return val
}

func main() {
	wd, _ := os.Getwd()
	numbers := retrieveInputNumbers(filepath.Join(wd, "10/input.txt"))

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	diff1, diff2, diff3 := findDifferences(0, numbers)
	diff3++ // For the device adapter

	fmt.Printf("The number of differences are: %d, %d, %d\n\n", diff1, diff2, diff3)
	fmt.Printf("Part I multiplied: %d\n", diff1*diff3)

	adapters := make(map[int64]int64)

	combinations := findCombinations(0, numbers, adapters)
	fmt.Printf("Part II combinations: %d\n", combinations)
}
