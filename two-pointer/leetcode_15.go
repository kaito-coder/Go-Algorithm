package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	leftPointer := 0
	rightPointer := len(nums) - 1
	res := [][]int{}
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		for leftPointer < rightPointer {
			sum := nums[leftPointer] + nums[rightPointer] + nums[i]
			if sum > 0 {
				rightPointer--
			} else if sum < 0 {
				leftPointer++
			} else {
				res = append(res, []int{nums[i], nums[leftPointer], nums[rightPointer]})
				rightPointer--
				leftPointer++
			}
		}

	}
	return res
}
