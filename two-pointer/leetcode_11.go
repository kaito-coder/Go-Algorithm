package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
		res = int(math.Max(float64(res), float64(area)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res

}
