package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number: ", i)

		if i == 5 {
			break
		}
	}

	// Simple iteration over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("The vlaue of i is ", i)
	// }

	// iteration over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	//fmt.Println(index, value)
	// 	fmt.Printf("Index: %v, Value:%d\n", index, value)
	// }

	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces for stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
