package main

import "fmt"

func main() {
	emailCh := make(chan string, 3) // Create a buffered channel with capacity of 3

	emailCh <- "debjyoti1@gmail.com"
	emailCh <- "debjyoti2@gmail.com"
	emailCh <- "debjyoti3@gmail.com"
	//emailCh <- "debjyoti4@gmail.com"

	fmt.Println("Email ID is : ", <-emailCh)
	fmt.Println("Email ID is : ", <-emailCh)
	fmt.Println("Email ID is : ", <-emailCh)
	//fmt.Println("Email ID is : ", <-emailCh)
}
