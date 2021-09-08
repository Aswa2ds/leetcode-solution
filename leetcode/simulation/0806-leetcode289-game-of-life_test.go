package simulation

import (
	"fmt"
	"testing"
	"time"
)

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				board: [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, line := range tt.args.board {
				for _, life := range line {
					fmt.Print(life, "\t")
				}
				fmt.Println()
			}
			fmt.Println()
			for true {
				time.Sleep(time.Second)
				gameOfLife(tt.args.board)
				for _, line := range tt.args.board {
					for _, life := range line {
						fmt.Print(life, "\t")
					}
					fmt.Println()
				}
				fmt.Println()
			}
		})
	}
}
