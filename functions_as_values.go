package main

import "fmt"

func add(x, y int) int {

	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func compute(f func(int, int) int) {
	fmt.Println("The value is ", f(2, 3))
}

func main() {

	//var f func(int, int) int

	compute(add)
	compute(multiply)

	// f = add

	// fmt.Println("The value of addition is ", f(2, 3))

}
