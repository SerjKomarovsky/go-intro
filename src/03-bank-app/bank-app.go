package main

import "fmt"

func main() {
	var choice int
	var accountBalance float64 = 1000

	chooseOperation(&choice)
	fmt.Printf("Your choice: %v\n", choice)

	wantsCheckBalance := choice == 1
	wantsDepositMoney := choice == 2
	wantsWithdrawMoney := choice == 3

	if wantsCheckBalance {
		fmt.Printf("Your balance is %.2f$\n", accountBalance)
	} else if wantsDepositMoney {
		var depositValue float64
		fmt.Print("Enter the amount you want to deposit, $: ")
		fmt.Scan(&depositValue)
		accountBalance += depositValue
		fmt.Printf("Balance updated! New amount: %.2f$\n", accountBalance)
	} else if wantsWithdrawMoney {
		var withdrawValue float64
		fmt.Print("Enter the amount you want to withdraw: ")
		fmt.Scan(&withdrawValue)
		accountBalance -= withdrawValue
		fmt.Printf("You've successfully withdrawed %.2f$! Current balance: %.2f$\n", withdrawValue, accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}

func chooseOperation(choice *int) {
	fmt.Print(`
    Welcome to Go Bank!
    What do you want to do?
    1. Check balance
    2. Deposit money
    3. Withdraw money
    4. Exit
    `)

	fmt.Scan(choice)
}
