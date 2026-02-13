package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	fmt.Printf("Hello, %s!\n", "world")

	// var x int = 10

	// x *= 2

	// fmt.Println(x)

	var myFirstinitial rune = 'g'

	fmt.Printf("%c\n",myFirstinitial)

	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1,sum2)
}
