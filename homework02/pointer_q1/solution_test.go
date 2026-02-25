package main

import "testing"

func TestAddTen(t *testing.T) {
	value := 7
	AddTen(&value)

	if value != 17 {
		t.Fatalf("expected 17, got %d", value)
	}
}

func TestAddTenNilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("AddTen should not panic, got %v", r)
		}
	}()

	AddTen(nil)
}

func TestSolve(t *testing.T) {
	got := Solve(20)
	if got != 30 {
		t.Fatalf("expected 30, got %d", got)
	}
}
