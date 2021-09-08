package codetop

import (
	"reflect"
	"testing"
)

func generateLinkedListFromSlice(ints []int) *ListNode {
	emptyHead := &ListNode{}
	tail := emptyHead
	for _, num := range ints {
		tail.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		tail = tail.Next
	}
	return emptyHead.Next
}

func Test_deleteDuplicates(t *testing.T) {
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
			args: args{
				head: generateLinkedListFromSlice([]int{1, 2, 3, 3, 4, 4}),
			},
			want: generateLinkedListFromSlice([]int{1, 2, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
