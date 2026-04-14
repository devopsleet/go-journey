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
	var x = [3]int{10, 20, 30}
	fmt.Println("The array values are ", x)

	var y = [...]int{1, 2, 3, 4, 3, 3, 4, 45, 56, 5}
	fmt.Println(y)
}
