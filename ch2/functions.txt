package main

import "fmt"

//var myFuncVariable func(string) int

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}

	return total
}

func main() {

	var myFuncVariable func(string) int

	//myFuncVariable2 := func(string) int

	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result)

	// myFuncVariable2 = f2
	// result2 := myFuncVariable2("Gagan")
	// fmt.Println(result2)

}
