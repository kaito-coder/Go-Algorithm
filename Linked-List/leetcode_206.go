package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

//     nil -> 1 ->  2 ->  3  -> 4 ->  5
//	   nil <- 1 <-  2
//prev
//	  head
//  next
