package main

import "fmt"

// DoubleSlice multiplies each element in nums by 2.
func DoubleSlice(nums *[]int) {
	if nums == nil {
		return
	}
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	DoubleSlice(&nums)
	fmt.Println(nums)
}
