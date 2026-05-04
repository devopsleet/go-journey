package main

import "fmt"

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}

	return num / denom
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func Myfunc(Opts MyFuncOpts) {

	fmt.Printf("The data is %+v\n", Opts)

}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, val := range vals {
		out = append(out, base+val)
	}

	return out
}

func main() {

	fmt.Println(addTo(3, 4))

	result := div(5, 2)
	fmt.Println("result is ", result)

	Myfunc(MyFuncOpts{
		LastName: "Singla",
		Age:      32,
	})

}
