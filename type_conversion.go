package main

import "fmt"

func main() {
	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)

	fmt.Println(x, y, z, d, e, f, g)

	const a int64 = 10

	const (
		idKey   = "id"
		nameKey = "name"
	)

	const b = 20 * 10

	fmt.Println("The value of b is ", b)
}
