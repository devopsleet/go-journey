package main

import (
	"fmt"
	"math"
)

// var x int = 10
// var y float64 = 30.2

// var sum1 float64 = float64(x) + y
// var sum2 = x + int(y)

// const a int64 = 10

// // untyped constant
// const x = 10

// var y int = x
// var z float64 = x
// var b byte = x

var b byte = math.MaxUint8
var smallI int32 = math.MaxInt32
var bigI uint64 = math.MaxUint64

func main() {
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}