package main

import "fmt"

type Person struct {
	FirtName string
	LastName string
	Age      int
}

func (p Person) printInfo() string {
	return fmt.Sprintf("%s is the FirstName", p.FirtName)
}

func main() {

	p := Person{
		"Gagan",
		"Singla",
		32,
	}

	fmt.Println(p.printInfo())
}
