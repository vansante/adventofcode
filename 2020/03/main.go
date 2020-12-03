package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type row []bool
type grid []row

func retrieveInputLines(file string) []string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(content), "\n")

	var input []string
	for i := range split {
		line := strings.TrimSpace(split[i])
		if line != "" {
			input = append(input, line)
		}
	}
	return input
}

func processLines(lines []string) grid {
	var rows []row
	for i := range lines {
		var r row
		for _, char := range strings.Split(lines[i], "") {
			switch char {
			case ".":
				r = append(r, false)
			case "#":
				r = append(r, true)
			default:
				panic(fmt.Sprintf("unexpected char: %v", char))
			}
		}
		rows = append(rows, r)
	}
	return rows
}

type slope struct {
	x int
	y int
}

func findSlopeTrees(g grid, sl slope) int {
	var x, trees int

	for y := 0; y < len(g); y += sl.y {
		curX := x
		if curX >= len(g[y]) {
			curX %= len(g[y])
		}

		//fmt.Printf("x: %d y: %d | %v\n", x, y, g[y][curX])
		if g[y][curX] {
			trees++
		}

		x += sl.x
	}
	return trees
}

func main() {
	wd, _ := os.Getwd()
	g := processLines(retrieveInputLines(filepath.Join(wd, "03/input.txt")))

	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	answerPt2 := 0
	for i := range slopes {
		trees := findSlopeTrees(g, slopes[i])
		fmt.Printf("Amount of trees in %d lines with slope %v: %d\n", len(g), slopes[i], trees)
		if i == 0 {
			answerPt2 = trees
		} else {
			answerPt2 *= trees
		}
	}

	fmt.Printf("All trees multiplied: %d", answerPt2)
}
