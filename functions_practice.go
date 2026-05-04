// functions are values

package main

import "fmt"

var myFuncVariable func(string) int

func f1(a string) int {
	return len(a)
}

func f2(b string) int {
	total := 0

	for _, val := range b {
		total += int(val)
	}

	return total
}

func main() {

	myFuncVariable := f1
	fmt.Println(myFuncVariable("Hello"))

	myFuncVariable = f2
	fmt.Println("The sum of all the ASCII characters is ", myFuncVariable("a"))

}
