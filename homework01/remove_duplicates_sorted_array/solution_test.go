package removeduplicatessortedarray

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantK    int
		wantNums []int
	}{
		{
			name:     "basic case",
			nums:     []int{1, 1, 2},
			wantK:    2,
			wantNums: []int{1, 2},
		},
		{
			name:     "long sample",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantK:    5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "empty input",
			nums:     []int{},
			wantK:    0,
			wantNums: []int{},
		},
		{
			name:     "single element",
			nums:     []int{7},
			wantK:    1,
			wantNums: []int{7},
		},
		{
			name:     "with negative values",
			nums:     []int{-1, -1, 0, 0, 0, 3, 3},
			wantK:    3,
			wantNums: []int{-1, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]int(nil), tt.nums...)
			gotK := RemoveDuplicates(input)
			if gotK != tt.wantK {
				t.Fatalf("RemoveDuplicates(%v) returned k=%d, want %d", tt.nums, gotK, tt.wantK)
			}
			if !sameInts(input[:gotK], tt.wantNums) {
				t.Fatalf("RemoveDuplicates(%v) returned nums[:k]=%v, want %v", tt.nums, input[:gotK], tt.wantNums)
			}
		})
	}
}

func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
