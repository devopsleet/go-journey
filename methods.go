package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// method defined only at the package level

func (p Person) String() string {
	return fmt.Sprintf("My name is %s and my age is %d", p.FirstName, p.Age)
}

func main() {
	p := Person{
		"fred",
		"Fredson",
		22,
	}

	output := p.String()

	fmt.Println("The output is ", output)

}
