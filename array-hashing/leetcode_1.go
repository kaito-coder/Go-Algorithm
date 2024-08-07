package main

import "fmt"

func main() {
	nums := []int{2, 2, 11, 15}
	fmt.Println(twoSum(nums, 4))
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil

}
