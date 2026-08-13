package main

import "fmt"

// variable of type function
var myFuncVariable func(string) int

func f1(a string) int {

	return len(a)

}

func main() {

	myFuncVariable = f1
	output := myFuncVariable("Welcome")
	fmt.Println("The length of the string is ", output)

}
