// lc 21: Merge Two Sorted Lists
package two_pointers

func MergeSortedArray(n1 *ListNode, n2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for n1 != nil && n2 != nil {
		if n1.Value < n2.Value {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return dummy.Next
}
