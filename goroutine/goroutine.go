package main

import (
	"fmt"
	"time"
)

func TaskID(ID int) {
	fmt.Println("TaskID", ID+1)
}

func main() {
	start := time.Now()
	for i := 0; i < 15; i++ {
		go TaskID(i)
	}
	end := time.Now()
	fmt.Println("Time taken:", end.Sub(start))
	time.Sleep(2 * time.Second)
}
