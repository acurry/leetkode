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
		{
			name: "Level 6",
			args: args{
				numRows: 6,
			},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
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

func Test_getPascalTriangleRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Row 1",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			name: "Row 2",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
		{
			name: "Row 3",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPascalTriangleRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPascalTriangleRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
