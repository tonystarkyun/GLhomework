package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestRunChannelDemo(t *testing.T) {
	var buf bytes.Buffer
	got := RunChannelDemo(10, &buf)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 10 {
		t.Fatalf("expected 10 printed lines, got %d", len(lines))
	}
}

func TestRunChannelDemoZero(t *testing.T) {
	var buf bytes.Buffer
	got := RunChannelDemo(0, &buf)

	if len(got) != 0 {
		t.Fatalf("expected empty result, got %v", got)
	}
	if buf.Len() != 0 {
		t.Fatalf("expected empty output, got %q", buf.String())
	}
}
