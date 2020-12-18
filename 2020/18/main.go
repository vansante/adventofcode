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

func tokenize(in string) []string {
	in = strings.ReplaceAll(in, "(", "( ")
	in = strings.ReplaceAll(in, ")", " )")
	tkns := strings.Split(in, " ")

	var nwTkns []string
	for i := range tkns {
		nwTkns = append(nwTkns, strings.TrimSpace(tkns[i]))
	}
	return nwTkns
}

// https://en.wikipedia.org/wiki/Shunting-yard_algorithm#The_algorithm_in_detail
func parse(tkns []string, ptII bool) []string {
	var operators, ret []string

	// while there are tokens to be read: read a token.
	for _, tkn := range tkns {
		// if the token is a number, then: push it to the output queue.
		_, err := strconv.ParseInt(tkn, 10, 64)
		if err == nil {
			ret = append(ret, tkn)
			continue
		}
		// else if the token is a left parenthesis (i.e. "("), then: push it onto the operator stack
		if tkn == "(" {
			operators = append(operators, tkn)
			continue
		}
		// else if the token is a right parenthesis (i.e. ")"), then:
		if tkn == ")" {
			foundLeftParenthesis := false
			// while the operator at the top of the operator stack is not a left parenthesis:
			//   pop the operator from the operator stack onto the output queue
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				if top == "(" {
					foundLeftParenthesis = true
					break
				}
				ret = append(ret, top)
				operators = operators[:len(operators)-1]
			}

			if !foundLeftParenthesis {
				panic(fmt.Sprintf("%v : mismatched parentheses found", tkns))
			}

			// if there is a left parenthesis at the top of the operator stack, then: pop the operator from the operator stack and discard it
			if operators[len(operators)-1] == "(" {
				operators = operators[:len(operators)-1]
			}
			continue
		}

		// while the operator at the top of the operator stack is not a left parenthesis:
		//   pop the operator from the operator stack onto the output queue.
		for len(operators) > 0 {
			top := operators[len(operators)-1]
			if top == "(" {
				break
			}
			if ptII && top == "*" && tkn == "+" {
				break
			}

			operators = operators[:len(operators)-1]
			ret = append(ret, top)
		}

		operators = append(operators, tkn)
	}

	// After while loop, if operator stack not null, pop everything to output queue
	for i := range operators {
		if operators[i] == "(" {
			panic(fmt.Sprintf("%v : mismatched parentheses found", tkns))
		}
		// pop the operator from the operator stack onto the output queue.
		ret = append(ret, operators[i])
	}
	return ret
}

func calculate(tkns []string) int64 {
	var stack []int64
	for _, tkn := range tkns {
		num, err := strconv.ParseInt(tkn, 10, 64)
		if err == nil {
			stack = append(stack, num)
			continue
		}

		var a, b int64
		a, b, stack = stack[len(stack)-1], stack[len(stack)-2], stack[:len(stack)-2]
		switch tkn {
		case "+":
			stack = append(stack, a+b)
		case "*":
			stack = append(stack, a*b)
		}
	}
	return stack[0]
}

func main() {
	wd, _ := os.Getwd()
	lines := retrieveInputLines(filepath.Join(wd, "18/input.txt"))

	var totalPtI, totalPtII int64
	for i := range lines {
		totalPtI += calculate(parse(tokenize(lines[i]), false))
		totalPtII += calculate(parse(tokenize(lines[i]), true))
	}
	fmt.Println(totalPtI)

	// > 337395165925601
	fmt.Println(totalPtII)
}
