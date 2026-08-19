package main

import (
	"fmt"
	"slices"
)

var x = [3]int{10, 20, 30}
var y = [12]int{1, 5: 4, 6, 10: 100, 12}
var z = [...]int{10, 20, 30}

var a = []int{1, 2, 3}
var b = []int{1, 2, 3}
var c = []int{1, 2, 3, 4}
var d = []string{"a", "b", "c"}

func main() {
	fmt.Println(x == z)
	fmt.Println(x[2])

	// i := 7
	// //fmt.Println(x[i])
	// x[i] = 100

	var s []int
	fmt.Println(s == nil)

	fmt.Println(slices.Equal(a, b))
	fmt.Println(slices.Equal(b, c))
	//fmt.Println(slices.Equal(c, d))

	// shadowing the scope of outer x and y
	var x = []int{1, 2, 3}
	y := []int{20, 30, 40}
	x = append(x, y...)
}
