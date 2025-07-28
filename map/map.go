package main
import "fmt"

func main(){
	studentGrades := make(map[string]any)
	studentGrades["Debjyoti"] = 90
	studentGrades["Debjyoti Shit"] = 95		
	studentGrades["Debjyoti Shit 2"] = 98
	studentGrades["Debjyoti Shit 3"] = 99
	studentGrades["Debjyoti Shit 4"] = 100
	fmt.Println("The value of studentGrades is:",studentGrades)

	// For delete the value from the map
	delete(studentGrades, "Debjyoti Shit 2")
	fmt.Println("The value of studentGrades is:",studentGrades)

	// For update the value from the map
	studentGrades["Debjyoti"] = 95
	fmt.Println("The value of studentGrades is:",studentGrades)

	// For get the value from the map
	fmt.Println("The value of studentGrades is:",studentGrades["Debjyoti"])

	// For get the value from the map
	fmt.Println("The value of studentGrades is:",studentGrades["Debjyoti Shit"])

	// For get the value from the map
	fmt.Println("The value of studentGrades is:",studentGrades["Debjyoti Shit 2"])
	
}