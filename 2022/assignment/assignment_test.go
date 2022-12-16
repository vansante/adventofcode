package assignment

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/vansante/adventofcode/2022/util"
)

func Test_Day_01_SolveI(t *testing.T) {
	d := Day01{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(1, "example")))
	valid := "24000"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_01_SolveII(t *testing.T) {
	d := Day01{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(1, "example")))
	valid := "45000"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_02_SolveI(t *testing.T) {
	d := Day02{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(2, "example")))
	valid := "15"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_02_SolveII(t *testing.T) {
	d := Day02{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(2, "example")))
	valid := "12"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_03_SolveI(t *testing.T) {
	d := Day03{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(3, "example")))
	valid := "157"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_03_SolveII(t *testing.T) {
	d := Day03{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(3, "example")))
	valid := "70"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_04_SolveI(t *testing.T) {
	d := Day04{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(4, "example")))
	valid := "2"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_04_SolveII(t *testing.T) {
	d := Day04{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(4, "example")))
	valid := "4"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_05_SolveI(t *testing.T) {
	d := Day05{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(5, "example")))
	valid := "CMZ"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_05_SolveII(t *testing.T) {
	d := Day05{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(5, "example")))
	valid := "MCD"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_06_SolveI(t *testing.T) {
	d := Day06{}

	examples := util.SplitLines(getInput(6, "1_example"))
	answers := []int{5, 6, 10, 11}
	for i := range examples {
		a := d.SolveI(examples[i])
		if a != answers[i] {
			t.Errorf("%s : %v is not equal to %v", examples[i], a, answers[i])
		}
	}
}

func Test_Day_06_SolveII(t *testing.T) {
	d := Day06{}
	examples := util.SplitLines(getInput(6, "2_example"))
	answers := []int{19, 23, 23, 29, 26}
	for i := range examples {
		a := d.SolveII(examples[i])
		if a != answers[i] {
			t.Errorf("%s : %v is not equal to %v", examples[i], a, answers[i])
		}
	}
}

func Test_Day_07_SolveI(t *testing.T) {
	d := Day07{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(7, "example")))
	valid := "95437"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_07_SolveII(t *testing.T) {
	d := Day07{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(7, "example")))
	valid := "24933642"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_08_SolveI(t *testing.T) {
	d := Day08{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(8, "example")))
	valid := "21"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_08_SolveII(t *testing.T) {
	d := Day08{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(8, "example")))
	valid := "8"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_09_SolveI(t *testing.T) {
	d := Day09{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(9, "1_example")))
	valid := "13"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_09_SolveII(t *testing.T) {
	d := Day09{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(9, "2_example")))
	valid := "36"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_10_SolveI(t *testing.T) {
	d := Day10{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(10, "example")))
	valid := "13140"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_10_SolveII(t *testing.T) {
	d := Day10{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(10, "example")))
	valid := `░▓░░▓▓░░▓▓░░▓▓░░▓▓░░▓▓░░▓▓░░▓▓░░▓▓░░▓▓░░
▓▓▓░░░▓▓▓░░░▓▓▓░░░▓▓▓░░░▓▓▓░░░▓▓▓░░░▓▓▓░
▓▓▓▓░░░░▓▓▓▓░░░░▓▓▓▓░░░░▓▓▓▓░░░░▓▓▓▓░░░░
▓▓▓▓▓░░░░░▓▓▓▓▓░░░░░▓▓▓▓▓░░░░░▓▓▓▓▓░░░░░
▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓░░░░░░▓▓▓▓▓▓░░░░░░▓▓▓▓
▓▓▓▓▓▓▓░░░░░░░▓▓▓▓▓▓▓░░░░░░░▓▓▓▓▓▓▓░░░░░
`

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_11_SolveI(t *testing.T) {
	d := Day11{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(11, "example")))
	valid := "10605"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_11_SolveII(t *testing.T) {
	d := Day11{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(11, "example")))
	valid := "2713310158"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_12_SolveI(t *testing.T) {
	d := Day12{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(12, "example")))
	valid := "31"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_12_SolveII(t *testing.T) {
	d := Day12{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(12, "example")))
	valid := "29"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_13_SolveI(t *testing.T) {
	d := Day13{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(13, "example")))
	valid := "13"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_13_SolveII(t *testing.T) {
	d := Day13{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(13, "example")))
	valid := "140"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_14_SolveI(t *testing.T) {
	d := Day14{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(14, "example")))
	valid := "24"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_14_SolveII(t *testing.T) {
	d := Day14{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(14, "example")))
	valid := "93"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_15_SolveI(t *testing.T) {
	d := Day15{}
	sum := d.findImpossible(getInput(15, "example"), 10).sumImpossible(10)
	valid := 26

	if sum != valid {
		t.Errorf("%v is not equal to %v", sum, valid)
	}
}

func Test_Day_15_SolveII(t *testing.T) {
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

func Test_Day_16_SolveI(t *testing.T) {
	d := Day16{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(16, "example")))
	valid := "1651"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_16_SolveII(t *testing.T) {
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
	input, err := os.ReadFile(path.Join(dir, fmt.Sprintf("../%02d/%s.txt", day, fileName)))
	util.CheckErr(err)
	return string(input)
}
