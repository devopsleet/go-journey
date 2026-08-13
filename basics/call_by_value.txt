package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func modStruct(p person) {
	p.Name = "abc"
	fmt.Printf("%+v\n", p)
}

func modMap(m map[string]int) {
	m["a"] = 10
	m["c"] = 30
	fmt.Println(m)

}

func modSlice(s []int) {
	s[0] = 10
	fmt.Println("The copied slice is ", s)
}

func main() {

	p := person{
		Name: "Gagan",
		Age:  32,
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	modStruct(p)
	modMap(m)
	fmt.Println("The original map is ", m)

	s := []int{1, 2, 3, 4}

	modSlice(s)
	fmt.Println("The original slice is ", s)
}
