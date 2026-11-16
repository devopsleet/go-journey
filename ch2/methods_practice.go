package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) sayHello() {
	fmt.Println("Hello ", p.Name)
}

func main() {

	p := Person{
		Name: "Gagan",
	}

	// method value

	greet := p.sayHello
	greet()

	// method expression
	greet2 := Person.sayHello

	greet2(p)

	fmt.Println(1 << 1)

}
