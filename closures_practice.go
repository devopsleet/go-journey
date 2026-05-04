package main

import "fmt"

func helper(x, multiplier int) int {
	return x * multiplier
}

func process(nums []int, multiplier int) {
	for _, val := range nums {
		fmt.Println(helper(val, multiplier))
		//val * helper(val, multiplier)
	}
}

func main() {

	nums := []int{1, 2, 3, 4}
	process(nums, 3)

}
