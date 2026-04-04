package main

import (
	"fmt"
	//"slices"
	"math"
)

// variable declarations

// const x = 10

//const a = 40


func main() {
	const inflationRate = 6.5
	investmentAmount, years := 10000.0, 10.0
	expectedRunRate := 5.5

	fmt.Print("Enter the input value = ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+ expectedRunRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureRealValue)
	fmt.Println(futureValue)

}