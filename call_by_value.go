package main

import "fmt"

// map is not call by value
// map is a reference type so it changes the original map
// func modifyMap(m map[string]int) {
// 	m["a"] = 10
// 	m["c"] = 30
// 	delete(m, "a")
// }

func modifySlice(s []int) {
	for idx, _ := range s {
		s[idx] *= 2
	}

	s = append(s, 10)
}

func main() {

	// m := make(map[string]int)
	// m["a"] = 1
	// m["b"] = 2
	// fmt.Println("Map before call by value ", m)
	// modifyMap(m)
	// fmt.Println("Map after call by value ", m)

	s := make([]int, 3)
	s = append(s, 4)
	fmt.Println("Slice before call by value = ", s)
	modifySlice(s)
	fmt.Println("Slice after call by value = ", s)

}
