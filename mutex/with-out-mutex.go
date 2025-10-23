package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.count++
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
