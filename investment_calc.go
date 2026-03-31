package main

import "fmt"

// variable declarations

const x = 10


func main() {
	var y int = x
	fmt.Println(y)
	var z float64 = x
	fmt.Println(z)
	var d byte = x
	fmt.Println(d)

	fmt.Print(x)
	
}