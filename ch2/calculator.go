package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num, denom int) (result, remainder int, err error) {

	if denom == 0 {
		err = errors.New("Cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

func main() {

	res, rem, err := divAndRemainder(0, 2)

	if err != nil {
		fmt.Println("The result is ", res, " and remainder is ", rem)
	} else {
		fmt.Println("Error : ", err)
	}

}
