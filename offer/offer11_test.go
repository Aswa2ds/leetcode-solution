package offer

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				numbers: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				numbers: []int{1},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				numbers: []int{1, 3, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
