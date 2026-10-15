package main

import (
	"fmt"
	//"slices"
)

func main() {

	xSlice := []int{10, 20, 30, 40}
	subSlice := xSlice[1:]

	fmt.Println("The Full slice is ", xSlice)
	fmt.Println("The Sub Slice is ", subSlice)

	// Sub Slice share the underlying array with the main Slice
	subSlice[2] = 50

	fmt.Println("xSlice: ", xSlice)
	fmt.Println("subSlice: ", subSlice)

}
