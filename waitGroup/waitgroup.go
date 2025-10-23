package main

import (
	"fmt"
	"sync"
)

func TaskID(ID int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("TaskID", ID+1)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go TaskID(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
