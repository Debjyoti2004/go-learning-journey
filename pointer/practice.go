package main

import "fmt"

// pass by value
func sum(a int, b int) {
	result := a + b
	fmt.Println("The sum is:", result)
}

// pass by reference

func updateTheValue(number *int) {
	*number = 100

	fmt.Println("The value of number is inside the function:", *number)
}

func main() {
	fmt.Println("--------------------------")
	number := 10
	fmt.Println("The value of number is Before to pass to the function:", number)
	updateTheValue(&number)
	fmt.Println("The value of number is After to pass to the function:", number)
	fmt.Println("--------------------------")
	fmt.Println("Pass by value example:")
	num1 := 5
	num2 := 10
	fmt.Println("The values are:", num1, "and", num2)
	sum(num1, num2)
	fmt.Println("The values after function call are still:", num1, "and", num2)
	fmt.Println("--------------------------")
	fmt.Println("Pass by reference example:")
	check := 12
	fmt.Println("Address of check is:", &check)
	address := &check
	fmt.Println("The value at address is:", *address)
	fmt.Println("--------------------------")
}
