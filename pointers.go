package main

import "fmt"

func update(y []int) {
	fmt.Println("lenth is ", len(y))
	fmt.Println("capacity is ", cap(y))
	y[2] = 30
	y = append(y, 40)
	fmt.Println("lenth(after) is ", len(y))
	fmt.Println("capacity(after) is ", cap(y))
	y[1] = 20
	fmt.Println("The updated slice is ", y)
}

func main() {

	// nil pointer
	x := []int{1, 2, 3, 4}
	update(x)
	fmt.Println("The original slice is ", x)

}
