package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lead := &ListNode{0, head}
	fast := lead
	slow := lead
	for ; n > 0; n-- {
		fast = fast.Next
	}
	for fast.Next != nil && fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return lead.Next

}
