package main

import (
	"example/bank/fileops"
	"fmt"
)

const BALANCE_FILENAME = "balance.txt"

func main() {
	var choice int
	accountBalance, accountError := fileops.LoadFloatFromFile(BALANCE_FILENAME)

    if accountError != nil {
        fmt.Println(accountError)
        fmt.Println("Exiting...")
        return
    }

	fmt.Println("Welcome to Go Bank!")
	for choice != 4 {
		chooseOperation(&choice)

		switch choice {
			case 1:
				fmt.Printf("Your balance is %.2f$\n", accountBalance)
			case 2:
				depositMoney(&accountBalance)
			case 3:
				withdrawMoney(&accountBalance)
			case 4:
				return
			default:
				fmt.Println("invalid operation")
		}
	}

	fileops.WriteFloatToFile(accountBalance, BALANCE_FILENAME)
	fmt.Println(("Thanks for choosing our bank!"))
}

func withdrawMoney(accountBalance *float64) {
	var withdrawValue float64
	fmt.Print("Enter the amount you want to withdraw: ")
	fmt.Scan(&withdrawValue)

	if withdrawValue <= 0 {
		fmt.Print("Withdrawal value must be greater than 0")
		return
	}

	if withdrawValue > *accountBalance {
		fmt.Print("Withdrawal value must be less than current balance")
		return
	}

	*accountBalance -= withdrawValue
	fmt.Printf("You've successfully withdrawed %.2f$! Current balance: %.2f$\n", withdrawValue, *accountBalance)
}

func depositMoney(accountBalance *float64) {
	var depositValue float64
	fmt.Print("Enter the amount you want to deposit, $: ")
	fmt.Scan(&depositValue)

	if depositValue <= 0 {
		fmt.Print("Deposit value must be greater than 0")
		return
	}
	*accountBalance += depositValue
	fmt.Printf("Balance updated! New amount: %.2f$\n", *accountBalance)
}
