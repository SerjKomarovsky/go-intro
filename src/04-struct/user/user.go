package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

// receiver argument
func (u User) PrintUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthday)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("first name, last name and birthday are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}
