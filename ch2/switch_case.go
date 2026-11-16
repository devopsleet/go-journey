package main

import "fmt"

func main() {

	// Switch Statement (switch case deafult) (fallthrough)
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown Fruit")
	}

	day := "Friday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a Weekday")
	case "Sunday":
		fmt.Println("It's a Weekend")
	default:
		fmt.Println("Invalid day")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

}
