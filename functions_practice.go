package main

import (
	"errors"
	"fmt"
)

func Denom(num, denom int) (result, remainder int, err error) {

	if denom == 0 {
		return 0, 0, errors.New("divide by zero")
	}

	result, remainder = num/denom, num%denom

	return
}

func main() {
	res, rem, err := Denom(5, 2)
	fmt.Println(res, rem, err)
}
