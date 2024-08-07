package main

import (
	"fmt"
	"math"
)

func main() {
	// arr := [][]int32{
	// 	{-1, -1, 0, -9, -2, -2},
	// 	{-2, -1, -6, -8, -2, -5},
	// 	{-1, -1, -1, -2, -3, -4},
	// 	{-1, -9, -2, -4, -4, -5},
	// 	{-7, -3, -3, -2, -9, -9},
	// 	{-1, -3, -1, -2, -4, -5},
	// }
	arr := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	fmt.Println(arrayManipulation(10, arr))
}
func arrayManipulation(n int32, queries [][]int32) int64 {
    // Write your code here

	

}
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n)
	for _, row := range queries {
		for i := row[0] - 1; i < row[2]-1; i++ {
			arr[i] += int64(row[1])
		}
	}
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}
func hourglassSum(arr [][]int32) int32 {
	result := int32(math.MinInt32)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i+2 < len(arr) && j+2 < len(arr[i]) {
				sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
				if sum > int32(result) {
					result = sum
				}
			}
		}
	}
	return result
}
func dynamicArray(n int32, queries [][]int32) []int32 {
	result := make([]int32, 0)
	lastAnswer := int32(0)
	seqList := make([][]int32, n)
	for i := range seqList {
		seqList[i] = make([]int32, 0)
	}
	for _, row := range queries {
		index := (row[1] ^ lastAnswer) % n
		if row[0] == 1 {
			seqList[index] = append(seqList[index], row[2])
		} else {
			lastAnswer = seqList[index][row[2]%int32(len(seqList[index]))]
			result = append(result, lastAnswer)
		}
	}
	return result
}
func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	length := int32(len(arr))
	d = d % length
	result := append(arr[d:], arr[:d]...)
	return result
}
func matchingStrings(stringList []string, queries []string) []int32 {
	// Write your code here
	result := make([]int32, 0)
	for _, query := range queries {
		count := int32(0)
		for _, str := range stringList {
			if query == str {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
