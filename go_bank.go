package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	
	for {
		userChoice := getUserChoice()

		if userChoice == 1 {
			fmt.Println("Your balance is: USD", accountBalance)
			continue
		} else if userChoice == 2 {
			accountBalance = getUserDeposit(accountBalance)
			fmt.Println("Your balance is now USD", accountBalance) 
			continue
		} else if userChoice == 3 {
			accountBalance = getUserWithdraw(accountBalance)
			fmt.Println("Your balance is now USD", accountBalance) 
			continue
		} else if userChoice == 4 {
			fmt.Println("Thank you for choosing Go Bank!")
			break
		} else {
			fmt.Println("Enter a valid input!")
			continue
		}
	}
}

func getUserChoice() int {
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

func getUserDeposit(accountBalance float64) float64 {
	var depositAmount float64
	var confirmDeposit string

	fmt.Print("How much do you want to deposit?: ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		return accountBalance
	}

	fmt.Print("The value of USD", depositAmount, " will be deposited into your account. Do you wish to continue? (Y/N)")
	fmt.Scan(&confirmDeposit)

	if (confirmDeposit == "Y" || confirmDeposit == "y") {
		accountBalance += depositAmount		
	} else if (confirmDeposit == "N" || confirmDeposit == "n") {
		fmt.Println("Deposit cancelled")
	}

	return accountBalance
}

func getUserWithdraw(accountBalance float64) float64 {
	var withdrawAmount float64
	var confirmWithdraw string

	fmt.Print("How much do you want to withdraw?: ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		return accountBalance
	}
	
	if accountBalance < withdrawAmount {
		fmt.Println("Your withdraw amount can't exceed your balance of USD", accountBalance)
		return accountBalance
	}

	fmt.Print("The value of USD", withdrawAmount, " will be withdrawn from your account. Do you wish to continue? (Y/N)")
	fmt.Scan(&confirmWithdraw)

	if (confirmWithdraw == "Y" || confirmWithdraw == "y") {
		accountBalance -= withdrawAmount		
	} else if (confirmWithdraw == "N" || confirmWithdraw == "n") {
		fmt.Println("Withdraw cancelled")
	}

	return accountBalance
}