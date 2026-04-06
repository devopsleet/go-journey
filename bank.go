package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!!!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Print("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsCheckBalance := choice == 1

	if wantsCheckBalance {

		fmt.Println("Your balance is ", accountBalance)

	} else if choice == 2 {
		fmt.Print("How much? = ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("New balance is ", accountBalance)

	}

	fmt.Println("Your choice: ", choice)

}
