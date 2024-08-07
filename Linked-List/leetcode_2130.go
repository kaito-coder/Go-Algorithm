package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	curr := slow
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	max := 0
	secondHalf := prev

	for secondHalf != nil {
		if (secondHalf.Val + head.Val) > max {
			max = secondHalf.Val + head.Val
		}
		secondHalf = secondHalf.Next
		head = head.Next

	}
	return max

}
