package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}
	k := 2
	fmt.Println(nums, k)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	// Print the sorted slice
	fmt.Println(nums) // Output: [5 5 5 5 4 4 4 3 3 3 2 2 2 2 1 1 1 1]

}
func maxOperations(nums []int, k int) int {
	l, r := 0, len(nums)-1
	count := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	for l < r {
		sum := nums[l] + nums[r]
		if sum > k {
			r--
		} else if sum < k {
			l++
		} else {
			count++
			l++
			r--
		}
	}
	return count
}
