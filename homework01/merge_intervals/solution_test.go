package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "sample one",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "touching intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "same end",
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
		{
			name:      "no overlap",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
