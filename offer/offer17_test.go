package offer

import (
	"reflect"
	"testing"
)

func Test_printNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{n: 3},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
