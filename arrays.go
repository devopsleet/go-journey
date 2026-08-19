package main

import "fmt"

var x = [3]int{10, 20, 30}
var y = [12]int{1, 5: 4, 6, 10: 100, 12}
var z = [...]int{10, 20, 30}

func main() {
	fmt.Println(x == z)
}
