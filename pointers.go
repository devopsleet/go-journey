package main

import "fmt"

func myFunc(g *int) {

	*g = 45
}
func main() {

	//var f *int
	x := 10
	myFunc(&x)
	fmt.Println(x)

}
