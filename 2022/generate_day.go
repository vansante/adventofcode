//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func main() {
	day, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		panic(err)
	}

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

	newAssignment := fmt.Sprintf("assignment/%02d.go", day)
	_, err = os.Stat(newAssignment)
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

	mainGo, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}

	mainGo = bytes.Replace(
		mainGo,
		[]byte("// <generator:add:days>"),
		[]byte(fmt.Sprintf("%d: &assignment.Day%02d{},\n\t\t// <generator:add:days>", day, day)),
		1,
	)
	err = os.WriteFile("main.go", mainGo, 0755)
	if err != nil {
		panic(err)
	}
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
