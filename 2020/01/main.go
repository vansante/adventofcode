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
	nums := retrieveNumbers(retrieveInputLines(filepath.Join(wd, "01/input.txt")))
	const lookFor = 2020

	var foundTwo, foundThree bool

	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if !foundTwo && nums[i]+nums[j] == lookFor {
				fmt.Printf("%d with 2 numbers: %d + %d = %d\n\n", lookFor, nums[i], nums[j], nums[i]*nums[j])
				foundTwo = true
			}

			if foundThree || nums[i]+nums[j] >= lookFor { // Some early elimination
				continue
			}

			for k := range nums {
				if i == k || j == k {
					continue
				}

				if nums[i]+nums[j]+nums[k] == lookFor {
					fmt.Printf("%d with 3 numbers: %d + %d + %d = %d\n\n", lookFor, nums[i], nums[j], nums[k], nums[i]*nums[j]*nums[k])
					foundThree = true
					break
				}
			}
		}
	}
}
