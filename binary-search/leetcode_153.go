package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 5, 1, 2}

	fmt.Println(findMin(arr))
}
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return nums[l]

}
