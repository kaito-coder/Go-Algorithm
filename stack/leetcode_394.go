package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]")) // "aaabcbc"

}
func decodeString(s string) string {
	numberStack := []int{}
	stringStack := []string{}
	currNum := 0
	currStr := ""
	for _, v := range s {
		if '0' <= v && v <= '9' {
			num, _ := strconv.Atoi(string(v))
			currNum = currNum*10 + num
		} else if v == '[' {
			stringStack = append(stringStack, currStr)
			numberStack = append(numberStack, currNum)
			currNum = 0
			currStr = ""
		} else if v == ']' {
			repeat := numberStack[len(numberStack)-1]
			numberStack = numberStack[:len(numberStack)-1]
			preStr := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			currStr = string(preStr) + strings.Repeat(currStr, repeat)
		} else {
			currStr += string(v)
		}
	}
	return currStr

}
