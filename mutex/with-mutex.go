package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock the mutex before modifying the count
	// and unlock it after the modification is done
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	CounterNum := Counter{count: 0}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go CounterNum.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter value:", CounterNum.count)
}
