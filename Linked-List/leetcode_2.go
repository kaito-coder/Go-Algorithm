package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(numbers []int) *ListNode {
	if len(numbers) == 0 {
		return nil
	}
	head := &ListNode{numbers[0], nil}
	current := head
	for _, v := range numbers[1:] {
		current.Next = &ListNode{v, nil}
		current = current.Next
	}
	return head
}

func printLinkedList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
}

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func main() {
	l1 := createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createLinkedList([]int{9, 9, 9, 9})
	result := addTwoNumbers(l1, l2)
	printLinkedList(result) // Output: 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{0, nil}
	current := head
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{1, nil}
	}
	return head.Next
}
