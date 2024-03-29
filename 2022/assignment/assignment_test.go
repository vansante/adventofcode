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
	valid := "1707"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_17_SolveI(t *testing.T) {
	d := Day17{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(17, "example")))
	valid := "3068"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_17_SolveII(t *testing.T) {
	d := Day17{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(17, "example")))
	valid := "1514285714288"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_18_SolveI(t *testing.T) {
	d := Day18{}

	answer := d.SolveI("1,1,1\n2,1,1")
	if answer != int64(10) {
		t.Errorf("%v is not equal to 10", answer)
	}

	answer = fmt.Sprintf("%v", d.SolveI(getInput(18, "example")))
	valid := "64"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_18_SolveII(t *testing.T) {
	d := Day18{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(18, "example")))
	valid := "58"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_19_SolveI(t *testing.T) {
	d := Day19{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(19, "example")))
	valid := "33"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_19_SolveII(t *testing.T) {
	d := Day19{}
	bps := d.getBlueprints(getInput(19, "example"))

	bc := d.makeBotCollections(bps[0])
	answer := bc.collect(32, d19Bots{oreBots: 1}, d19Resources{})
	if answer != 56 {
		t.Errorf("%v is not equal to %v", answer, 56)
	}

	bc = d.makeBotCollections(bps[1])
	answer = bc.collect(32, d19Bots{oreBots: 1}, d19Resources{})
	if answer != 62 {
		t.Errorf("%v is not equal to %v", answer, 62)
	}
}

func Test_Day_20_SolveI(t *testing.T) {
	d := Day20{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(20, "example")))
	valid := "3"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_20_SolveII(t *testing.T) {
	d := Day20{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(20, "example")))
	valid := "1623178306"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_21_SolveI(t *testing.T) {
	d := Day21{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(21, "example")))
	valid := "152"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_21_SolveII(t *testing.T) {
	d := Day21{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(21, "example")))
	valid := "301"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_22_SolveI(t *testing.T) {
	d := Day22{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(22, "example")))
	valid := "6032"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_22_SolveII(t *testing.T) {
	in := getInput(22, "input")

	d := Day22{}
	grid, directions := d.getNotes(in)
	coord, facing := grid.walk(grid.findStart(), d22FaceRight, directions, grid.wrapCube)
	if coord.x != 37 {
		t.Errorf("X: %v is not %v", coord.x, 37)
	}
	if coord.y != 103 {
		t.Errorf("Y: %v is not %v", coord.y, 103)
	}
	if facing != d22FaceUp {
		t.Errorf("Dir: %v is not %v", facing, d22FaceUp)
	}

	answer := d.SolveII(in).(int)
	if answer != 37415 {
		t.Errorf("%v is not %v", answer, 37415)
	}
}

func Test_Day_23_SolveI(t *testing.T) {
	d := Day23{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(23, "example")))
	valid := "110"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_23_SolveII(t *testing.T) {
	d := Day23{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(23, "example")))
	valid := "20"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_24_SolveI(t *testing.T) {
	d := Day24{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(24, "example")))
	valid := "18"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_24_SolveII(t *testing.T) {
	d := Day24{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput(24, "example")))
	valid := "54"

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func Test_Day_25_SolveI(t *testing.T) {
	d := Day25{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput(25, "example")))
	valid := "2=-1=0"

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
