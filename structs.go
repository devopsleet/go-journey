package main

import "fmt"

type test struct {
	Name string
	Age  int
}

func main() {

	t := test{
		Name: "Gagan",
	}

	fmt.Println(t)

}
