package main

func main() {

}
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	// Separate odd and even nodes
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	// Attach even list to the end of odd list
	odd.Next = evenHead

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
