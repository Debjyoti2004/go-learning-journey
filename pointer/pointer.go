// Explore on pointer

package main

import "fmt"

func main(){
	var color string = "red"
	var prt *string   // there we can say it is a box that can hold the address of an string.
	prt = &color  // prt stores the address of color using &color
	fmt.Println("The value of prt is:",prt)  // this give us the address of the color.
	fmt.Println("The value of *prt is:",*prt) // the give the color value.
}