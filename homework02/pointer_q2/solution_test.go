package main

import (
	"reflect"
	"testing"
)

func TestDoubleSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	DoubleSlice(&nums)

	want := []int{2, 4, 6, 8}
	if !reflect.DeepEqual(nums, want) {
		t.Fatalf("expected %v, got %v", want, nums)
	}
}

func TestDoubleSliceEmpty(t *testing.T) {
	nums := []int{}
	DoubleSlice(&nums)

	if len(nums) != 0 {
		t.Fatalf("expected empty slice, got %v", nums)
	}
}

func TestDoubleSliceNilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("DoubleSlice should not panic, got %v", r)
		}
	}()

	DoubleSlice(nil)
}
