package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}
func maxProfit(prices []int) int {
	res := 0
	lowest := prices[0]
	for _, v := range prices {
		if v < lowest {
			lowest = v
		}
		if (v - lowest) > res {
			res = v - lowest
		}
	}
	return res
}
