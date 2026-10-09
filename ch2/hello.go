package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("Hello, %s!\n", "Gagan")

	var myFirstInitial rune = 'G'
	var myLastInitial int32 = 'S'

	fmt.Println(myFirstInitial)
	fmt.Println(myLastInitial)

	// var x int = 10
	// var y float64 = 30.2

	// var sum1 float64 = float64(x) + y
	// fmt.Println(sum1)

	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)

	fmt.Println(x, y, z, d, e, f, g)
}
