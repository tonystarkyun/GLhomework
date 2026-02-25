package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func RunAtomicCounter(workers, increments int) int64 {
	counter := &AtomicCounter{}
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
	fmt.Println(RunAtomicCounter(10, 1000))
}
