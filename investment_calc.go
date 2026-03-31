package main

import "fmt"

// variable declarations

// const x = 10

const a = 40


func main() {
	 x := 10
	 x = 20
	 fmt.Println(x)
	 x = 30

	 _0 := 0o_1_2_3
	 fmt.Println(_0)

	 var a = [...]int{10,20,30}
	 var b = [3]int{10,20,30}

	 fmt.Println(a == b)

	 var c [2][3]int

	 c[0] = [3]int{1,2,3}
	 c[0][0] = 10

	 //c[0] = 10

	 fmt.Println(len(c))




	
}