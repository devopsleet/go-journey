package main

import "fmt"

func main() {
	// var x int = 10
	// var pointerToX *int
	x := 10
	pointerToX := &x

	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	var y = new(int)
	println(y)
	println(*y)

}
