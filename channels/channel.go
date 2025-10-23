package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numberChannel(numch chan int) {
	for num := range numch {
		fmt.Println("Printed number:", num)
		time.Sleep(time.Second)
	}
}
func main() {
	numCh := make(chan int)
	go numberChannel(numCh)

	for {
		numCh <- rand.Intn(100) // Send a random number to the channel

	}

}
