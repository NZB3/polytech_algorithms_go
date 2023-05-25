package main

import (
	"algorithms/packages/stack"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your expression: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	expression := strings.ReplaceAll(line, " ", "")
	expression = strings.Trim(expression, "\n")

	calculated_expr, err := calculate(expression)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculated_expr)

}

func getPriority(operator string) int {
	switch operator {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	case "(", ")":
		return -1
	default:
		return 0
	}
}

func calculate(expression string) (result float64, err error) {
	operators := stack.NewStack[string]()
	numbers := stack.NewStack[float64]()

	pattern := `(\d+\.\d+|\d+|\(|\)|\+|\-|\*|\^|\/)`
	re := regexp.MustCompile(pattern)
	tokens := re.FindAllString(expression, -1)

	for _, token := range tokens {
		token_priority := getPriority(token)

		if token_priority == 0 {
			floatToken, err := strconv.ParseFloat(token, 64)
			if err != nil {
				log.Fatal("unexpected token", err)
			}

			numbers.Push(floatToken)
			continue
		}

		top_operator, operatorsIsEmpty := operators.Top()
		if operatorsIsEmpty {
			operators.Push(token)
			continue
		}

		top_operator_priority := getPriority(top_operator)
		for !operatorsIsEmpty && top_operator_priority >= token_priority {
			operator, operatorsIsEmpty := operators.Pop()

			operand1, numbersIsEmpty := numbers.Pop()
			if numbersIsEmpty && !operatorsIsEmpty {
				log.Fatalf("found extra operator")
			}

			operand2, numbersIsEmpty := numbers.Pop()
			if numbersIsEmpty && !operatorsIsEmpty {
				log.Fatalf("invalid expression, missing second operand")
			}

			switch operator {
			case "*":
				numbers.Push(operand1 * operand2)
			case "/":
				numbers.Push(operand1 / operand2)
			case "+":
				numbers.Push(operand1 + operand2)
			case "-":
				numbers.Push(operand1 - operand2)
			case "^":
				numbers.Push(math.Pow(float64(operand1), float64(operand2)))
			}
		}
	}

	result, numbersIsEmpty := numbers.Pop()
	if numbersIsEmpty {
		err = fmt.Errorf("calculate faled")
	}
	return result, err
}

