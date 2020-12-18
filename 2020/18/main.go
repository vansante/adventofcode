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
func parse(tkns []string, precedence map[string]int) []string {
	var operators, out []string

	// while there are tokens to be read: read a token.
	for _, tkn := range tkns {
		// if the token is a number, then: push it to the output queue.
		_, err := strconv.ParseInt(tkn, 10, 64)
		if err == nil {
			out = append(out, tkn)
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
					// if there is a left parenthesis at the top of the operator stack, then: pop the operator from the operator stack and discard it
					operators = operators[:len(operators)-1]
					break
				}
				out = append(out, top)
				operators = operators[:len(operators)-1]
			}
			if !foundLeftParenthesis {
				panic(fmt.Sprintf("%v : mismatched parentheses found", tkns))
			}
			continue
		}

		// while ((there is an operator at the top of the operator stack)
		for len(operators) > 0 {
			// and ((the operator at the top of the operator stack has greater precedence)
			//     or (the operator at the top of the operator stack has equal precedence and the token is left associative))
			// and (the operator at the top of the operator stack is not a left parenthesis)):
			top := operators[len(operators)-1]
			if top == "(" {
				break
			}
			if precedence[tkn] > precedence[top] {
				break
			}

			// pop the operator from the operator stack onto the output queue.
			operators = operators[:len(operators)-1]
			out = append(out, top)
		}

		operators = append(operators, tkn)
	}

	// After while loop, if operator stack not null, pop everything to output queue
	for len(operators) > 0 {
		var top string
		operators, top = operators[:len(operators)-1], operators[len(operators)-1]
		if top == "(" {
			panic(fmt.Sprintf("%v : mismatched parentheses found", tkns))
		}
		// pop the operator from the operator stack onto the output queue.
		out = append(out, top)
	}
	return out
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
		totalPtI += calculate(parse(tokenize(lines[i]), map[string]int{}))
		totalPtII += calculate(parse(tokenize(lines[i]), map[string]int{"+": 1}))
	}
	fmt.Printf("Part I: %d\n\n", totalPtI)

	fmt.Printf("Part II: %d\n\n", totalPtII)
}
