package main

import "fmt"

// const x int64 = 10

func main() {

	x := 5
	y := 10

	// compile time error
	//const z = x + y

	z := x + y

	fmt.Println(z)

}
