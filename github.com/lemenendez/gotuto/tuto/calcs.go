/*
Create a function hat simulates a basic calculator's ability to evaluate arithmetic (+,-,*,/)

Input: an array of single character strings, representing key strokes on a calculator, they can be numbers 1-9,+,-,*,/=

Output: an array of integer representing what would be shown on the calculator screen after each key press

The order of the operator does not apply, i.e. 1+2*3=9, 1*2+3=5, 1+2*3+4=14

To keep things simple, there is no need to handle negative numbers.

Example:
input: ["1" "1" "+" "2" "=" ]
ouput: [1, 11, 11, 2, 13]

input: ["1" "2" "*" "3" "="]
output: [1, 1, 2, 3, 3, 9]


*/

package main

import (
	"fmt"
	"strconv"
)

func operate(ope string, a int, b int) int {
	switch ope {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func parse(num string) int {
	A, err := strconv.Atoi(num)
	if err != nil {
		return -1
	}
	return A
}

func add(a []int, num string) []int {
	A, err := strconv.Atoi(num)
	if err != nil {
		return a
	}
	a = append(a, A)
	return a
}

func calculate(inputs []string) []int {
	r := []int{}
	operandA := ""
	operandB := ""
	flag := true
	ope := ""
	for i := 0; i < len(inputs); i++ {
		if inputs[i] == "=" {
			A := parse(operandA)
			B := parse(operandB)
			C := operate(ope, A, B)
			r = append(r, C)
			operandA = ""
			operandB = ""
		} else if inputs[i] == "+" || inputs[i] == "-" || inputs[i] == "*" || inputs[i] == "/" {
			if operandA != "" && operandB != "" {
				A := parse(operandA)
				B := parse(operandB)
				C := operate(ope, A, B)
				r = append(r, C)
				operandA = ""
			}

			flag = !flag
			ope = inputs[i]
		} else if inputs[i] == "0" || inputs[i] == "1" || inputs[i] == "2" || inputs[i] == "3" || inputs[i] == "4" || inputs[i] == "5" || inputs[i] == "6" || inputs[i] == "7" || inputs[i] == "8" || inputs[i] == "9" {
			if flag {
				operandA = operandA + inputs[i]
				r = add(r, operandA)

			} else {
				operandB = operandB + inputs[i]
				r = add(r, operandB)
			}
		}
	}
	return r
}

func main() {
	x1 := []string{"1", "+", "1", "="}
	x2 := []string{"2", "+", "2", "="}
	x3 := []string{"4", "*", "4", "="}

	a := []string{"1", "1", "+", "2", "="}
	b := []string{"1", "+", "2", "*", "3", "="}

	c := []string{"2", "5", "*", "2", "/", "5", "="}

	fmt.Println(calculate(x1))
	fmt.Println(calculate(x2))
	fmt.Println(calculate(x3))
	fmt.Println(calculate(a))
	fmt.Println(calculate(b))
	fmt.Println(calculate(c))

}
