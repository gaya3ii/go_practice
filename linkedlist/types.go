package linkedlist

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
