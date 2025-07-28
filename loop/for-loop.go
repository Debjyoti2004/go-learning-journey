package main
import "fmt"


func main(){
	for i :=0; i<=10; i++{
		fmt.Println("The value of i is:",i)
	}


	animal := [] string{"Dog","Cat","Bird"}

	for i :=0; i<len(animal); i++{
		fmt.Println("The value of animal is:",animal[i])
	}

	for index, value := range animal{
		fmt.Println("The value of animal is:",value,"and the index is:",index)
	}

	for _, value := range animal{
		fmt.Println("The value of animal is:",value)
	}
	
	for index := range animal{
		fmt.Println("The index of animal is:",index)
	}

	for index, _ := range animal{
		fmt.Println("The index of animal is:",index)
	}
}