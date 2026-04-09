package main

import (
	"fmt"
	//"math/rand"
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

}
