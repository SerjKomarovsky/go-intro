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

    admin := user.NewAdmin("test@example.com", "test123")

    admin.PrintUserDetails()
    admin.ClearUserName()
    admin.PrintUserDetails()

	appUser.PrintUserDetails()
	appUser.ClearUserName()
	appUser.PrintUserDetails()
}
