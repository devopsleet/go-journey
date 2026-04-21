package main

import "fmt"

type Person struct {
	FirstName string
	Lastname  *string
	age       int
}

func makePointer[T any](t T) *T {
	return &t
}

func main() {

	p := Person{
		FirstName: "Gagan",
		Lastname:  makePointer("Singla"),
		age:       14,
	}

	fmt.Println(p)

	// variables
	var x int32 = 10

	pointerX := &x

	var y bool = true
	pointerY := &y

	fmt.Println("X and Y are ", pointerX, pointerY)

	// var z *string
	// fmt.Println("The value of pointer z is ", z)

	a := "hello"
	pointerToA := &a
	fmt.Println(*pointerToA)

	z := 5 + *pointerX
	fmt.Println(z)

}
