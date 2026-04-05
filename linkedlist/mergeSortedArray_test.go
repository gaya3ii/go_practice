package linkedlist

import "testing"

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		n1   *ListNode
		n2   *ListNode
		want []int
	}{
		{
			n1:   &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 4}}},
			n2:   &ListNode{Value: 1, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4}}},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			n1:   nil,
			n2:   &ListNode{Value: 1},
			want: []int{1},
		},
		{
			n1:   nil,
			n2:   nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		got := mergeSortedArray(tt.n1, tt.n2)
		for i, want := range tt.want {
			if got == nil || got.Value != want {
				t.Errorf("index %d: got nil or wrong value, want %d", i, want)
				break
			}
			got = got.Next
		}
	}
}
