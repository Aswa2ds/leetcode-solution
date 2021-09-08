package codetop

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "test",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				s: "tmmzuxt",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
