package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := search(arr, target)
	fmt.Println(result)
}
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		temp := l + (r-l)/2

		if nums[temp] == target {
			return temp
		}

		if nums[temp] > target {
			r = temp - 1
		} else if nums[temp] < target {
			l = temp + 1
		}
	}
	return -1
}
