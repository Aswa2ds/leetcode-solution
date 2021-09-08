package basicalgorithm

import (
	"testing"
)

func TestPowerOf2(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				a: 1 << 42+1 << 20,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerOf2(tt.args.a); got != tt.want {
				t.Errorf("PowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
