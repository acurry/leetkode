package main

import (
	"reflect"
	"testing"
)

func Test_generatePascalsTriangle(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Level 1",
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "Level 2",
			args: args{
				numRows: 2,
			},
			want: [][]int{
				{1},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePascalsTriangle(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePascalsTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
