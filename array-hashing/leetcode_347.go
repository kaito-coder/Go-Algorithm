package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, -1, 2, -1, 2, 3}

	k := 2
	fmt.Println(topKFrequent(nums, k))
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	// Sort the keys
	sort.Ints(keys)
	return keys[:k]
}
