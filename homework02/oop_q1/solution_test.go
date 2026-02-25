package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestRectangleImplementsShape(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	if !almostEqual(rect.Area(), 15) {
		t.Fatalf("expected area 15, got %f", rect.Area())
	}
	if !almostEqual(rect.Perimeter(), 16) {
		t.Fatalf("expected perimeter 16, got %f", rect.Perimeter())
	}
}

func TestCircleImplementsShape(t *testing.T) {
	circle := Circle{Radius: 2}
	if !almostEqual(circle.Area(), 4*math.Pi) {
		t.Fatalf("unexpected circle area: %f", circle.Area())
	}
	if !almostEqual(circle.Perimeter(), 4*math.Pi) {
		t.Fatalf("unexpected circle perimeter: %f", circle.Perimeter())
	}
}
