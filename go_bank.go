package main

import (
	"fmt"
	"github.com/mtravassos96/go-bank/fileops"
)


const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't continue the application.")
	}

	fmt.Println("Welcome to Go Bank!")
	
	for {
		userChoice := getUserChoice()

		switch userChoice {
		case 1:
			fmt.Println("Your balance is: USD", accountBalance)
		case 2:
			accountBalance = getUserDeposit(accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your balance is now USD", accountBalance)
		case 3:
			accountBalance = getUserWithdraw(accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your balance is now USD", accountBalance)
		case 4:
			fmt.Println("Thank you for choosing Go Bank!")
			return
		default:
			fmt.Println("Enter a valid input!")
		}
	}
	
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