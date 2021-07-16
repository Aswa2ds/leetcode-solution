package offer

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				s: "xxxxx xxxx",
			},
			want: "xxxxx%02xxxx",
		},
		{
			name: "test 2",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "test 3",
			args: args{
				s: "  ",
			},
			want: "%02%02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
