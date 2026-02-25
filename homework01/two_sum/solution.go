package twosum

// TwoSum returns the indices of the two numbers that add up to target.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		need := target - num
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}
