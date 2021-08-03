package offer

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{0,0,1},
				k:   6,
			},
			want: []int{1,2,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
