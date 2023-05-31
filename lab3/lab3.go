package main

import (
	"alg/pkgs/measure"
	"alg/pkgs/stack"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operator struct {
	name       string
	precedence int
	operation  func(float64, float64) float64
}

var operators = map[string]Operator{
	"+": {"+", 1, func(a, b float64) float64 { return a + b }},
	"-": {"-", 1, func(a, b float64) float64 { return a - b }},
	"*": {"*", 2, func(a, b float64) float64 { return a * b }},
	"/": {"/", 2, func(a, b float64) float64 { return a / b }},
	"^": {"^", 3, func(a, b float64) float64 { return math.Pow(a, b) }},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression: ")
	expression, _ := reader.ReadString('\n')

	measure.Measure(func() {
		fmt.Println("Result: ", calculate(expression))
	})

}

func calculate(expression string) float64 {
	tokens := tokenize(expression)
	rpn := toRPN(tokens)
	fmt.Println(tokens)
	fmt.Println(rpn)
	var stack = stack.NewStack[float64]()
	for _, token := range rpn {
		if isNumber(token) {
			number, _ := strconv.ParseFloat(token, 64)
			stack.Push(number)
		} else if isOperator(token) {
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			stack.Push(operators[token].operation(b, a))
		}
	}

	result, _ := stack.Pop()

	return result
}

func toRPN(tokens []string) []string {
	var rpn []string
	var operatorStack = stack.NewStack[string]()

	for _, token := range tokens {
		if isNumber(token) {
			rpn = append(rpn, token)
		} else if isOperator(token) {
			for !operatorStack.IsEmpty() {
				top, _ := operatorStack.Peek()
				if isOperator(top) && operators[top].precedence >= operators[token].precedence {
					operatorStack.Pop()
					rpn = append(rpn, top)
				} else {
					break
				}
			}
			operatorStack.Push(token)
		} else if token == "(" {
			operatorStack.Push(token)
		} else if token == ")" {
			for !operatorStack.IsEmpty() {
				top, _ := operatorStack.Peek()
				if top == "(" {
					operatorStack.Pop()
					break
				} else {
					operatorStack.Pop()
					rpn = append(rpn, top)
				}
			}
		}
	}

	for !operatorStack.IsEmpty() {
		top, _ := operatorStack.Pop()
		rpn = append(rpn, top)
	}

	return rpn
}

func tokenize(expression string) []string {
	expression = strings.Trim(expression, "\n")
	expression = strings.ReplaceAll(expression, " ", "")

	var tokens []string
	var token string

	var bracketStack = stack.NewStack[string]()

	for _, c := range expression {
		if c == '(' {
			bracketStack.Push("(")
			if isNumber(token) {
				tokens = append(tokens, token)
				tokens = append(tokens, "*")
				tokens = append(tokens, "(")
				token = ""
			} else {
				tokens = append(tokens, "(")
			}
		} else if c == ')' {
			if _, empty := bracketStack.Pop(); empty {
				panic("Invalid expression")
			}
			if isNumber(token) {
				tokens = append(tokens, token)
				tokens = append(tokens, ")")
				token = ""
			}
		} else if isOperator(string(c)) {
			if c == '-' && token == "" && len(tokens) != 0 && tokens[len(tokens)-1] == "(" {
				tokens = append(tokens, "0")
			} else {
				if isNumber(token) {
					tokens = append(tokens, token)
					token = ""
				}
			}

			tokens = append(tokens, string(c))
		} else {
			token += string(c)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	return tokens
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	_, ok := operators[token]
	return ok
}
