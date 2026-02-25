package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestCollectOddEvenContainsAllNumbers(t *testing.T) {
	msgs := CollectOddEven(10)
	if len(msgs) != 10 {
		t.Fatalf("expected 10 messages, got %d", len(msgs))
	}

	seen := make(map[int]string, 10)
	for _, msg := range msgs {
		seen[msg.Value] = msg.Kind
	}

	for i := 1; i <= 10; i++ {
		kind, ok := seen[i]
		if !ok {
			t.Fatalf("number %d is missing", i)
		}

		if i%2 == 1 && kind != "odd" {
			t.Fatalf("number %d should be odd, got %s", i, kind)
		}
		if i%2 == 0 && kind != "even" {
			t.Fatalf("number %d should be even, got %s", i, kind)
		}
	}
}

func TestPrintOddEven(t *testing.T) {
	var buf bytes.Buffer
	PrintOddEven(10, &buf)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 10 {
		t.Fatalf("expected 10 lines, got %d", len(lines))
	}

	seen := make(map[int]string, 10)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			t.Fatalf("unexpected line format: %q", line)
		}
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			t.Fatalf("invalid number in line %q: %v", line, err)
		}
		seen[num] = parts[0]
	}

	for i := 1; i <= 10; i++ {
		kind, ok := seen[i]
		if !ok {
			t.Fatalf("number %d is missing from output", i)
		}
		if i%2 == 1 && kind != "odd" {
			t.Fatalf("number %d should be odd, got %s", i, kind)
		}
		if i%2 == 0 && kind != "even" {
			t.Fatalf("number %d should be even, got %s", i, kind)
		}
	}
}
