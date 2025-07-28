package main

import "fmt"


type person struct{
	name string
	age int
}

type employee struct{
	person
	salary float64
}

func main(){

	employee1 := employee{
		person: person{
			name: "Debjyoti",
			age: 20,
		},
		salary: 100000,
	}

	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.salary)

}