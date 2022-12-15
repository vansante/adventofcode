package assignment

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/vansante/adventofcode/2022/util"
)

func TestDay13_SolveI(t *testing.T) {
	d := Day13{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(13, "example")))
	valid := "13"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay13_SolveII(t *testing.T) {
	d := Day13{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(13, "example")))
	valid := "140"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay14_SolveI(t *testing.T) {
	d := Day14{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(14, "example")))
	valid := "24"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay14_SolveII(t *testing.T) {
	d := Day14{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(14, "example")))
	valid := "93"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay15_SolveI(t *testing.T) {
	d := Day15{}
	answer := fmt.Sprintf("%v", d.findImpossible(getInput(15, "example"), 10))
	valid := "26"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay15_SolveII(t *testing.T) {
	d := Day15{}
	answer := d.findPossible(getInput(15, "example"), 0, 20)
	valid := &d15Coord{14, 11}

	if answer.x != valid.x {
		t.Errorf("%v is not equal to %v", answer.x, valid.x)
	}
	if answer.y != valid.y {
		t.Errorf("%v is not equal to %v", answer.y, valid.y)
	}
}

func TestDay16_SolveI(t *testing.T) {
	d := Day16{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(16, "example")))
	valid := "" // FIXME

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay16_SolveII(t *testing.T) {
	d := Day16{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(16, "example")))
	valid := "" // FIXME

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

// <generator:add:days>

func getInput(day int, fileName string) string {
	dir, err := os.Getwd()
	util.CheckErr(err)
	input, err := os.ReadFile(path.Join(dir, fmt.Sprintf("../%d/%s.txt", day, fileName)))
	util.CheckErr(err)
	return string(input)
}
