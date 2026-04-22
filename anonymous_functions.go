package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {

		// inner anonymous function
		func(j int) {
			fmt.Println("printing ", j, "from inside the anonymous function")
		}(i)

	}

}
