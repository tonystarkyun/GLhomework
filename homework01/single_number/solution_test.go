package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic case",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "middle unique",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "negative number",
			nums: []int{-1, -1, -2},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			if got != tt.want {
				t.Fatalf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
