package main

import "fmt"


func main() {
	var choice int
	var accountBalance float64 = 1000

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

    fmt.Println(("Thanks for choosing our bank!"))
}

func chooseOperation(choice *int) {
	fmt.Print(`
    What do you want to do?
    1. Check balance
    2. Deposit money
    3. Withdraw money
    4. Exit
    `)

	fmt.Scan(choice)
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
