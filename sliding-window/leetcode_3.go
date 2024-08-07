package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	res := 1
	l := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok && m[s[i]] >= l {
			l = m[s[i]] + 1
		}
		m[s[i]] = i
		if i-l+1 > res {
			res = i - l + 1
		}
	}
	return res
}
