package main

import "fmt"

func main() {

	// nil slice
	var a []int
	if a == nil {
		fmt.Println(`It's a NIL slice`)
	}

	// empty slice but not a nil slice
	y := []int{}

	fmt.Println("Its an empty slice")
	fmt.Println(len(y))

	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(x))

}
