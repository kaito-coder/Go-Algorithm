package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 6, 7, 11}
	target := 8

	result := minEatingSpeed(arr, target)
	fmt.Println(result)
}
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	l, r := 1, max
	for l <= r {
		mid := l + (r-l)/2
		times := 0
		for _, v := range piles {
			times += int(math.Ceil(float64(v) / (float64(mid))))
		}
		if times == h {
			return mid
		}
		if times > h {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
