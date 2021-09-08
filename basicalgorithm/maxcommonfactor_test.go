package basicalgorithm

import "testing"

func TestEuclideanAlgorithm(t *testing.T) {
	type args struct {
		a int
		b int
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
				a: 1000000000000000001,
				b: 1000000000000000000,
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				a: 100000000000,
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EuclideanAlgorithm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EuclideanAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGengXiangJianSun(t *testing.T) {
	type args struct {
		a int
		b int
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
				a: 100000000001,
				b: 100000000000,
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				a: 100000000000,
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GengXiangJianSun(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GengXiangJianSun() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCommonFactor(t *testing.T) {
	type args struct {
		a int
		b int
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
				a: 100000000001,
				b: 100000000000,
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				a: 100000000000,
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCommonFactor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxCommonFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}