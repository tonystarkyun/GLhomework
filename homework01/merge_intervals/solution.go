package mergeintervals

import "sort"

// Merge merges all overlapping intervals.
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sortedIntervals := make([][]int, len(intervals))
	for i := range intervals {
		sortedIntervals[i] = []int{intervals[i][0], intervals[i][1]}
	}

	sort.Slice(sortedIntervals, func(i, j int) bool {
		if sortedIntervals[i][0] == sortedIntervals[j][0] {
			return sortedIntervals[i][1] < sortedIntervals[j][1]
		}
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	merged := [][]int{sortedIntervals[0]}
	for i := 1; i < len(sortedIntervals); i++ {
		last := merged[len(merged)-1]
		current := sortedIntervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
			continue
		}

		merged = append(merged, current)
	}

	return merged
}
