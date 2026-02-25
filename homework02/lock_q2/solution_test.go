package main

import "testing"

func TestRunAtomicCounter(t *testing.T) {
	got := RunAtomicCounter(10, 1000)
	var want int64 = 10000

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestRunAtomicCounterZeroWorkers(t *testing.T) {
	got := RunAtomicCounter(0, 1000)
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}
