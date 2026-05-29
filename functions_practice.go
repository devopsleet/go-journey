package main

import "fmt"

// Public exported struct
type MyFuncOpts struct {
	FirstName string
	Lastname  string
	Age       int
}

// struct as named and positional parameters
func MyFunc(opts MyFuncOpts) string {
	return fmt.Sprintf("%s", opts.Lastname)

}

func main() {

	output := MyFunc(MyFuncOpts{
		FirstName: "Gagan",
	})

	fmt.Println("The output is ", output)

}
