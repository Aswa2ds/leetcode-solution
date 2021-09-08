package codetop

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{4, nil, nil},
				},
			}},
			want: []int{1,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
