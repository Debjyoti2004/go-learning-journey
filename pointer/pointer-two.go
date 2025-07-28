package main 

import "fmt"

func main(){
	var color string = "red"
	fmt.Println("The value of color is:",color)
	changeColor(&color)
	fmt.Println("After the value of color is:",color)
}

func changeColor(color *string){
	*color = "blue"
}