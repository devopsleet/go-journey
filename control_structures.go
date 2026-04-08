package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if n := rand.Intn(10); n == 0 {
		fmt.Println("This is too low")
	} else if n > 5 {
		fmt.Println("That's too big", n)
	} else {
		fmt.Println("That's a good number ", n)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

}
