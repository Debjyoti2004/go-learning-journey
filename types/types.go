package main

import (
	"fmt"
	"time"
)

type User struct {
	Name        string
	Age         int
	Email       string
	Password    string
	IsActive    bool
	dateOfBirth time.Time
}

func main() {
	person := User{
		Name:        "Debjyoti",
		Age:         21,
		Email:       "debjyoti27@gmail.com",
		Password:    "123456",
		IsActive:    true,
		dateOfBirth: time.Date(2004, 4, 27, 0, 0, 0, 0, time.UTC),
	}

	// Print the full struct but format the date separately
	fmt.Printf("The value of person is:\nName: %s\nAge: %d\nEmail: %s\nPassword: %s\nIsActive: %t\ndateOfBirth: %s\n",
		person.Name, person.Age, person.Email, person.Password, person.IsActive, person.dateOfBirth.Format("02-Jan-2006"))
}
