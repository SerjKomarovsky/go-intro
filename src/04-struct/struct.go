package main

import (
	"example/struct/user"
	"fmt"
)

func getUserData(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)

	return input
}

func main() {
	appUser, err := user.New(
		getUserData("First name: "),
		getUserData("Last name: "),
		getUserData("Birthday (MM/DD/YYYY):"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// passing pointer to avoid creating of one more variable (inside printUserDetails)
	appUser.PrintUserDetails()
	appUser.ClearUserName()
	appUser.PrintUserDetails()
}
