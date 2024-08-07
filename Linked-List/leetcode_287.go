package main

func main() {

}
func findDuplicate(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			return v
		}
	}
	return -1
}
