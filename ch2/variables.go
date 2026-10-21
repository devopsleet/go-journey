package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	fmt.Println(true)
	true := 10
	fmt.Println("The value of true is ", true)

	if n := rand.Int(10); n == 0 {
		fmt.Println("That's too low)
	}

}
