package main

import "fmt"

// const x int64 = 10

func main() {

	// untyped constant
	// const x = 10

	// var y int = x
	// var z float64 = x
	// var d byte = x

	// fmt.Println("The values are ", y, z, d)

	// // typed constants

	// const typedX int = 10

	a := 10

	fmt.Println("The value of a  is ", a)

	const b = 20

	var x = [3]int{10, 20}
	fmt.Println(x[2])

	var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("The sparse array is ", y)

	var z = [...]int{10, 20, 30}
	fmt.Println("Teh value of array z is ", z)
}
