package basicalgorithm

import (
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

func TestGetIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
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
				headA: generateLinkListByList([]int{1,2,3,4,5}),
				headB: generateLinkListByList([]int{6,7,8}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReserveList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{head: generateLinkListByList([]int{1,2,3,4,5})},
			want: generateLinkListByList([]int{5,4,3,2,1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseParList(t *testing.T) {
	head := generateLinkListByList([]int{1, 2, 3, 4, 5})
	emptyhead := &ListNode{Next: head}
	pre, next := emptyhead, emptyhead.Next.Next.Next.Next.Next.Next
	left, right := ReversePartList(head, head.Next.Next.Next.Next)
	pre.Next = left
	right.Next = next
	head = emptyhead.Next
	if !reflect.DeepEqual(head, generateLinkListByList([]int{5,4,3,2,1})) {
		t.Error("reverse error")
	}
}

func TestIsExistLoop(t *testing.T) {
	head := generateLinkListByList([]int{1})
	if IsExistLoop(head) {
		t.Error("Loop Judge Error")

	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = head
	if !IsExistLoop(head) {
		t.Error("Loop Judge Error")
	}
}