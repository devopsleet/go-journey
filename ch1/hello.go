package main

import "fmt"

var x int = 10
var y float64 = 30.2

var sum1 float64 = float64(x) + y
var sum2 = x + int(y)
func main() {
	fmt.Println(sum1, sum2)
}