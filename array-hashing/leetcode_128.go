package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	// 1 ,2 ,3 ,4 ,100,200
	res := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i-1] == nums[i]-1 {
			count++
			fmt.Println(count)
		} else {
			if count > res {
				res = count
			}
			count = 1
		}
		if count > res {
			res = count
		}

	}
	return res
}
