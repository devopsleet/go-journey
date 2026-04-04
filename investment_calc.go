package main

import (
	"fmt"
	//"slices"
	//"math"
)

// variable declarations

// const x = 10

//const a = 40

func main() {

	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter the revenue you are making = ")
	fmt.Scan(&revenue)

	fmt.Print("How much are your expenses? = ")
	fmt.Scan(&expenses)

	fmt.Print("What is the new tax rate? = ")
	fmt.Scan(&tax_rate)

	earning_before_taxes := revenue - expenses

	profit := earning_before_taxes * (1 - tax_rate/100)

	ratio := earning_before_taxes / profit

	//fmt.Println(earning_before_taxes)
	formattedFEB := fmt.Sprintf("Earning before taxes: %.1f\n", earning_before_taxes)
	fmt.Printf(formattedFEB)

	fmt.Println(profit)

	fmt.Println(ratio)

}
