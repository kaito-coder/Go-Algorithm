package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	scan := 0
	write := 0
	length := len(chars)
	for scan < length {
		chars[write] = chars[scan]
		count := 0
		for scan < length && chars[scan] == chars[write] {
			count++
			scan++
		}
		if count > 1 {
			for _, c := range []byte(strconv.Itoa(count)) {
				write++
				chars[write] = c
			}
		}
	}
	return write
}
func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))

}
