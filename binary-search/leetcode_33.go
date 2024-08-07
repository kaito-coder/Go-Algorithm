package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	result := search(arr, target)
	fmt.Println(result)
}
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
	}
	return -1
}
