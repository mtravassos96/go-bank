package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	userChoice := getUserChoice()

	if userChoice == 1 {
		fmt.Println("Your balance is:")
	} else if userChoice == 2 {
		fmt.Print("How much do you want to deposit?:")
	} else if userChoice == 3 {
		fmt.Print("How much do you want to withdraw?:")
	} else if userChoice == 4 {
		fmt.Print("Thank for you for choosing Go Bank!")
	}
}

func getUserChoice() (int) {
	var userChoice int
	
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
	fmt.Scan(&userChoice)

	return userChoice
}
