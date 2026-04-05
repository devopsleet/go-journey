package main

import "fmt"

func main() {
	//revenue, expenses, tax_rate := 0.0, 0.0, 0.0

	rv, exp, trate := getInput()
	fmt.Println(rv)
	fmt.Println(exp)
	fmt.Println(trate)

}

func getInput() (revenue float64, expenses float64, tax_rate float64) {

	//revenue, expenses, tax_rate := 0.0, 0.0, 0.0

	fmt.Println("Input all the values")

	fmt.Scan(&revenue)
	fmt.Scan(&expenses)
	fmt.Scan(&tax_rate)

	return
}
