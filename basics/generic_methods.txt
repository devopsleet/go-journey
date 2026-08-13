package main

import (
	"fmt"
	"strconv"
)

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))

	for idx, val := range s {
		r[idx] = f(val)
	}

	return r
}

func main() {

	nums := []int{1, 2, 3}
	result := Map(nums, func(n int) string {
		return strconv.Itoa(n)
	})
	fmt.Println("The result is", result)

}
