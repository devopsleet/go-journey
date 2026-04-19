package main

import "fmt"

func main() {

	// Simple iteration
	for i := 1; i <= 5; i++ {
		fmt.Println("The value of i is ", i)
	}

	// iteration over collection
	numbers := []int{1, 2, 3, 4, 5}
	for idx, val := range numbers {
		fmt.Printf("Index: %v, Value: %d\n", idx, val)

	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Numbe encountered: ", i)
		if i == 5 {
			break
		}
	}

	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println() // Move to he next line
	}

	for l := range 10 {
		l += 2
		fmt.Println("The value of l is ", l)
	}

}
