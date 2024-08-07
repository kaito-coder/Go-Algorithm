package main

import (
	"fmt"
)

func main() {
	str := []string{"abccd", "asdas"}
	encode := encode(str)
	fmt.Println(encode)
	fmt.Println(decode(encode))
}
func encode(strs []string) string {
	var res string
	for _, s := range strs {
		res = res + string(len(s)) + "#" + s
	}
	return res
}

// neet code
// 4#need#4#code
func decode(str string) []string {
	i := 0
	res := []string{}
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		length := int(str[i])
		i = j + 1
		j = i + length
		res = append(res, str[i:j])
		i = j
	}

	return res
}
