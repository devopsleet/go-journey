package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("It's a main function")

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(x, y))

	x = append(x, 6)
	fmt.Println("The new slice is ", x)

	z := []int{10, 20, 30}
	x = append(x, z...)

	fmt.Println("The new vlaue of slice x is ", x)
}
