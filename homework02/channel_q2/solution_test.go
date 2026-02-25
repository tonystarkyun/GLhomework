package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestRunBufferedChannelDemo(t *testing.T) {
	var buf bytes.Buffer
	got := RunBufferedChannelDemo(100, 16, &buf)

	if len(got) != 100 {
		t.Fatalf("expected 100 numbers, got %d", len(got))
	}

	wantFirstTen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(got[:10], wantFirstTen) {
		t.Fatalf("unexpected first 10 numbers: %v", got[:10])
	}
	if got[99] != 100 {
		t.Fatalf("expected last number to be 100, got %d", got[99])
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 100 {
		t.Fatalf("expected 100 printed lines, got %d", len(lines))
	}
}

func TestRunBufferedChannelDemoWithInvalidBuffer(t *testing.T) {
	var buf bytes.Buffer
	got := RunBufferedChannelDemo(5, 0, &buf)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
