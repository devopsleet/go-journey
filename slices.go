package main

import "fmt"

var x = []int{}

var m = map[int]int{}

func main() {

	fmt.Println(x[2])
	fmt.Println(m[2])
	fmt.Println(x == nil)
	fmt.Println(m == nil)

}
