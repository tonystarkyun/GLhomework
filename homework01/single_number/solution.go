package singlenumber

// SingleNumber returns the element that appears exactly once.
func SingleNumber(nums []int) int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	for _, num := range nums {
		if counts[num] == 1 {
			return num
		}
	}

	return 0
}
