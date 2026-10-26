package main

import (
	"errors"
	"fmt"
	"os"
)

type MyFuncOpts struct {
	First string
	Last  string
	Age   int
}

func MyFunc(opts MyFuncOpts) {

	fmt.Println("My FirtName is ", opts.First)
}

func addTo(base int, vals ...int) []int {

	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}

	return out

}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {

	MyFunc(MyFuncOpts{
		First: "Gagan",
		Last:  "Singla",
	})

	fmt.Println(addTo(3))
	fmt.Println(addTo(2, 3))
	fmt.Println(3, 2, 4, 6, 8)

	a := []int{4, 3}

	fmt.Println(addTo(3, a...))

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println("The error is ", err)
		os.Exit(1)
	}

	fmt.Println(result, remainder)

}
