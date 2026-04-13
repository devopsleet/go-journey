package main

import (
	"errors"
	"fmt"
	"os"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	return nil
}

func addTobase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func main() {

	result, remainder, err := divAndRemainder(5, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
	// fmt.Println(addTobase(3, 1, 2))
	// fmt.Println(addTobase(3, 2, 4, 6, 8, 10))
	// fmt.Println(addTobase(3, []int{1, 2, 3}...))

	// nums := []int{1, 2, 3}
	// for _, n := range nums {
	// 	fmt.Println("The value is ", n)
	// }

	// MyFunc(MyFuncOpts{
	// 	LastName: "Patel",
	// 	Age:      50,
	// })
	// MyFunc(MyFuncOpts{
	// 	FirstName: "Joe",
	// 	LastName:  "Smith",
	// })

}
