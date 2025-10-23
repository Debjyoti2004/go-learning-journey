package main

func main() {
	messageCh := make(chan string) // Create a channel to communicate string messages

	go func() {
		messageCh <- "Hello from the goroutine!" // Send a message to the channel
	}()

	message := <-messageCh // Receive the message from the channel
	println(message)       // Print the received message
}
