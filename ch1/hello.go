package main

import "fmt"

const x int64 = 10

const (
	idKey = "id"
	namekey = "name"
)

const z = 20 * 10

func main() {
	x := 10
	x = 20

	fmt.Println(x)

	x = 30
}
