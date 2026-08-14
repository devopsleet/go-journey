package main

import "fmt"

func main() {

	var r rune = 'a'
	fmt.Println("Rune literal value is ", r)
	fmt.Printf("Character value is %c\n", r)

	str := "Greetings and\n\"Salutations"
	fmt.Printf("The value of a string is %s", str)

	var flag bool
	fmt.Println("The default value of flag is", flag)

	var isAwesome = true
	fmt.Println("The value of isAwesome is ", isAwesome)

	var myFirstInitial rune = 'G'
	var myLastinitial int32 = 'S'

	fmt.Printf("The first initial is %c and last initial is %c", myFirstInitial, myLastinitial)
}
