package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

// method with a value receiver
func (p Person) String() string {
	//fmt.Printf("Address of received P is %p\n", &p)
	//p.Age++
	return fmt.Sprintf("Name is %s and my age is %d", p.Name, p.Age)

}

// method with a pointer receiver
func (pp *Person) IncrementAge() {
	//fmt.Printf("Address of received P is %p\n", pp)
	//fmt.Println(*(pp))
	pp.Age++
}

func main() {

	var p Person
	p = Person{
		Name: "Gagan",
		Age:  32,
	}
	fmt.Println(p.String())
	//fmt.Printf("Address of Original P is %p\n", &p)
	// fmt.Println("The value of Original P is ", p)
	fmt.Println(strings.Repeat("*", 10))
	p.IncrementAge()
	fmt.Println(p.String())

}
