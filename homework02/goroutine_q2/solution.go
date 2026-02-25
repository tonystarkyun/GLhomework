package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

type Task struct {
	Name string
	Run  func(ctx context.Context) error
}

type TaskResult struct {
	Name     string
	Duration time.Duration
	Err      error
}

// RunTasksConcurrently runs all tasks concurrently and records execution time for each task.
func RunTasksConcurrently(ctx context.Context, tasks []Task) []TaskResult {
	if len(tasks) == 0 {
		return nil
	}

	results := make(chan TaskResult, len(tasks))
	var wg sync.WaitGroup

	for _, task := range tasks {
		t := task
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			err := t.Run(ctx)
			results <- TaskResult{
				Name:     t.Name,
				Duration: time.Since(start),
				Err:      err,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	collected := make([]TaskResult, 0, len(tasks))
	for result := range results {
		collected = append(collected, result)
	}

	sort.Slice(collected, func(i, j int) bool {
		return collected[i].Name < collected[j].Name
	})
	return collected
}

func SleepTask(name string, d time.Duration, err error) Task {
	return Task{
		Name: name,
		Run: func(ctx context.Context) error {
			select {
			case <-time.After(d):
				return err
			case <-ctx.Done():
				return ctx.Err()
			}
		},
	}
}

func main() {
	tasks := []Task{
		SleepTask("task-A", 120*time.Millisecond, nil),
		SleepTask("task-B", 80*time.Millisecond, nil),
		SleepTask("task-C", 100*time.Millisecond, errors.New("simulated failure")),
	}

	results := RunTasksConcurrently(context.Background(), tasks)
	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("%s took %v, error: %v\n", result.Name, result.Duration, result.Err)
			continue
		}
		fmt.Printf("%s took %v\n", result.Name, result.Duration)
	}
}
