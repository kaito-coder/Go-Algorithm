package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	array := []byte(strings.ToLower(s))
	i, j := 0, len(array)-1
	for i < j {
		if !isChar(array[i]) {
			i++
		} else if !isChar(array[j]) {
			j--
		} else if array[i] == array[j] {
			i++
			j--
		} else {
			return false
		}

	}
	return true
}
func isChar(s byte) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	if s >= 'a' && s <= 'z' {
		return true
	}
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}
