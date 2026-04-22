package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
