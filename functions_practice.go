package main

import (
	"fmt"
)

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func main() {
	result := div(5, 2)
	fmt.Println(result)

	// map declarations
	//var m = map[string]int{}
	d := map[string]int{
		"Kitten": 1,
	}
	m := make(map[string]int)

	// zeo value of amp is nil
	if m == nil {
		fmt.Println("Map is Nil")
	}

	m["x"] = 1
	m["y"] = 2

	fmt.Println(m["z"])

	val, ok := m["z"]
	fmt.Println(val, ok)

	val, ok = m["world"]
	fmt.Println(val, ok)

	// m["a"] = 1

	var n = map[string]int{}

	n["a"] = 1

	fmt.Println(n["a"])

	// type person struct {
	// 	name string
	// 	age  int
	// }

	// var beth person

	// beth.name = "Beth"
	// beth.age = 40
	//beth.animal = "Cat"

	// beth := person{
	// 	"Beth",
	// 	32,
	// }

	// fmt.Println(beth.name)

	// Anonymous Struct

	var person struct {
		name string
		age  int
		pet  string
	}

	var fred person
}
