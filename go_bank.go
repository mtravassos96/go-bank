package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	userChoice := getUserChoice()

	if userChoice == 1 {
		fmt.Println("Your balance is: USD", accountBalance)
	} else if userChoice == 2 {
		accountBalance = getUserDeposit(accountBalance)
		fmt.Print("Your balance is now USD", accountBalance) 
	} else if userChoice == 3 {
		accountBalance = getUserWithdraw(accountBalance)
		fmt.Print("Your balance is now USD", accountBalance) 
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

func getUserDeposit(accountBalance float64) (float64) {
	var depositAmount float64
	var confirmDeposit string

	fmt.Print("How much do you want to deposit?: ")
	fmt.Scan(&depositAmount)
	fmt.Print("The value of USD", depositAmount, " will be deposited into your account. Do you wish to continue? (Y/N)")
	fmt.Scan(&confirmDeposit)

	if (confirmDeposit == "Y" || confirmDeposit == "y") {
		accountBalance += depositAmount		
	} else if (confirmDeposit == "N" || confirmDeposit == "n") {
		fmt.Print("Deposit cancelled")
	}

	return accountBalance
}

func getUserWithdraw(accountBalance float64) (float64) {
	var withdrawAmount float64
	var confirmWithdraw string

	fmt.Print("How much do you want to withdraw?: ")
	fmt.Scan(&withdrawAmount)
	fmt.Print("The value of USD", withdrawAmount, " will be withdrawn into your account. Do you wish to continue? (Y/N)")
	fmt.Scan(&confirmWithdraw)

	if (confirmWithdraw == "Y" || confirmWithdraw == "y") {
		accountBalance -= withdrawAmount		
	} else if (confirmWithdraw == "N" || confirmWithdraw == "n") {
		fmt.Print("Withdraw cancelled")
	}

	return accountBalance
}