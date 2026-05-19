package main

import "fmt"

func main() {

	var i interface{} = "hello"

	switch v := i.(type) {
	case string:
		fmt.Println(v)
	}

}
