package main

import "fmt"

// simple function declaration
func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}

	return num / denom
}

func main() {

	output := div(2, 0)
	fmt.Println("The finals answer is ", output)

}
