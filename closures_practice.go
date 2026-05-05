package main

import "fmt"

func multiplier(base int) func(int) int {

	return func(factor int) int {
		return base * factor
	}

}

func main() {

	twoBase := multiplier(2)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i))
	}

}
