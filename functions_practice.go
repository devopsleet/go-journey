// program to add a base to the numbers

package main

import "fmt"

func AddToBase(base int, vals ...int) []int {
	s := make([]int, 0, len(vals))

	for _, val := range vals {
		s = append(s, base+val)
	}

	return s
}

func main() {

	input := []int{3, 4, 5, 6}
	output := AddToBase(2, input...)
	fmt.Println("The output is ", output)

}
