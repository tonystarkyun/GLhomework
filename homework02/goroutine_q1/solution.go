package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type NumberMessage struct {
	Kind  string
	Value int
}

// CollectOddEven starts two goroutines and collects odd/even numbers up to max.
func CollectOddEven(max int) []NumberMessage {
	if max < 1 {
		return nil
	}

	out := make(chan NumberMessage, max)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= max; i += 2 {
			out <- NumberMessage{Kind: "odd", Value: i}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= max; i += 2 {
			out <- NumberMessage{Kind: "even", Value: i}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	result := make([]NumberMessage, 0, max)
	for msg := range out {
		result = append(result, msg)
	}
	return result
}

func PrintOddEven(max int, w io.Writer) {
	for _, msg := range CollectOddEven(max) {
		fmt.Fprintf(w, "%s:%d\n", msg.Kind, msg.Value)
	}
}

func main() {
	PrintOddEven(10, os.Stdout)
}
