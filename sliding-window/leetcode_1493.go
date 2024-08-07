package main

import "fmt"

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
func longestSubarray(nums []int) int {
	res := 0
	start := 0
	deleteCount := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			deleteCount++
		}
		for deleteCount > 1 {
			if nums[start] == 0 {
				deleteCount--
			}
			start++
		}
		res = max(res, end-start)
	}
	return res
}
