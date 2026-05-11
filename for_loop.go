package main

import "fmt"

func main() {

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Odd number value is ", i)
	// }

	// rows := 7

	// Outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")

	// 	}
	// 	// inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10 {
		fmt.Println(i)
	}

}
