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

func retrieveNumbers(input []string) []int64 {
	var nums []int64
	for i := range input {
		num, err := strconv.ParseInt(input[i], 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func main() {
	wd, _ := os.Getwd()
	lines := retrieveInputLines(filepath.Join(wd, "2021/01/input.txt"))
	depths := retrieveNumbers(lines)

	var totalPtI, totalPtII int64
	for i := range depths {
		if i == 0 {
			continue
		}

		if depths[i-1] < depths[i] {
			totalPtI++
		}

		if i <= 2 {
			continue
		}

		sumI := depths[i-1] + depths[i-2] + depths[i-3]
		sumII := depths[i] + depths[i-1] + depths[i-2]

		if sumII > sumI {
			totalPtII++
		}
	}

	fmt.Printf("Part I: %d\n\n", totalPtI)
	fmt.Printf("Part II: %d\n\n", totalPtII)
}
