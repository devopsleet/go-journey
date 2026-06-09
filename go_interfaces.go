package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Your name is %s and age is %d", p.Name, p.Age)
}

func (p *Person) Increment() {
	p.Age++
}

type Stringer interface {
	String() string
}

type Incrementer interface {
	Increment()
}

func main() {

	p := &Person{
		"Gagan",
		32,
	}

	var myStringer Stringer

	myStringer = p

	var myIncrementer Incrementer

	myIncrementer = p

	fmt.Println(myStringer.String())
	fmt.Println(myIncrementer.Increment)

}
