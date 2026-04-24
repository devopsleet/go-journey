package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s", p.name)
}

func main() {

	p := Person{
		"Gagan",
		30,
	}

	fmt.Println(p.String())

}
