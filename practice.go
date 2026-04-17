package main

import "fmt"

func main() {

	r := "I'm Go language\\"
	s := "Hello\nWorld"
	fmt.Println(r)
	fmt.Println(s)

	arr := []int{1, 2, 3}
	fmt.Println(arr[2])

	// var (
	// 	x int
	// 	y = 20
	// )

	// fmt.Println("x and y", x, y)

	// var myFirstInitial rune = "G"
	// var mySecondInitial int32 = "F"

	// arrays
	// var x = [3]int{10, 20, 30}
	// fmt.Println("The array values are ", x)

	// var y = [...]int{1, 2, 3, 4, 3, 3, 4, 45, 56, 5}
	// fmt.Println(y)

	// var z = []int{10, 20, 30}
	// fmt.Println(z)

	// var a []int
	// fmt.Println(a == nil)

	// a = append(a, 1, 2)
	// fmt.Println(a)

	// var b = []int{10, 20, 30}
	// b = append(b, 40, 50)

	// a = append(a, b...)
	// fmt.Println(a)

	// fmt.Println(b)

	//var y []int
	//fmt.Println(y == nil)
	//fmt.Println(len(y))

	z := []int{}
	fmt.Println(z == nil)
	fmt.Println(len(z))

	// x := make([]int, 5)
	// x = append(x, 10)
	// fmt.Println("Length of slice is ", len(x))
	// fmt.Println("Capacity of slice is ", cap(x))

	// Emptying a Slice
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println(x)
	fmt.Println(y)

	c := make([]int, 6)
	fmt.Println(len(c))
	//fmt.Println("Capacity is ", cap(c))

	ages := make(map[int][]string, 10)
	fmt.Println(len(ages))

}
