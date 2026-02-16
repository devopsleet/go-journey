package main

import "fmt"

const x int64 = 10

const (
	idKey = "id"
	namekey = "name"
)

const z = 20 * 10

func main() {
	const x = 10

	var y int = x
	var z float64= x
	var d byte = x

	fmt.Println(y,z,d)
	fmt.Printf("%f\n",z)
}
