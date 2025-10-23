package main

import "fmt"

func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("Task is being executed")
}

func main() {
	done := make(chan bool)
	go task(done)
	<-done
	fmt.Println("Task completed signal received")
}
