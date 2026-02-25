package main

import (
	"fmt"
	"io"
	"os"
)

func producer(max int, out chan<- int) {
	defer close(out)
	for i := 1; i <= max; i++ {
		out <- i
	}
}

func consumer(in <-chan int, w io.Writer, done chan<- []int) {
	received := make([]int, 0)
	for n := range in {
		fmt.Fprintln(w, n)
		received = append(received, n)
	}
	done <- received
}

// RunChannelDemo starts producer/consumer goroutines and returns consumed values.
func RunChannelDemo(max int, w io.Writer) []int {
	numbers := make(chan int)
	done := make(chan []int, 1)

	go producer(max, numbers)
	go consumer(numbers, w, done)

	return <-done
}

func main() {
	RunChannelDemo(10, os.Stdout)
}
