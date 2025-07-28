package main
import "fmt"

func main(){
	var name []string  // this is slice, not array

	name = append(name, "Debjyoti")
	name = append(name, "Debjyoti1")
	name = append(name, "Debjyoti2")
	name = append(name, "Debjyoti3")
	fmt.Println(name)

	var name2 [5]string
	name2[0] = "Debjyoti"
	name2[1] = "Debjyoti1"
	name2[2] = "Debjyoti2"
	name2[3] = "Debjyoti3"
	fmt.Println(name2)

	// For get the length of the array
	fmt.Println("The length of the array is:",len(name2))
	
}