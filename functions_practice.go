package main

import (
	"fmt"
)

// func div(num int, denom int) int {
// 	if denom == 0 {
// 		return 0
// 	}
// 	return num / denom
// }

// type MyFuncOpts struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func MyFunc(opts MyFuncOpts) error {

// 	fmt.Println(opts.FirstName)
// 	return nil

// }

// func addTo(base int, vals ...int) []int {
// 	output := make([]int, 0, len(vals))

// 	for _, v := range vals {
// 		output = append(output, base+v)
// 	}

// 	return output
// }

func f1(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}

	return total
}

func main() {

	var myFuncVariable func(string) int

	myFuncVariable = f1

	result := myFuncVariable("Hello")
	fmt.Println("The result is ", result)

	// result := div(5, 2)
	// fmt.Println(result)

	// MyFunc(MyFuncOpts{
	// 	"Gagan",
	// 	"Singla",
	// 	32,
	// })

	// fmt.Println(addTo(1, 3, 5, 7, 9))
	// b := []int{2, 4, 6, 8}
	// fmt.Println(addTo(1, b...))
	// map declarations
	//var m = map[string]int{}
	// d := map[string]int{
	// 	"Kitten": 1,
	// }
	// m := make(map[string]int)

	// // zeo value of amp is nil
	// if m == nil {
	// 	fmt.Println("Map is Nil")
	// }

	// m["x"] = 1
	// m["y"] = 2

	// fmt.Println(m["z"])

	// val, ok := m["z"]
	// fmt.Println(val, ok)

	// val, ok = m["world"]
	// fmt.Println(val, ok)

	// // m["a"] = 1

	// var n = map[string]int{}

	// n["a"] = 1

	// fmt.Println(n["a"])

}
