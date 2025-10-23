package main

import (
	"fmt"
	"time"
)

func emailQueue(emailCh <-chan string, done chan<- bool) { // Receive-only channel there we can't send data only we can receive

	defer func() { done <- true }() // Signal completion when done

	// emailCh <- "something@gmail.com" // This will cause a compile-time error

	for email := range emailCh { // Receive emails from the channel
		fmt.Println("Email ID is : ", email)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	emailCh := make(chan string, 10) // Create a buffered channel with capacity of 10
	done := make(chan bool)

	go emailQueue(emailCh, done)

	for i := 1; i <= 10; i++ {
		emailCh <- fmt.Sprintf("debjyoti%d@gmail.com", i)
	}
	fmt.Println("All email IDs sent to channel")

	close(emailCh) // Close the channel after sending all emails, else the goroutine will block forever because of range loop
	<-done
}
