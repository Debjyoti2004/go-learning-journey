package main
import "fmt"

type User struct {
	Name string
}

func main(){

	var myUser User
	myUser.Name = "Debjyoti"
	fmt.Println("The value of myUser is:",myUser)

	myUser2 := User{
		Name: "Debjyoti Shit",
	}
	fmt.Println("The value of myUser2 is:",myUser2)

	// For pointer with function
	var myUserPointer *User
	myUserPointer = &myUser
	fmt.Println("The value of myUserPointer is:",*myUserPointer)
	
}