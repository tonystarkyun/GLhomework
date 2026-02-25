package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunTasksConcurrently(t *testing.T) {
	taskErr := errors.New("task failed")
	tasks := []Task{
		SleepTask("task-A", 80*time.Millisecond, nil),
		SleepTask("task-B", 120*time.Millisecond, nil),
		SleepTask("task-C", 60*time.Millisecond, taskErr),
	}

	start := time.Now()
	results := RunTasksConcurrently(context.Background(), tasks)
	elapsed := time.Since(start)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	resultMap := make(map[string]TaskResult, 3)
	for _, result := range results {
		resultMap[result.Name] = result
	}

	if resultMap["task-A"].Duration < 80*time.Millisecond {
		t.Fatalf("task-A duration too short: %v", resultMap["task-A"].Duration)
	}
	if resultMap["task-B"].Duration < 120*time.Millisecond {
		t.Fatalf("task-B duration too short: %v", resultMap["task-B"].Duration)
	}
	if resultMap["task-C"].Duration < 60*time.Millisecond {
		t.Fatalf("task-C duration too short: %v", resultMap["task-C"].Duration)
	}

	if !errors.Is(resultMap["task-C"].Err, taskErr) {
		t.Fatalf("expected task-C error %v, got %v", taskErr, resultMap["task-C"].Err)
	}
	if resultMap["task-A"].Err != nil || resultMap["task-B"].Err != nil {
		t.Fatalf("unexpected errors in successful tasks")
	}

	// Sequential duration would be around 260ms; concurrent run should be much lower.
	if elapsed >= 240*time.Millisecond {
		t.Fatalf("tasks did not run concurrently enough, elapsed: %v", elapsed)
	}
}
