package main

import "fmt"

var x *int

type Person struct {
	FirstName string
	LastName  *string
	Age       int
}

//fmt.Println("Outside")

func failedUpdate(g *int) {

	x := 10

	g = &x
}
func main() {

	fmt.Println(x == nil)

	p := &Person{}

	fmt.Println("The value of Person p is ", p)

	fmt.Printf("the address of Person p is %p\n", p)

	var f *int // nil

	failedUpdate(f)

	fmt.Println(f)

}
