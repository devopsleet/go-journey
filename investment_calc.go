package main

import (
	"fmt"
	"slices"
	//"math"
)

// variable declarations

// const x = 10

//const a = 40


func main() {

	//Slices
	x:= []int{1,2,3}
	y:= []int{1,2,3,4}

	fmt.Println((slices.Equal(x,y)))


}