package main

import (
	//"crypto/rand"
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(true)
	true := 10
	fmt.Println("The value of true is ", true)

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big ", n)
	} else {
		fmt.Println("That's a good number ", n)
	}

}
