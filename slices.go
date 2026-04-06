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

	z := make([]int, 5)
	z = append(z, 1, 2, 3, 4)
	fmt.Println("z is ", z)

	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4)
	fmt.Println("s is ", s)
	fmt.Println("Capacity of s is", cap(s))

}
