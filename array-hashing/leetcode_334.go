package main

import (
	"math"
)

func main() {

}

// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
func increasingTriplet(nums []int) bool {
	i, j := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v < i {
			i = v
		} else if v < j {
			j = v
		} else {
			return true
		}
	}
	return false
}
