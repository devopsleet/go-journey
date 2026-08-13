package main

import "fmt"

func main() {

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	//slice := make([]int, 5) // Initiate a Slice

	a := [5]int{1, 2, 3, 4, 5}
	slice2 := a[1:4]

	fmt.Println(slice2)

}
