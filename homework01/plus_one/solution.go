package plusone

// PlusOne adds one to a non-negative integer represented by digits.
func PlusOne(digits []int) []int {
	result := make([]int, len(digits))
	copy(result, digits)

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] < 9 {
			result[i]++
			return result
		}
		result[i] = 0
	}

	return append([]int{1}, result...)
}
