package main

import (
	"fmt"
	//"math"
)

// variable declarations

// const x = 10

const a = 40


func main() {
	 
	var x [3]int

	fmt.Println(x[0])

	var a = [...]int{1,2,3}
	var b = [3]int{1,2,3}

	fmt.Println(a==b)
	 
	var y = [3]int{1,2}
	fmt.Println(y[1])

	
}