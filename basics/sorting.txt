package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 3, 5, 7}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println(nums)

}
