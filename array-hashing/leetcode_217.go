package main

import "fmt"

func main() {
	nums := []int{2, 2, 11, 15}
	fmt.Println(containsDuplicate(nums))
}
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
