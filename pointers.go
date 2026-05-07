package main

import "fmt"

func main() {

	x := 10
	var pointerToX *int
	pointerToX = &x
	//pointerToX := &x
	fmt.Println(*pointerToX)

}
