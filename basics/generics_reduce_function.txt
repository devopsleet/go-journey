package main

func Reduce[T1, T2 any](s []T1, initializer T2,
	f func(T2, T1) T2,
) T2 {

}

func main() {

	nums := []int{1, 2, 3}

	result := Reduce(nums, func(n, v int) int {
		return n + v
	})

}
