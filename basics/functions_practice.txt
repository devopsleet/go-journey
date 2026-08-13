package main

import "fmt"

func main() {

	// f := func(j int) {
	// 	fmt.Println("printing ", j)
	// }

	for i := 0; i < 5; i++ {

		func(j int) {
			fmt.Println("printing ", j)
		}(i)
	}

}
