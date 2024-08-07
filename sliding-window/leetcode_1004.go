package main

func main() {

}

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
func longestOnes(nums []int, k int) int {
	res := 0
	zeroCount := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		res = max(res, i-start+1)
	}
	return res
}
