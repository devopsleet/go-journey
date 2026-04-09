package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//shadowed variable
	x := 10
	if x > 5 {
		fmt.Println(x)
		// shadowing variable
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// shadowing language identifiers
	fmt.Println(true)
	true := 10
	fmt.Println(true)

	// declare variabels scoped to the condition
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number: ", n)
	}

}
