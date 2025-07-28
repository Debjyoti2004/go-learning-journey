package main
import "fmt"

type User struct {
	Name string
}

func  (u User) changeName() string{
	return u.Name
}

func main(){

	var myUser User
	myUser.Name = "Debjyoti"

	myUser2 := User{
		Name: "Debjyoti Shit",
	}
	fmt.Println("The value of myUser is:",myUser.changeName())
	fmt.Println("The value of myUser2 is:",myUser2.changeName())

	// For pointer with function
	var myUserPointer *User
	myUserPointer = &myUser
	fmt.Println("The value of myUserPointer is:",*myUserPointer)
	
}