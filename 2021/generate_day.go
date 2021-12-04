//go:build ignore

package main

import (
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

	_ = os.Mkdir(fmt.Sprintf("%02d", day), 0644)

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
	defer file.Close()

	dayTemplate.Execute(file, struct {
		Num int64
	}{
		Num: day,
	})
}

var dayTemplate = template.Must(template.New("").Parse(`
package assignment

type Day{{ .Num }} struct{}

func (d *Day{{ .Num }}) SolveI(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}

func (d *Day{{ .Num }}) SolveII(input string) int64 {
	// TODO: FIXME: Implement me!
	panic("no result")
}

`))
