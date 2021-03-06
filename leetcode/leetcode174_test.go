package leetcode

import "testing"

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{dungeon: [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
