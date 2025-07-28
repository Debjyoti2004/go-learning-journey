package main
import "fmt"


func main(){

	type User struct{
		FirstName  string
		LastName   string
		Email      string
		Age        int
	}
	var user []User

	user = append(user,User{"Debjyoti","Shit","debjyoti@gmail.com",20})
	user = append(user,User{"Debjyoti1","shit1","debjyoti1@gmail.com",21})
	user = append(user,User{"Debjyoti2","shit2","debjyoti2@gmail.com",22})
	user = append(user,User{"Debjyoti3","shit3","debjyoti3@gmail.com",23})
	user = append(user,User{"Debjyoti4","shit4","debjyoti4@gmail.com",24})
	user = append(user,User{"Debjyoti5","shit5","debjyoti5@gmail.com",25})
	user = append(user,User{"Debjyoti6","shit6","debjyoti6@gmail.com",26})
	

	for _, user := range user{
		fmt.Println(user.FirstName,user.LastName,user.Email,user.Age)
	}
	
}