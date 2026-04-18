package main

import (
	"fmt"
	"math"
)

func main() {

	// Variables
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction ", result)

	result = a * b
	fmt.Println("Multiply ", result)

	x := 22.0 / 7
	fmt.Println(x)

	var maxInt int64 = 1<<63 - 1
	maxInt += 1
	fmt.Println("Max value of Int64 is ", maxInt)

	fmt.Println("The max vlaue of int64 is ", math.MaxInt64)

	z := math.MaxFloat64 / 3
	fmt.Println("Tha value of z is ", z)

}
