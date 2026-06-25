package main

import "fmt"

func main() {

	// var arrayName [Size]elementType
	a, b := someFunction()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("")

	// comparinfg arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println(array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("The original array is ", originalArray)
	fmt.Println("The copied array is ", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
