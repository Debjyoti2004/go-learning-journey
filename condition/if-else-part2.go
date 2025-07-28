package main

import "fmt"

func main(){
	var age int

	fmt.Println("Enter your age:")
	_, err:=fmt.Scanln(&age)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if age >=18 && age <= 60 {
		fmt.Println("You are an adult")
	} else if age >= 13 && age < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}
}