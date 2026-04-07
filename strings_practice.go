package main

import "fmt"

func main() {

	var s string = "Hello World"
	var b byte = s[6]
	fmt.Println("Character is ", string(s[6]))
	fmt.Println("character value is ", b)

}
