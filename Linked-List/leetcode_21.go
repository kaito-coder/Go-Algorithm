package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("hello world")
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	for list1 != nil && list1.Next != nil {
		temp1 := list1.Next
		temp2 := list2.Next
		list1.Next = list2
		list2.Next = temp1
		list1 = temp1
		list2 = temp2
	}
	return list1
}

//     nil -> 1 ->  2 ->  3  -> 4 ->  5
//	   nil <- 1 <-  2
//prev
//	  head
//  next
