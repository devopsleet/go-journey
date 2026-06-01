package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Test() string {
	p.Age++
	p.Name = "Test"
	return fmt.Sprintf("%s and %d", p.Name, p.Age)
}

func main() {

	var p Person
	p = Person{
		Name: "Gagan",
		Age:  31,
	}

	output := p.Test()
	fmt.Println(output)

	fmt.Println("The origianls p is", p)

}
