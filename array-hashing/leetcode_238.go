package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
func productExceptSelf(nums []int) []int {
	size := len(nums)

	res := make([]int, size)

	left := 1 // product all left of nums[i]

	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}

	right := 1 // product all right of nums[i]
	for i := size - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

