package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func subtract(x, y int) int {
	return x - y
}

// function type declaration
type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": subtract,
}

func main() {
	var op string
	fmt.Println("Please enter the operator you want to use ")
	fmt.Scan(&op)

	expressions := [][]string{
		{"2", op, "3"},
		{"two", "+", "three"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("It's an invalid expression ", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println("Error is ", err)
			continue

		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator ", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println("The error is ", err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println("The result is ", result)

	}

}
