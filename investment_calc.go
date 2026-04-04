package main

import (
	"fmt"
	"math"
)

const expectedRunRate = 5.5

func main() {

	investmentAmount := 0.0
	years := 0.0
	inflation := 5.0

	fmt.Print("Enter the inverstment Amount = ")
	_, err := fmt.Scan(&investmentAmount)
	if err != nil {
		fmt.Print("Invalid Input")
		return
	}

	fmt.Print("Enter number of years = ")
	_, err = fmt.Scan(&years)
	if err != nil {
		fmt.Print("Invalid Input")
		return
	}

	fv := calculateFutureValue(investmentAmount, years)
	rfv := calculateRealFutureValue(fv, inflation, years)

	fmt.Println(fv)
	fmt.Println(rfv)

}

func calculateFutureValue(investmentAmount float64, years float64) float64 {

	return (investmentAmount * math.Pow(1+expectedRunRate/100, years))

}

func calculateRealFutureValue(futureVal float64, inflation float64, years float64) float64 {
	return (futureVal / math.Pow(1+inflation/100, years))
}
