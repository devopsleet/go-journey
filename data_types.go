package main

import (
	"fmt"
	"strings"
)

func main() {

	var x int
	fmt.Println(x == 0)
	fmt.Println(strings.Repeat("*", 10))

	var y int = 12
	fmt.Println(y)

	var z = 13
	fmt.Println(z)

	a := 14
	fmt.Println(a)

	var b = new(int)
	fmt.Println(*b == 0)
	fmt.Println(b == nil)
	fmt.Println(b)
}
