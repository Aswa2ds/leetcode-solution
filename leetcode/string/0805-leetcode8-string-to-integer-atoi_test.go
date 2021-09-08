package string

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "test",
			args: args{
				s: "-42",
			},
			want: -42,
		},
		{
			name: "test",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "test",
			args: args{
				s: "word front 4193",
			},
			want: 0,
		},
		{
			name: "test",
			args: args{
				s: "21474836471",
			},
			want: 2147483647,
		},
		{
			name: "test",
			args: args{
				s: "-21474836471",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
