package main

import (
	"fmt"
)

func main() {
	s := []int{2, 7, 11, 15}
	fmt.Println(twoSum(s, 9))
}
func twoSum(numbers []int, target int) []int {
	leftPointer := 0
	rightPointer := len(numbers) - 1
	for leftPointer < rightPointer {
		sum := numbers[leftPointer] + numbers[rightPointer]
		if sum > target {
			rightPointer--
		} else if sum < target {
			leftPointer++
		} else {
			return []int{numbers[leftPointer], numbers[rightPointer]}
		}
	}
	return []int{}
}
