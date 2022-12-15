//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func generateInput(day uint8) {
	_ = os.Mkdir(fmt.Sprintf("%02d", day), 0755)

	example := fmt.Sprintf("%02d/example.txt", day)
	if _, err := os.Stat(example); err != nil {
		file, err := os.Create(example)
		if err != nil {
			panic(err)
		}
		_ = file.Close()
	}

	input := fmt.Sprintf("%02d/input.txt", day)
	if _, err := os.Stat(input); err != nil {
		file, err := os.Create(input)
		if err != nil {
			panic(err)
		}
		_ = file.Close()
	}
}

func generateCodeTemplate(day uint8) {
	newAssignment := fmt.Sprintf("assignment/%02d.go", day)
	_, err := os.Stat(newAssignment)
	if err == nil {
		panic("assignment exists!")
	}

	file, err := os.Create(newAssignment)
	if err != nil {
		panic(err)
	}

	dayTemplate.Execute(file, struct {
		Num string
	}{
		Num: fmt.Sprintf("%02d", day),
	})
	_ = file.Close()
}

func addToMain(day uint8) {
	mainFile, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}

	mainFile = bytes.Replace(
		mainFile,
		[]byte("// <generator:add:days>"),
		[]byte(fmt.Sprintf("%d: &assignment.Day%02d{},\n\t\t// <generator:add:days>", day, day)),
		1,
	)
	err = os.WriteFile("main.go", mainFile, 0755)
	if err != nil {
		panic(err)
	}
}

func addTestFunctions(day uint8) {
	testFile, err := os.ReadFile("assignment/assignment_test.go")
	if err != nil {
		panic(err)
	}

	output := &bytes.Buffer{}
	testTemplate.Execute(output, struct {
		Num string
	}{
		Num: fmt.Sprintf("%02d", day),
	})

	testFile = bytes.Replace(
		testFile,
		[]byte("// <generator:add:days>"),
		append(output.Bytes(), []byte("// <generator:add:days>")...),
		1,
	)
	err = os.WriteFile("assignment/assignment_test.go", testFile, 0755)
	if err != nil {
		panic(err)
	}
}

func main() {
	day64, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		panic(err)
	}

	day := uint8(day64)
	generateInput(day)
	generateCodeTemplate(day)
	addToMain(day)
	addTestFunctions(day)
}

var dayTemplate = template.Must(template.New("").Parse(
	`package assignment

type Day{{ .Num }} struct{}

func (d *Day{{ .Num }}) SolveI(input string) any {
	return "Not Implemented Yet"
}

func (d *Day{{ .Num }}) SolveII(input string) any {
	return "Not Implemented Yet"
}
`))

var testTemplate = template.Must(template.New("").Parse(
	`func TestDay{{ .Num }}_SolveI(t *testing.T) {
	d := Day{{ .Num }}{}
	answer := fmt.Sprintf("%v", d.SolveI(getInput({{ .Num }}, "example")))
	valid := "" // FIXME

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

func TestDay{{ .Num }}_SolveII(t *testing.T) {
	d := Day{{ .Num }}{}
	answer := fmt.Sprintf("%v", d.SolveII(getInput({{ .Num }}, "example")))
	valid := "" // FIXME

	if answer != valid {
		t.Errorf("%v is not equal to %v", answer, valid)
	}
}

`))
