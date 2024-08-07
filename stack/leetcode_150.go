package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens)) // Output: 9
}
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		switch v {
		case "+":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b+a)
		case "-":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		default:
			temp, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			stack = append(stack, temp)
		}
	}
	return stack[0]
}
