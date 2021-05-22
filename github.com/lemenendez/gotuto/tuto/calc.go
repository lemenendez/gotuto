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

func parseNum(num string) int {
	A, err := strconv.Atoi(num)
	if err != nil {
		return -1
	}
	return A
}

func calculate(inputs []string) []int {
	r := []int{}
	lastOpe := ""
	lastA := ""
	lastB := ""

	operFlag := true
	for i := 0; i < len(inputs); i++ {

		if inputs[i] == "=" {
			A := parseNum(lastA)
			B := parseNum(lastB)
			r = append(r, operate(lastOpe, A, B))
			fmt.Print(lastOpe)
			lastA = ""
			lastB = ""
		} else if inputs[i] == "+" || inputs[i] == "-" || inputs[i] == "*" || inputs[i] == "/" {
			prevOpe := lastOpe
			lastOpe = inputs[i]
			operFlag = !operFlag
			A := parseNum(lastA)
			r = append(r, A)

			if lastA != "" && lastB != "" {
				A := parseNum(lastA)
				B := parseNum(lastB)

				res := operate(prevOpe, A, B)
				r = append(r, res)
				lastA = strconv.Itoa(res)
				lastB = ""
				operFlag = false
			}
		} else if inputs[i] == "0" || inputs[i] == "1" || inputs[i] == "2" || inputs[i] == "3" || inputs[i] == "4" || inputs[i] == "5" || inputs[i] == "6" || inputs[i] == "7" || inputs[i] == "8" || inputs[i] == "9" {
			if operFlag {
				lastA = lastA + inputs[i]
				A := parseNum(lastA)
				r = append(r, A)
			} else {
				lastB = lastB + inputs[i]
				B := parseNum(lastB)
				r = append(r, B)
			}
		}
	}
	return r
}

func main() {
	x1 := []string{"1", "+", "1", "="}
	//a := []string{"1", "1", "+", "2", "="}
	//b := []string{"1", "+", "2", "*", "3", "="}
	// c := []string{"2", "+", "2", "="}
	fmt.Println(calculate(x1))
	//fmt.Println(calculate(a))
	//fmt.Println(calculate(b))
	// fmt.Println(calculate(c))

}
