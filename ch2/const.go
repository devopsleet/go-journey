package main

import "fmt"

// const x int64 = 10

func main() {

	// untyped constant
	const x = 10

	var y int = x
	var z float64 = x
	var d byte = x

	fmt.Println("The values are ", y, z, d)

	// typed constants

	const typedX int = 10

}
