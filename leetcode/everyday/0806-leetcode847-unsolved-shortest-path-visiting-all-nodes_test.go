package everyday

import "testing"

func Test_shortestPathLength(t *testing.T) {
	type args struct {
		graph [][]int
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
				graph: [][]int{{1},{0,2,4},{1,3,4},{2},{1,2}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathLength(tt.args.graph); got != tt.want {
				t.Errorf("shortestPathLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
