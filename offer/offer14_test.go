package offer

import "testing"

func Test_cuttingRope(t *testing.T) {
	type args struct {
		n int
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
				n: 16,
			},
			want: 324,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
