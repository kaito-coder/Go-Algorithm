package main

import (
	"fmt"
)

func main() {
	tokens := []int{73, 74, 75}
	fmt.Println(dailyTemperatures(tokens)) // Output: 9
}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{} // this will store the indices of the temperatures

	for i := 0; i < n; i++ {
		// While the current temperature is higher than the temperature at the index
		// at the top of the stack, pop the stack and update the answer for that index.
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[topIndex] = i - topIndex
			fmt.Println(i)
			fmt.Println(topIndex)
		}
		// Push the current index onto the stack.
		stack = append(stack, i)
	}

	return answer

}
