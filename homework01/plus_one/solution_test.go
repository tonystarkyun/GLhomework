package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "simple increment",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "carry in tail",
			digits: []int{4, 3, 2, 9},
			want:   []int{4, 3, 3, 0},
		},
		{
			name:   "single nine",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "all nines",
			digits: []int{9, 9, 9},
			want:   []int{1, 0, 0, 0},
		},
		{
			name:   "zero value",
			digits: []int{0},
			want:   []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("PlusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
