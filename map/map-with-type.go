package main
import "fmt"


type User struct {
	Name string
}

func main(){
	myMap := make(map[string]User)
	myMap["User1"] = User{
		Name: "Debjyoti",
	}

	fmt.Println("The value of myMap is:",myMap["User1"].Name)
}