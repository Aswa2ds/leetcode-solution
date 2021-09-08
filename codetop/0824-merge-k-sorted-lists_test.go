package codetop

import (
	"fmt"
	"reflect"
	"testing"
)

func generateLinkListByList(list []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, num := range list {
		node := &ListNode{
			Val: num,
		}
		tail.Next = node
		tail = node
	}
	return head.Next
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				lists: []*ListNode{
					generateLinkListByList([]int{1,4,5}),
					generateLinkListByList([]int{1,3,4}),
					generateLinkListByList([]int{2,6}),
				},
			},
			want: generateLinkListByList([]int{1,1,2,3,4,4,5,6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				for p := got; p != nil; p = p.Next {
					fmt.Print(p.Val)
				}
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
