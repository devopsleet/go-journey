package main

import "fmt"

func deferExample(a int) int {
	//a = 10
	defer func(a int) {
		fmt.Println("Function 1 ", a)
	}(a)

	a = 20
	defer func(a int) {
		fmt.Println("Function 2 ", a)
	}(a)

	a = 30
	fmt.Println("Exiting call")
	return a

}
func main() {

	result := deferExample(10)
	fmt.Println("The final result is ", result)

}
