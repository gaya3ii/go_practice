package two_pointers

import "testing"

func TestHasCycle(t *testing.T) {
	// build a cycle: 1 -> 2 -> 3 -> back to 2
	cycleNode := &ListNode{Value: 2}
	cycleNode.Next = &ListNode{Value: 3, Next: cycleNode}
	cycleHead := &ListNode{Value: 1, Next: cycleNode}

	tests := []struct {
		head *ListNode
		want bool
	}{
		{head: cycleHead, want: true},
		{head: &ListNode{Value: 1, Next: &ListNode{Value: 2}}, want: false},
		{head: nil, want: false},
	}

	for _, tt := range tests {
		got := HasCycle(tt.head)
		if got != tt.want {
			t.Errorf("HasCycle() = %v, want %v", got, tt.want)
		}
	}
}
