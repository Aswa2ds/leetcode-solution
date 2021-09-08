package hash

import "testing"

func Test_numFriendRequests(t *testing.T) {
	type args struct {
		ages []int
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
				ages: []int{20,30,100,110,120},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numFriendRequests(tt.args.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
