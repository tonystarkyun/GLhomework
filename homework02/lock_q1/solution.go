package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func RunMutexCounter(workers, increments int) int {
	counter := &MutexCounter{}
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	return counter.Value()
}

func main() {
	fmt.Println(RunMutexCounter(10, 1000))
}
