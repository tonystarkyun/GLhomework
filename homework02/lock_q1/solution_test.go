package main

import "testing"

func TestRunMutexCounter(t *testing.T) {
	got := RunMutexCounter(10, 1000)
	want := 10000

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestRunMutexCounterZeroWorkers(t *testing.T) {
	got := RunMutexCounter(0, 1000)
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}
