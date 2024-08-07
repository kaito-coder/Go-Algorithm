package main

import "fmt"

func main() {
	// fmt.Println(isValid("()")) // true
	fmt.Println(isValid("()[]{}")) // true
	// fmt.Println(isValid("(]"))     // false
	// fmt.Println(isValid("([)]"))   // false
	// fmt.Println(isValid("{[]}"))   // true

}
func isValid(s string) bool {
	m := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, v := range s {
		if _, ok := m[v]; !ok {
            fmt.Printf("Pushing: %c\n", v)
			stack = append(stack, v)
			fmt.Println(m[v])
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != m[v] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
