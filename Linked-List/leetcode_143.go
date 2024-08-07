package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// find the middle of the linked list by way use two pointers : slow , fast
	// slow moves one step at time, while fast moves two steps at a time .
	// when fast reaches end of the list, 'slow' will be at a middle
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the list
	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}

	// Merge two halves
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}

// Helper function to create a linked list from a slice.
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print the linked list.
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Example usage:
	list := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println("Original list:")
	printLinkedList(list)

	reorderList(list)
	fmt.Println("Reordered list:")
	printLinkedList(list)
}

// 1 ,2 ,3 ,4, 5
// 1 ,5 ,2, 4, 3
