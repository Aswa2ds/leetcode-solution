package codetop

import (
	"fmt"
	"testing"
)

func TestReversList(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	head = reverseKGroup(head, 2)
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
