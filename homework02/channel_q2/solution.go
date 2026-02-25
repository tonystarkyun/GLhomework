package main

import (
	"fmt"
	"io"
	"os"
)

func bufferedProducer(total int, out chan<- int) {
	defer close(out)
	for i := 1; i <= total; i++ {
		out <- i
	}
}

func bufferedConsumer(in <-chan int, w io.Writer, done chan<- []int) {
	received := make([]int, 0)
	for n := range in {
		fmt.Fprintln(w, n)
		received = append(received, n)
	}
	done <- received
}

// RunBufferedChannelDemo starts producer/consumer goroutines with a buffered channel.
func RunBufferedChannelDemo(total, buffer int, w io.Writer) []int {
	if buffer <= 0 {
		buffer = 1
	}

	ch := make(chan int, buffer)
	done := make(chan []int, 1)

	go bufferedProducer(total, ch)
	go bufferedConsumer(ch, w, done)

	return <-done
}

func main() {
	RunBufferedChannelDemo(100, 16, os.Stdout)
}
