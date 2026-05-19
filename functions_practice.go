package main

import "fmt"

func sum(numbers ...int) (result int) {

	for _, val := range numbers {
		result += val
	}

	return
}

func main() {

	fmt.Println("The total sum is ", sum(1, 2, 3, 4, 5))

}
