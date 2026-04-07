package main

import "fmt"

func main() {

	// x := make([]int, 4)
	// fmt.Println(x)

	// y := make([]int, 0, 10)
	// y = append(y, 1, 2, 3, 4)
	// fmt.Println("New slice is ", y)

	// copy a slice
	// x := []int{1, 4, 6}
	// y := make([]int, 2)
	// num := copy(y, x[2:])
	// fmt.Println(y, num)

	// // Convert arrays to slice
	// xArray := [4]int{1, 2, 3, 4}
	// xSlice := xArray[:]
	// fmt.Println("Slice is ", xSlice)

	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)

}
