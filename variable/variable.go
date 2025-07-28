// There we declare the variable
package main
import "fmt"

func main() {
	// For String
	var whatIsYourName string
	whatIsYourName = "Debjyoti"
	fmt.Println("Your Name is:",whatIsYourName)
	whatIsYourName = "Debjyoti Shit"
	fmt.Println("Your Name is:",whatIsYourName)

	// For Integer
	var age int 
	age = 21   // pointer take the value from the address of the variable
	fmt.Println("Your Age is:",age)

	// For Float 

	var pi float32
	pi = 3.14
	fmt.Println("The value of pi is:",pi)

	// For Boolean
	var isStudent bool
	isStudent = true
	fmt.Println("Are you a Student:",isStudent)

	// For Array
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("The value of numbers is:",numbers)

	// For Map
	studentGrades := map[string]any{
		"Name": "Debjyoti",
		"Age": 21,
		"IsStudent": true,
	}
	fmt.Println("The value of studentGrades is:",studentGrades)

	// For Slice
	numbersSlice := []int{1,2,3,4,5}
	fmt.Println("The value of numbersSlice is:",numbersSlice)

	// For Struct
	type Person struct {
		Name string
		Age int
		IsStudent bool
	}
	person := []Person{
		{
			Name: "Debjyoti",
			Age: 21,
			IsStudent: true,
		},
		{
			Name: "Debjyoti Shit",
			Age: 21,
			IsStudent: true,
		},
	}
	fmt.Println("The value of person 1 is:",person[0])
	fmt.Println("The value of person 2 is:",person[1])
	fmt.Println("The Name of person 2 is:",person[1].Name)
	
	// For Pointer
	var ptr *int
	ptr = &age
	fmt.Println("The value of ptr is:",ptr)
	fmt.Println("The value of *ptr is:",*ptr)


	city := "kolkata"

	var ptrCity *string
	ptrCity = &city
	fmt.Println("The value of ptrCity is:",ptrCity)
	fmt.Println("The value of *ptrCity is:",*ptrCity)

	// For Constant
	const dateOfBirth = 2004
	// dateOfBirth = 2005  // This will give an error because const is a constant and cannot be changed
	fmt.Println("The value of dateOfBirth is:",dateOfBirth)

	// For Constant with type
	const dateOfBirthCopy int = 2004
	fmt.Println("The value of dateOfBirthCopy is:",dateOfBirthCopy)

	
}