package main

import (
	"fmt"
	"math"
)

// variable declarations

// const x = 10

const a = 40


func main() {
	 var investmentAmount float64 = 1000
	 var expectedReturnRate = 5.5
	 var years float64 = 10
	
	 var futureValue = investmentAmount* math.Pow(1 + expectedReturnRate/100, years)
	 fmt.Println(futureValue)
	 

	
}