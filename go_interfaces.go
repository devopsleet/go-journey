package main

import "fmt"

type Speaker interface {
	Speak()
}

// Defins a Function type
type SpeakerFunc func()

// Attach a method to the function type
func (f SpeakerFunc) Speak() {
	f() // call the underlying function
}

func sayHello() {
	fmt.Println("Hello from a function")
}

func main() {

	var s Speaker = SpeakerFunc(sayHello)
	s.Speak()

}
