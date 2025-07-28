package main
import "fmt"


func updateValue(number *int){
	*number = 100

	fmt.Println("The value of number is inside the function:",*number)
}

func main(){
	number := 10
	fmt.Println("The value of number is Before to pass to the function:",number)
	updateValue(&number)
	fmt.Println("The value of number is After to pass to the function:",number)
}