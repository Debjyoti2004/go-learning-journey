package main

import "fmt"

func main(){
	var number [5]int

	number[0] = 1
	number[1] = 2
	number[2] = 3
	number[3] = 4
	number[4] = 5

	fmt.Println("The value of number is:",number)

	number2 := [6]int{1,2,3,4,5}
	fmt.Println("The value of number2 is:",number2)

	// For get the length of the array
	fmt.Println("The length of the array is:",len(number2))

	// For get the capacity of the array
	fmt.Println("The capacity of the array is:",cap(number2))
	
}